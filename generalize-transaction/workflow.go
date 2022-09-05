package generalizetransaction

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

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
	TokenId uint
}

type ReceiptReadySignal struct {
	Receipt *types.Receipt
}

func Workflow(ctx workflow.Context, params Mint) (MintResult, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting a workflow: %v", params)

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
		StartToCloseTimeout: 10 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval: 0,
			MaximumAttempts: 1,
		},
	}

	b := Blockchain{}

	var result MintResult
	var tx Transaction
	for i := 0; i < 5; i++ {
		ctx = workflow.WithActivityOptions(ctx, activityOptionsSendTx)

		nonce := -1
		if tx.Nonce > 0 {
			nonce = int(tx.Nonce)
		}

		err := workflow.ExecuteActivity(ctx, b.SendTx, params, nonce).Get(ctx, &tx)
		if err != nil {
			logger.Error("Activity failed", "Error", err)
			return MintResult{}, err
		}

		if len(tx.Hash) == 0 {
			return result, fmt.Errorf("Failed to produce TX: %v", tx)
		}

		logger.Info("Submitted TX, waiting for event", "Hash", tx.Hash)

		ctx = workflow.WithActivityOptions(ctx, activityOptionsGetEvent)
		err = workflow.ExecuteActivity(ctx, b.GetEvent, tx, params.To).Get(ctx, &result)
		if err != nil {
			logger.Error("Failed to get event", "Error", err)
		}
		if result.Success {
			break
		}
	}

	return result, nil

}
