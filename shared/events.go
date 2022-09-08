package shared

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vpavlin/temporal-experiment/bindings"
	"go.temporal.io/sdk/activity"
)

type Transfer struct {
	TxHash  string
	From    string
	To      string
	TokenId int64
}

func ERC721WaitForTransferEvent(
	ctx context.Context,
	client *ethclient.Client,
	contractAddress common.Address,
	fromFilter []common.Address,
	toFilter []common.Address,
	tokenIdFilter []*big.Int,
	expectedEventCount int,
) ([]Transfer, error) {
	logger := activity.GetLogger(ctx)
	var result []Transfer
	parsed, err := abi.JSON(strings.NewReader(bindings.ERC721ABI))
	if err != nil {
		return result, err
	}

	var from []interface{}
	for _, item := range fromFilter {
		from = append(from, item)
	}

	var to []interface{}
	for _, item := range toFilter {
		to = append(to, item)
	}

	var tokenIds []interface{}

	for _, item := range tokenIdFilter {
		tokenIds = append(tokenIds, item)
	}

	boundContract := bind.NewBoundContract(contractAddress, parsed, nil, nil, client)
	logs, sub, err := boundContract.WatchLogs(
		&bind.WatchOpts{Context: ctx},
		"Transfer",
		from,
		to,
		tokenIds,
	)
	if err != nil {
		return result, err
	}

	logger.Info("Subscribed to events Transfer", "Contract", contractAddress)
	defer sub.Unsubscribe()

	cnt := 0
	for {
		select {
		case log := <-logs:
			// New log arrived, parse the event and forward to the user
			event := new(bindings.SCollectionV1Transfer)
			if err := boundContract.UnpackLog(event, "Transfer", log); err != nil {
				return result, err
			}
			event.Raw = log

			transfer := Transfer{
				TxHash:  log.TxHash.Hex(),
				From:    event.From.Hex(),
				To:      event.To.Hex(),
				TokenId: event.TokenId.Int64(),
			}
			result = append(result, transfer)
			cnt++

			if cnt == expectedEventCount {
				return result, nil
			}

		case err := <-sub.Err():
			return result, err
		}
	}
}
