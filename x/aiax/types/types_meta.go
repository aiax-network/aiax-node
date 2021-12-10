package types

import "fmt"

func CreateDenom(address string) string {
	return fmt.Sprintf("%s/%s", ModuleName, address)
}

func CreateDenomDescription(symbol string, name string, address string) string {
	return fmt.Sprintf("Aiax ERC20 %s/%s/%s", symbol, name, address)
}
