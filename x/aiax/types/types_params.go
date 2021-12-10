package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	ParamStoreKeyEnableErc20Mapping = []byte("EnableErc20Mapping")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(enableErc20Mapping bool) Params {
	return Params{
		EnableErc20Mapping: enableErc20Mapping,
	}
}

func DefaultParams() Params {
  return Params{
    EnableErc20Mapping: true,
  }
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyEnableErc20Mapping, &p.EnableErc20Mapping, validateBool),
	}
}

func (p Params) Validate() error {
	return nil
}
