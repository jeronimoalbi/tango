package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jeronimoalbi/tango/testutil/keeper"
	"github.com/jeronimoalbi/tango/x/tango/keeper"
	"github.com/jeronimoalbi/tango/x/tango/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TangoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
