package cmd

import (
	"fmt"
	"io"

	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

func RootCmd(
	out io.Writer,
	defaultUUIDFn func() (uuid.UUID, error)) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyNumberFlag(),
		)
	)

	cmd := &cobra.Command{
		Use:   "uuid",
		Short: fmt.Sprintf("CLI for generating UUIDs (default V%d)", uuid.Must(defaultUUIDFn()).Version()),
		RunE: func(cmd *cobra.Command, args []string) error {
			number, err := cmd.Flags().GetUint32(FlagNumber)
			if err != nil {
				return err
			}

			return writeMany(int(number), out, func() (string, error) {
				value, genErr := defaultUUIDFn()
				if genErr != nil {
					return "", fmt.Errorf("generating UUID: %w", genErr)
				}

				return value.String(), nil
			})
		},
	}

	cmd.SetOut(out)
	applyFlags(cmd)

	return cmd
}
