package transactioninactivity

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vpavlin/temporal-experiment/transaction-in-activity/bindings"
	"go.temporal.io/sdk/activity"
)

type TxCreationReverted error
type TxPending error

type Blockchain struct {
	client     *ethclient.Client
	chainId    uint
	privateKey *ecdsa.PrivateKey
}

func NewBlockchain(rpc string, chainId uint, privateKeyStr string) (*Blockchain, error) {
	c, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	return &Blockchain{
		client:     c,
		chainId:    chainId,
		privateKey: privateKey,
	}, nil
}

func (b *Blockchain) Mint(ctx context.Context, params Mint) (MintResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Starting activity on network %d: %v", b.chainId, params)

	contract, err := bindings.NewSCollectionV1(common.HexToAddress(params.Contract), b.client)
	if err != nil {
		logger.Error("Failed to instantiate contract", err)
		return MintResult{}, err
	}

	var result MintResult
	if activity.HasHeartbeatDetails(ctx) {
		err = activity.GetHeartbeatDetails(ctx, &result)
		if err != nil {
			logger.Error("Failed to get heartbeat details: ", err)
			return MintResult{}, err
		}
	}

	// Transaction exists, let's see if it succeeded and try to extract event logs
	if len(result.Hash) != 0 {
		logger.Info("Found existing TX, trying to get the receipt", "Hash", result.Hash, "Retries", result.Retries)

		receipt, err := b.client.TransactionReceipt(context.Background(), common.HexToHash(result.Hash))
		if err != nil {
			return MintResult{}, err
		}

		if receipt != nil && receipt.Status == 1 {
			for _, vLog := range receipt.Logs {
				event, err := contract.ParseTransfer(*vLog)
				if err != nil {
					if err.Error() == "event signature mismatch" {
						continue
					} else {
						return MintResult{}, err
					}
				} else {
					result.TokenId = uint(event.TokenId.Uint64())
					result.Success = true
					break
				}
			}
		}
	}

	// If this is the first run, submit new TX. If it's over retry limit, resend the TX (maybe the gasPrice was too low...)
	if result.Retries == 0 || result.Retries > 5 {
		senderAddress := crypto.PubkeyToAddress(b.privateKey.PublicKey)

		gasPrice, err := b.client.SuggestGasPrice(ctx)
		if err != nil {
			logger.Error("Failed to get gas price: ", err)
			return MintResult{}, err
		}

		txOpts, err := bind.NewKeyedTransactorWithChainID(b.privateKey, big.NewInt(int64(b.chainId)))
		if err != nil {
			logger.Error("Failed to produce new transactor: ", err)
			return MintResult{}, err
		}

		// Get Nonce in case of first submit
		if result.Nonce == 0 {
			nonce, err := b.client.PendingNonceAt(ctx, senderAddress)
			if err != nil {
				logger.Error("Failed to get nonce: ", err)
				return MintResult{}, err
			}
			result.Nonce = int64(nonce)
		}

		txOpts.Nonce = big.NewInt(result.Nonce)
		txOpts.GasPrice = gasPrice
		txOpts.NoSend = true

		// If this fails the TX is "wrong" and won't ever succeed
		// There is a eth_Call hidden in Geth, so if this is failing, it means the TX would get reverted and there is no point in submitting it
		tx, err := contract.Mint(txOpts, common.HexToAddress(params.To), big.NewInt(int64(params.Quantity)))
		if err != nil {
			logger.Error("Failed to produce transaction: ", err)
			return MintResult{}, TxCreationReverted(err)
		}

		result.Hash = tx.Hash().Hex()

		err = b.client.SendTransaction(ctx, tx)
		if err != nil {
			logger.Error("Failed to send TX: ", err)
			return result, err
		}
	}

	result.Retries++
	activity.RecordHeartbeat(ctx, result)

	// Fail, so that Temporal retries
	if !result.Success {
		return MintResult{}, TxPending(fmt.Errorf("Transaction %s still in pending state", result.Hash))
	}

	return result, nil
}
