package keeper_test

import (
	"testing"

	testkeeper "github.com/jeronimoalbi/tango/testutil/keeper"
	"github.com/jeronimoalbi/tango/x/tango/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TangoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
