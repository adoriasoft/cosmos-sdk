package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 2, "name does not exist")
)
