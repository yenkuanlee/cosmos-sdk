package codec

import (
	codectypes "github.com/yenkuanlee/cosmos-sdk/codec/types"
	"github.com/yenkuanlee/cosmos-sdk/crypto/keys/ed25519"
	"github.com/yenkuanlee/cosmos-sdk/crypto/keys/multisig"
	"github.com/yenkuanlee/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/yenkuanlee/cosmos-sdk/crypto/keys/secp256r1"
	cryptotypes "github.com/yenkuanlee/cosmos-sdk/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	var pk *cryptotypes.PubKey
	registry.RegisterInterface("cosmos.crypto.PubKey", pk)
	registry.RegisterImplementations(pk, &ed25519.PubKey{})
	registry.RegisterImplementations(pk, &secp256k1.PubKey{})
	registry.RegisterImplementations(pk, &multisig.LegacyAminoPubKey{})
	secp256r1.RegisterInterfaces(registry)
}
