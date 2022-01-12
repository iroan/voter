package keeper

import (
	"github.com/iroan/voter/x/voter/types"
)

var _ types.QueryServer = Keeper{}
