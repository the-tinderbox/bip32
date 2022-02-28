package run

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
)

func Decode(cmd *cobra.Command, args []string) error {
	prompt, err := getPromptStatus()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	if prompt {
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter key: "); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	}

	inputReader := bufio.NewReader(cmd.InOrStdin())
	keyString, err := inputReader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read key from input: %w", err)
	}
	keyString = strings.Trim(keyString, "\n")

	key, err := bip32.B58Deserialize(keyString)
	if err != nil {
		return fmt.Errorf("failed to decode key: %w", err)
	}

	jb, err := json.MarshalIndent(key, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize key: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), string(jb)); err != nil {
		return fmt.Errorf("failed to write serialized key to output: %w", err)
	}

	return nil
}
