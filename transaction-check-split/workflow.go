package transactionchecksplit

import (
	"fmt"
	"time"

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

func Workflow(ctx workflow.Context, params Mint) (MintResult, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting a workflow: %v", params)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        3 * time.Second,
			BackoffCoefficient:     1.2,
			MaximumAttempts:        2,
			NonRetryableErrorTypes: []string{"TxCreationReverted"},
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	b := Blockchain{}

	var tx Transaction
	err := workflow.ExecuteActivity(ctx, b.SendTx, params, -1).Get(ctx, &tx)
	if err != nil {
		logger.Error("Activity failed", "Error", err)
		return MintResult{}, err
	}

	var result MintResult
	if len(tx.Hash) == 0 {
		return result, fmt.Errorf("Failed to produce TX: %v", tx)
	}

	ao.RetryPolicy.MaximumAttempts = 5
	ctx = workflow.WithActivityOptions(ctx, ao)
	err = workflow.ExecuteActivity(ctx, b.CheckTx, tx).Get(ctx, &result)
	if err != nil {
		return result, err
	}

	return result, nil

}
