package queryworkflowstate

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var QueryType = "mintState"

type Mint struct {
	Contract string
	To       string
	Quantity int
}

type Transaction struct {
	Contract string
	Hash     string
	Nonce    int64
	Retries  int
}

type MintResult struct {
	Hash    string
	Success bool
	TokenId []uint
}

type ReceiptReadySignal struct {
	Receipt *types.Receipt
}

func Workflow(ctx workflow.Context, params Mint) (MintResult, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting a workflow: %v", params)

	currentState := "started" // This could be any serializable struct.

	err := workflow.SetQueryHandler(ctx, QueryType, func() (string, error) {
		return currentState, nil
	})
	if err != nil {
		return MintResult{}, err
	}

	activityOptionsSendTx := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        3 * time.Second,
			BackoffCoefficient:     1.2,
			MaximumAttempts:        2,
			NonRetryableErrorTypes: []string{"TxCreationReverted"},
		},
	}

	activityOptionsGetEvent := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval: 0,
			MaximumAttempts: 1,
		},
	}

	workflow.Sleep(ctx, time.Second*10)

	b := Blockchain{}

	var result MintResult
	var tx Transaction
	for i := 0; i < 5; i++ {
		currentState = fmt.Sprintf("retry %d, sending TX", i)
		ctx = workflow.WithActivityOptions(ctx, activityOptionsSendTx)

		nonce := -1
		if tx.Nonce > 0 {
			nonce = int(tx.Nonce)
		}

		workflow.Sleep(ctx, time.Second*10)

		err := workflow.ExecuteActivity(ctx, b.SendTx, params, nonce).Get(ctx, &tx)
		if err != nil {
			logger.Error("Activity failed", "Error", err)
			return MintResult{}, err
		}

		if len(tx.Hash) == 0 {
			return result, fmt.Errorf("Failed to produce TX: %v", tx)
		}

		currentState = fmt.Sprintf("retry %d, TX submitted (%s), waiting for event", i, tx.Hash)

		workflow.Sleep(ctx, time.Second*10)

		logger.Info("Submitted TX, waiting for event", "Hash", tx.Hash)

		ctx = workflow.WithActivityOptions(ctx, activityOptionsGetEvent)
		err = workflow.ExecuteActivity(ctx, b.WaitForMint, params).Get(ctx, &result)
		if err != nil {
			logger.Error("Failed to get event", "Error", err)
		}

		currentState = fmt.Sprintf("retry %d, event received", i)

		if result.Success {
			currentState = fmt.Sprintf("retry %d, mint successful, produced TokenId %d", i, result.TokenId)
			break
		}

		workflow.Sleep(ctx, time.Second*10)

	}

	return result, nil

}
