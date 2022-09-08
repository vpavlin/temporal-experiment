package shared

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vpavlin/temporal-experiment/bindings"
)

func ICollectionV1GetSpec(ctx context.Context, client *ethclient.Client, contractAddress common.Address) (bindings.ICollectionV1Spec, error) {
	var spec bindings.ICollectionV1Spec

	contract, err := bindings.NewSCollectionV1(contractAddress, client)
	if err != nil {
		return spec, err
	}

	spec, err = contract.GetSpec(&bind.CallOpts{Context: ctx})
	if err != nil {
		return spec, err
	}

	return spec, nil
}
