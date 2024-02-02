package cmd

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/spf13/cobra"
)

const (
	flagAddress = "address"
	flagPubkey  = "pubkey"
)

func valConsCmd(a *appState) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "valcons",
		Aliases: []string{"v"},
		Short:   "validator consensus address transformation",
		Args:    cobra.ExactArgs(1),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s valcons osmo --pubkey wC+QT4cw8WWOwRZhL/XZ8XusXSH7Q3kvhEnFFPagXis=
$ %s v osmo --pubkey wC+QT4cw8WWOwRZhL/XZ8XusXSH7Q3kvhEnFFPagXis=
$ %s v osmo --address 023DCF3F6AEA4E0098ABBA2AF23F3D65AC324851
$ %s v osmo --address 023DCF3F6AEA4E0098ABBA2AF23F3D65AC324851`,
			appName, appName, appName, appName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			prefix := args[0]
			address, _ := cmd.Flags().GetString(flagAddress)
			pubkey, _ := cmd.Flags().GetString(flagPubkey)

			if address == "" && pubkey == "" {
				return fmt.Errorf("either --address or --pubkey must be specified")
			}

			if address != "" && pubkey != "" {
				return fmt.Errorf("either --address or --pubkey must be specified, not both")
			}

			var err error
			var addrBz []byte

			if address != "" {
				addrBz, err = hex.DecodeString(address)
				if err != nil {
					return fmt.Errorf("failed to decode address: %s - %w", address, err)
				}
			} else {
				pubkeyBz, err := base64.StdEncoding.DecodeString(pubkey)
				if err != nil {
					return fmt.Errorf("failed to decode pubkey: %s - %w", pubkey, err)
				}
				hash := sha256.Sum256(pubkeyBz)
				addrBz = hash[:20]
			}

			valcons, err := bech32.ConvertAndEncode(prefix+"valcons", addrBz)
			if err != nil {
				return fmt.Errorf("failed to encode with prefix: %s - %w", prefix, err)
			}
			fmt.Println(valcons)

			return nil
		},
	}

	cmd.Flags().String(flagAddress, "", "validator hex address to transform")
	cmd.Flags().String(flagPubkey, "", "validator base64 pubkey to transform")
	return cmd
}
