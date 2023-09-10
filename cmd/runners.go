package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

type Runner func(cmd *cobra.Command, args []string)

type RunnerErr func(cmd *cobra.Command, args []string) error

func CopyClipboardRunErr(reader io.Reader, writer io.Writer) RunnerErr {
	return func(cmd *cobra.Command, args []string) error {
		hasFlag := cmd.Flags().Lookup(flagCopy) != nil
		if !hasFlag {
			return nil
		}

		val, err := cmd.Flags().GetBool(flagCopy)
		if err != nil {
			return fmt.Errorf("reading flag %q: %w", flagCopy, err)
		}
		if !val {
			return nil
		}

		bytes, err := io.ReadAll(reader)
		if err != nil {
			return fmt.Errorf("reading output: %w", err)
		}

		_, err = writer.Write(bytes)
		if err != nil {
			return fmt.Errorf("writing to clipboard: %w", err)
		}

		return nil
	}
}
