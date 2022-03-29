package types_test

import (
	"github.com/yenkuanlee/cosmos-sdk/simapp"
)

var (
	app                   = simapp.Setup(false)
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Marshaler, ecdc.Amino
)
