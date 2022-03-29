package gov_test

import (
	"strings"
	"testing"

	tmproto "github.com/yenkuanlee/tendermint/proto/tendermint/types"

	"github.com/yenkuanlee/cosmos-sdk/testutil/testdata"

	"github.com/stretchr/testify/require"

	sdk "github.com/yenkuanlee/cosmos-sdk/types"
	"github.com/yenkuanlee/cosmos-sdk/x/gov"
	"github.com/yenkuanlee/cosmos-sdk/x/gov/keeper"
)

func TestInvalidMsg(t *testing.T) {
	k := keeper.Keeper{}
	h := gov.NewHandler(k)

	res, err := h(sdk.NewContext(nil, tmproto.Header{}, false, nil), testdata.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized gov message type"))
}
