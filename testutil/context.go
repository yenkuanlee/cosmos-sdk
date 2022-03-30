package testutil

import (
	"github.com/yenkuanlee/tendermint/libs/log"
	tmproto "github.com/yenkuanlee/tendermint/proto/tendermint/types"
	dbm "github.com/yenkuanlee/tm-db"

	"github.com/yenkuanlee/cosmos-sdk/store"
	sdk "github.com/yenkuanlee/cosmos-sdk/types"
)

// DefaultContext creates a sdk.Context with a fresh MemDB that can be used in tests.
func DefaultContext(key sdk.StoreKey, tkey sdk.StoreKey) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tkey, sdk.StoreTypeTransient, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return ctx
}
