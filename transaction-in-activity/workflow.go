package transactioninactivity

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type Mint struct {
	Contract string
	To       string
	Quantity int
}

type MintResult struct {
	Hash    string
	Nonce   int64
	Retries int
	Success bool
	TokenId uint
}

func Workflow(ctx workflow.Context, params Mint) (MintResult, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting a workflow: %v", params)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			NonRetryableErrorTypes: []string{"TxCreationReverted"},
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result MintResult
	err := workflow.ExecuteActivity(ctx, "Mint", params).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed", "Error", err)
		return MintResult{}, err
	}

	return result, nil

}
