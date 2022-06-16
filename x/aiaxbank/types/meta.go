package types

import "fmt"

func CreateDenom(origin string, address string) string {
	return fmt.Sprintf("%s/%s", origin, address)
}

func CreateDenomDescription(symbol string, name string, address string) string {
	return fmt.Sprintf("Aiax ERC20 %s/%s/%s", symbol, name, address)
}
