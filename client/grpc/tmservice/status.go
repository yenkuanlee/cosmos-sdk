package tmservice

import (
	"context"

	ctypes "github.com/yenkuanlee/tendermint/rpc/core/types"

	"github.com/yenkuanlee/cosmos-sdk/client"
)

func getNodeStatus(ctx context.Context, clientCtx client.Context) (*ctypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}
	return node.Status(ctx)
}
