package types

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ParamStoreKeyEnableErc20Mapping = []byte("EnableErc20Mapping")
  ParamStoreKeyAiaxTokenContractAddress = []byte("AiaxTokenContractAddress");
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(enableErc20Mapping bool, aiaxTokenContractAddress string) Params {
	return Params{
		EnableErc20Mapping: enableErc20Mapping,
    AiaxTokenContractAddress: aiaxTokenContractAddress,
	}
}

func DefaultParams() Params {
	return Params{
		EnableErc20Mapping: true,
    AiaxTokenContractAddress: "",
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
    paramtypes.NewParamSetPair(ParamStoreKeyAiaxTokenContractAddress, &p.AiaxTokenContractAddress, validateAiaxTokenContractAddress),
	}
}

func validateAiaxTokenContractAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter of type: %T", i)
	}
	if !common.IsHexAddress(v) {
		return fmt.Errorf("not an ethereum address: %s", v)
	}
	return nil
}

func (p Params) Validate() error {
	if err := validateAiaxTokenContractAddress(p.AiaxTokenContractAddress); err != nil {
		return sdkerrors.Wrap(err, "Address of Ethereum mainnet Aiax token.")
	}
	return nil
}
