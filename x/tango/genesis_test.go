package tango_test

import (
	"testing"

	keepertest "github.com/jeronimoalbi/tango/testutil/keeper"
	"github.com/jeronimoalbi/tango/testutil/nullify"
	"github.com/jeronimoalbi/tango/x/tango"
	"github.com/jeronimoalbi/tango/x/tango/types"
	"github.com/stretchr/testify/require"
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
