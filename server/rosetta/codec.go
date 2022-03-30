package rosetta

import (
	"github.com/yenkuanlee/cosmos-sdk/codec"
	codectypes "github.com/yenkuanlee/cosmos-sdk/codec/types"
	cryptocodec "github.com/yenkuanlee/cosmos-sdk/crypto/codec"
	authcodec "github.com/yenkuanlee/cosmos-sdk/x/auth/types"
	bankcodec "github.com/yenkuanlee/cosmos-sdk/x/bank/types"
)

// MakeCodec generates the codec required to interact
// with the cosmos APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}
