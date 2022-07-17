package keeper

import (
	"github.com/zireael26/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
