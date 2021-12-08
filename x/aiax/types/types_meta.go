package types

import "fmt"

func CreateDenom(address string) string {
	return fmt.Sprintf("%s/%s", ModuleName, address)
}

func CreateDenomDescription(address string) string {
	return fmt.Sprintf("Aiax representation of Ethereum ERC20 %s", address)
}
