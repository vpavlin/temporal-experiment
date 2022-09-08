package propagators

import (
	"context"
	"fmt"
	"reflect"

	"github.com/vpavlin/temporal-experiment/shared/state"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

type statePropagator struct{}

type stateContextKey struct{}

var StateKey = stateContextKey{}
var StateKeyStr = reflect.TypeOf(StateKey).Name()

func NewStateContextPropagator() workflow.ContextPropagator {
	return &statePropagator{}
}

func (sp *statePropagator) Inject(ctx context.Context, writer workflow.HeaderWriter) error {
	valueInterface := ctx.Value(StateKey)
	if valueInterface == nil {
		return nil
	}

	err := write(valueInterface, writer)
	if err != nil {
		return err
	}

	return nil
}

func (sp *statePropagator) InjectFromWorkflow(ctx workflow.Context, writer workflow.HeaderWriter) error {
	valueInterface := ctx.Value(StateKey)
	if valueInterface == nil {
		return nil
	}

	logger := workflow.GetLogger(ctx)
	logger.Warn(fmt.Sprintf("Injecting: %v", valueInterface))

	err := write(valueInterface, writer)
	if err != nil {
		return err
	}

	return nil
}

func (sp *statePropagator) Extract(ctx context.Context, reader workflow.HeaderReader) (context.Context, error) {
	state, err := read(reader)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, StateKey, state)

	return ctx, nil
}

func (sp *statePropagator) ExtractToWorkflow(ctx workflow.Context, reader workflow.HeaderReader) (workflow.Context, error) {
	state, err := read(reader)
	if err != nil {
		return nil, err
	}

	logger := workflow.GetLogger(ctx)

	logger.Warn(fmt.Sprintf("Extracting: %v", state))

	ctx = workflow.WithValue(ctx, StateKey, state)

	return ctx, nil
}

/// Internal
func write(valueInterface interface{}, writer workflow.HeaderWriter) error {
	value, ok := valueInterface.(*state.StatesCollection)
	if !ok {
		return fmt.Errorf("Failed to cast context value as %s", StateKeyStr)
	}

	encodedValue, err := converter.GetDefaultDataConverter().ToPayload(value)
	if err != nil {
		return err
	}

	writer.Set(StateKeyStr, encodedValue)

	return nil
}

func read(reader workflow.HeaderReader) (*state.StatesCollection, error) {
	payload, ok := reader.Get(StateKeyStr)
	if !ok {
		return nil, nil
	}

	var state *state.StatesCollection
	err := converter.GetDefaultDataConverter().FromPayload(payload, &state)
	if err != nil {
		return nil, err
	}

	return state, nil
}
