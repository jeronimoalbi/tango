package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "tango/testutil/keeper"
	"tango/x/tango/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TangoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
