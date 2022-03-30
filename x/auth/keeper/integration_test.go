package keeper_test

import (
	tmproto "github.com/yenkuanlee/tendermint/proto/tendermint/types"

	"github.com/yenkuanlee/cosmos-sdk/simapp"
	sdk "github.com/yenkuanlee/cosmos-sdk/types"
	authtypes "github.com/yenkuanlee/cosmos-sdk/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}
