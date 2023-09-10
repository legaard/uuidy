package cmd

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
	"io"
)

func RootCmd(
	outputReadWriter io.ReadWriter,
	clipboardWriter io.Writer,
	defaultUUIDFn func() (uuid.UUID, error)) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyCopyClipboardFlag(),
			ApplyNumberFlag(),
		)
	)

	cmd := &cobra.Command{
		Use:   "uuid",
		Short: fmt.Sprintf("CLI for generating UUIDs (default V%d)", uuid.Must(defaultUUIDFn()).Version()),
		RunE: func(cmd *cobra.Command, args []string) error {
			number, err := cmd.Flags().GetUint32(flagNumber)
			if err != nil {
				return err
			}

			return writeMany(int(number), outputReadWriter, func() (string, error) {
				value, err := defaultUUIDFn()
				if err != err {
					return "", fmt.Errorf("generating UUID: %w", err)
				}

				return value.String(), nil
			})
		},
		PersistentPostRunE: CopyClipboardRunErr(outputReadWriter, clipboardWriter),
	}

	applyFlags(cmd)

	return cmd
}
