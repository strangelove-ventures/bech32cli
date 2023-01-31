package cmd

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func transformCmd(a *appState) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transform [bech32Address] [newBech32Prefix]",
		Aliases: []string{"t"},
		Short:   "Transforms bech32 string to new prefix",
		Args:    cobra.ExactArgs(2),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s transform cosmos1ge60jkvf2wygslexprqgshxgmzd6zqludz8wyt osmo
$ %s t cosmos1ge60jkvf2wygslexprqgshxgmzd6zqludz8wyt osmo`,
			appName, appName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			address, newPrefix := args[0], args[1]

			sdkAddr, err := types.AccAddressFromBech32(address)
			if err != nil {
				return fmt.Errorf("failed to decode [bech32Address]: %s - %w", address, err)
			}

			newAddr, err := types.Bech32ifyAddressBytes(newPrefix, sdkAddr)
			if err != nil {
				return fmt.Errorf("failed to encode with [newBech32Prefix]: %s - %w", newPrefix, err)
			}

			if newAddr == "" {
				return fmt.Errorf("failed to encode with [newBech32Prefix]: %s", newPrefix)
			}

			fmt.Fprintln(cmd.OutOrStdout(), newAddr)

			return nil
		},
	}
	return cmd
}
