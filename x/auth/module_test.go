package auth_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	abcitypes "github.com/yenkuanlee/tendermint/abci/types"
	tmproto "github.com/yenkuanlee/tendermint/proto/tendermint/types"

	"github.com/yenkuanlee/cosmos-sdk/simapp"
	"github.com/yenkuanlee/cosmos-sdk/x/auth/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	acc := app.AccountKeeper.GetAccount(ctx, types.NewModuleAddress(types.FeeCollectorName))
	require.NotNil(t, acc)
}
