package bundler

import (
	"context"

	"github.com/dbadoy/go-paymaster/pkg/userop"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type BundlerClient struct {
	entryPoint common.Address
	c          *rpc.Client
}

func (b *BundlerClient) SendUserOperation(ctx context.Context, op *userop.UserOperation) (common.Hash, error) {
	var result common.Hash
	err := b.c.CallContext(ctx, &result, "eth_sendUserOperation", op, b.entryPoint)
	return result, err
}

func (b *BundlerClient) EstimateUserOperationGas(ctx context.Context, op *userop.UserOperation) (*GasEstimates, error) {
	var estimate GasEstimates
	err := b.c.CallContext(ctx, &estimate, "eth_estimateUserOperationGas", op, b.entryPoint)
	if err != nil {
		return nil, err
	}
	return &estimate, nil
}
