package keeper

import (
	"github.com/jeronimoalbi/tango/x/tango/types"
)

var _ types.QueryServer = Keeper{}
