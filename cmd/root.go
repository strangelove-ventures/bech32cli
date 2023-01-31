package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const appName = "bech32"

type appState struct {
	Log *zap.Logger
}

func NewRootCmd(log *zap.Logger) *cobra.Command {
	a := &appState{Log: log}

	var rootCmd = &cobra.Command{
		Use:   "bech32",
		Short: "bech32 - a simple CLI to transform bech32 addresses",
		Long:  `bech32 is a CLI to transform bech32 addresses`,
	}

	rootCmd.AddCommand(
		transformCmd(a),
		versionCmd(a),
	)

	return rootCmd
}

func Execute() {
	rootCmd := NewRootCmd(nil)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
