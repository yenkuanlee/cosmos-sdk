package genutil

import (
	abci "github.com/yenkuanlee/tendermint/abci/types"

	"github.com/yenkuanlee/cosmos-sdk/client"
	sdk "github.com/yenkuanlee/cosmos-sdk/types"
	"github.com/yenkuanlee/cosmos-sdk/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) (validators []abci.ValidatorUpdate, err error) {
	if len(genesisState.GenTxs) > 0 {
		validators, err = DeliverGenTxs(ctx, genesisState.GenTxs, stakingKeeper, deliverTx, txEncodingConfig)
	}
	return
}
