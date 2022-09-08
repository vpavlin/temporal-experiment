package queryworkflowstate

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vpavlin/temporal-experiment/bindings"
	"github.com/vpavlin/temporal-experiment/shared"

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

func (b *Blockchain) WaitForMint(ctx context.Context, params Mint) (MintResult, error) {
	result, err := shared.ERC721WaitForTransferEvent(
		ctx,
		b.client,
		common.HexToAddress(params.Contract),
		[]common.Address{shared.ZeroAddress},
		[]common.Address{common.HexToAddress(params.To)},
		[]*big.Int{},
		params.Quantity,
	)
	if err != nil {
		return MintResult{}, err
	}

	if len(result) == 0 {
		return MintResult{}, nil
	}

	tokenIds := make([]uint, len(result))
	for i, item := range result {
		tokenIds[i] = uint(item.TokenId)
	}

	return MintResult{
		Hash:    result[0].TxHash,
		Success: true,
		TokenId: tokenIds,
	}, nil

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

	if nonce == -1 {
		gasPrice = big.NewInt(5000)
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
	//txOpts.NoSend = true

	// If this fails the TX is "wrong" and won't ever succeed
	// There is a eth_Call hidden in Geth, so if this is failing, it means the TX would get reverted and there is no point in submitting it
	tx, err := contract.Mint(txOpts, common.HexToAddress(params.To), big.NewInt(int64(params.Quantity)))
	if err != nil {
		logger.Error("Failed to produce transaction: ", err)
		return result, TxCreationReverted{err}
	}

	result.Hash = tx.Hash().Hex()
	result.Contract = params.Contract
	result.Retries++

	return result, nil
}
