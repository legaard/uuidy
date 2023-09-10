package cmd

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
	"io"
)

func V1Cmd(writer io.Writer) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyCopyClipboardFlag(),
			ApplyNumberFlag(),
		)
		cmd = &cobra.Command{
			Use:     "v1",
			Short:   "Generate an UUID V1",
			Long:    "UUID based on the current timestamp and MAC address",
			Example: "uuid v1",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(flagNumber)
				if err != nil {
					return err
				}

				return writeMany(int(number), writer, func() (string, error) {
					value, err := uuid.NewV1()
					if err != nil {
						return "", fmt.Errorf("generating UUID: %w", err)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V3Cmd(writer io.Writer, defaultNamespace uuid.UUID) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyCopyClipboardFlag(),
			ApplyNumberFlag(),
			ApplyNamespaceFlag(defaultNamespace.String()),
		)
		cmd = &cobra.Command{
			Use:     "v3 [value]",
			Short:   "Generate an UUID V3",
			Long:    "UUID based on the MD5 hash of the namespace UUID and name",
			Example: `uuid v3 "Hello v3"`,
			Args:    cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(flagNumber)
				if err != nil {
					return err
				}

				namespace, err := cmd.Flags().GetString(flagNamespace)
				if err != nil {
					return err
				}

				ns, err := uuid.FromString(namespace)
				if err != nil {
					return fmt.Errorf("invalid namespace: %w", err)
				}

				return writeMany(int(number), writer, func() (string, error) {
					return uuid.NewV3(ns, args[0]).String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V4Cmd(writer io.Writer) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyCopyClipboardFlag(),
			ApplyNumberFlag(),
		)
		cmd = &cobra.Command{
			Use:     "v4",
			Short:   "Generate an UUID V4",
			Long:    "Randomly generated UUID",
			Example: "uuid v4",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(flagNumber)
				if err != nil {
					return err
				}

				return writeMany(int(number), writer, func() (string, error) {
					value, err := uuid.NewV4()
					if err != nil {
						return "", fmt.Errorf("generating UUID: %w", err)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V5Cmd(writer io.Writer, defaultNamespace uuid.UUID) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyCopyClipboardFlag(),
			ApplyNumberFlag(),
			ApplyNamespaceFlag(defaultNamespace.String()),
		)
		cmd = &cobra.Command{
			Use:     "v5 [value]",
			Short:   "Generate an UUID V5",
			Long:    "UUID based on SHA-1 hash of the namespace UUID and value",
			Example: `uuid v5 "Hello v5"`,
			Args:    cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(flagNumber)
				if err != nil {
					return err
				}

				namespace, err := cmd.Flags().GetString(flagNamespace)
				if err != nil {
					return err
				}

				ns, err := uuid.FromString(namespace)
				if err != nil {
					return fmt.Errorf("invalid namespace: %w", err)
				}

				return writeMany(int(number), writer, func() (string, error) {
					return uuid.NewV5(ns, args[0]).String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V6Cmd(writer io.Writer) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyCopyClipboardFlag(),
			ApplyNumberFlag(),
		)
		cmd = &cobra.Command{
			Use:     "v6",
			Short:   "Generate an UUID V6",
			Long:    "K-sortable UUID based on a timestamp and 48 bits of pseudorandom data",
			Example: "uuid v6",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(flagNumber)
				if err != nil {
					return err
				}

				return writeMany(int(number), writer, func() (string, error) {
					value, err := uuid.NewV6()
					if err != nil {
						return "", fmt.Errorf("generating UUID: %w", err)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V7Cmd(writer io.Writer) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyCopyClipboardFlag(),
			ApplyNumberFlag(),
		)
		cmd = &cobra.Command{
			Use:   "v7",
			Short: "Generate an UUID V7",
			Long:  "K-sortable UUID based on the current millisecond precision",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(flagNumber)
				if err != nil {
					return err
				}

				return writeMany(int(number), writer, func() (string, error) {
					value, err := uuid.NewV7()
					if err != nil {
						return "", fmt.Errorf("generating UUID: %w", err)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func writeMany(number int, writer io.Writer, generatorFunc func() (string, error)) error {
	for i := 0; i < number; i++ {
		var (
			sep        = "\n"
			last       = i+1 == number
			value, err = generatorFunc()
		)
		if err != nil {
			return err
		}

		if last {
			sep = ""
		}

		_, err = writer.Write([]byte(value + sep))
		if err != nil {
			return err
		}
	}

	return nil
}
