package cmd_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/strangelove-ventures/bech32cli/cmd"
	"github.com/stretchr/testify/require"
)

type transformTestCase struct {
	name          string
	bech32Address string
	newPrefix     string
	expectedOut   string
	expectedErr   error
}

func TestCmdTransform(t *testing.T) {
	testCases := []transformTestCase{
		{
			name:          "happy path cosmos",
			bech32Address: "cosmos1ge60jkvf2wygslexprqgshxgmzd6zqludz8wyt",
			newPrefix:     "juno",
			expectedOut:   "juno1ge60jkvf2wygslexprqgshxgmzd6zqlumsy4rh",
			expectedErr:   nil,
		},
		{
			name:          "happy path juno",
			bech32Address: "juno1ge60jkvf2wygslexprqgshxgmzd6zqlumsy4rh",
			newPrefix:     "cosmos",
			expectedOut:   "cosmos1ge60jkvf2wygslexprqgshxgmzd6zqludz8wyt",
			expectedErr:   nil,
		},
		{
			name:          "invalid bech32 address",
			bech32Address: "invalid_bech32",
			newPrefix:     "cosmos",
			expectedErr:   fmt.Errorf("failed to decode [bech32Address]: invalid_bech32 - %w", fmt.Errorf("decoding bech32 failed: %w", fmt.Errorf("invalid separator index -1"))),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd := cmd.NewRootCmd(nil)

			rootCmd.SetArgs([]string{"transform", tc.bech32Address, tc.newPrefix})

			b := bytes.NewBufferString("")
			rootCmd.SetOut(b)

			err := rootCmd.Execute()

			if tc.expectedErr != nil {
				require.NotNil(t, err)
				require.Equal(t, tc.expectedErr.Error(), err.Error())
				return
			}
			require.NoError(t, err)

			out, err := ioutil.ReadAll(b)
			require.NoError(t, err)

			trimmedOut := strings.TrimSuffix(string(out), "\n")
			require.Equal(t, tc.expectedOut, trimmedOut)
		})
	}
}
