package tango_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "tango/testutil/keeper"
	"tango/testutil/nullify"
	"tango/x/tango"
	"tango/x/tango/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TangoKeeper(t)
	tango.InitGenesis(ctx, *k, genesisState)
	got := tango.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
