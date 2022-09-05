package transactionchecksplit

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

type TxCreationReverted struct {
	err error
}

func (e TxCreationReverted) Error() string {
	return e.err.Error()
}

type TxPending struct {
	err error
}

func (e TxPending) Error() string {
	return e.err.Error()
}

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

func (b *Blockchain) CheckTx(ctx context.Context, params Transaction) (MintResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Starting SendTX activity: %v", b.chainId, params)

	contract, err := bindings.NewSCollectionV1(common.HexToAddress(params.Contract), b.client)
	if err != nil {
		logger.Error("Failed to instantiate contract", err)
		return MintResult{}, err
	}

	if activity.HasHeartbeatDetails(ctx) {
		err = activity.GetHeartbeatDetails(ctx, &params)
		if err != nil {
			logger.Error("Failed to get heartbeat details: ", err)
			return MintResult{}, err
		}
	}

	logger.Info("Found existing TX, trying to get the receipt", "Hash", params.Hash, "Retries", params.Retries)

	receipt, err := b.client.TransactionReceipt(context.Background(), common.HexToHash(params.Hash))
	if err != nil {
		return MintResult{}, err
	}

	var result MintResult
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
				result.Hash = params.Hash
				break
			}
		}
	}

	params.Retries++
	activity.RecordHeartbeat(ctx, params)

	// Fail, so that Temporal retries
	if !result.Success {
		return MintResult{}, TxPending{fmt.Errorf("Transaction %s still in pending state", params.Hash)}
	}

	return result, nil
}

func (b *Blockchain) SendTx(ctx context.Context, params Mint, nonce int) (Transaction, error) {
	logger := activity.GetLogger(ctx)

	contract, err := bindings.NewSCollectionV1(common.HexToAddress(params.Contract), b.client)
	if err != nil {
		logger.Error("Failed to instantiate contract", err)
		return Transaction{}, err
	}

	senderAddress := crypto.PubkeyToAddress(b.privateKey.PublicKey)

	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		logger.Error("Failed to get gas price: ", err)
		return Transaction{}, err
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(b.privateKey, big.NewInt(int64(b.chainId)))
	if err != nil {
		logger.Error("Failed to produce new transactor: ", err)
		return Transaction{}, err
	}

	var result Transaction
	// Get Nonce in case of first submit
	if nonce < 0 {
		nonce, err := b.client.PendingNonceAt(ctx, senderAddress)
		if err != nil {
			logger.Error("Failed to get nonce: ", err)
			return result, err
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
		return result, TxCreationReverted{err}
	}

	err = b.client.SendTransaction(ctx, tx)
	if err != nil {
		if err.Error() == "nonce too low" {
			err = TxCreationReverted{err}
		}
		logger.Error("Failed to send TX: ", err)
		return result, err
	}

	result.Hash = tx.Hash().Hex()
	result.Contract = params.Contract
	result.Retries++

	return result, nil
}
