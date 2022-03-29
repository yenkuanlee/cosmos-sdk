package types

import (
	"github.com/yenkuanlee/cosmos-sdk/codec"
	cryptocodec "github.com/yenkuanlee/cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
