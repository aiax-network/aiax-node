package cli

import (
	"github.com/aiax-network/aiax-node/x/aiax/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all intrarelayer CLI query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the aiax module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	return cmd
}

