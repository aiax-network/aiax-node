package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrTokenMapping = sdkerrors.Register(ModuleName, 2, "erc20 token mapping error")
)
