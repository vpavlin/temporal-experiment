package shared

import "github.com/vpavlin/temporal-experiment/shared/state"

func QueryHandlerStatus(stateCol *state.StatesCollection) func() (state.StatesCollection, error) {
	return func() (state.StatesCollection, error) {
		return *stateCol, nil
	}
}
