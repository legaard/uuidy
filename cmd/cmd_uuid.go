package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

func V1Cmd() *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyNumberFlag(),
		)
		cmd = &cobra.Command{
			Use:     "v1",
			Short:   "Generate a UUID V1",
			Long:    "UUID based on the current timestamp and MAC address",
			Example: "uuid v1",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(FlagNumber)
				if err != nil {
					return err
				}

				return writeMany(int(number), cmd.OutOrStdout(), func() (string, error) {
					value, genErr := uuid.NewV1()
					if genErr != nil {
						return "", fmt.Errorf("generating UUID: %w", genErr)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V3Cmd(defaultNamespace uuid.UUID) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyNumberFlag(),
			ApplyNamespaceFlag(defaultNamespace.String()),
		)
		cmd = &cobra.Command{
			Use:     "v3 [value]",
			Short:   "Generate a UUID V3",
			Long:    "UUID based on the MD5 hash of the namespace UUID and name",
			Example: `uuid v3 "Hello v3"`,
			Args:    cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(FlagNumber)
				if err != nil {
					return err
				}

				namespace, err := cmd.Flags().GetString(FlagNamespace)
				if err != nil {
					return err
				}

				ns, err := uuid.FromString(namespace)
				if err != nil {
					return fmt.Errorf("invalid namespace: %w", err)
				}

				return writeMany(int(number), cmd.OutOrStdout(), func() (string, error) {
					return uuid.NewV3(ns, args[0]).String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V4Cmd() *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyNumberFlag(),
		)
		cmd = &cobra.Command{
			Use:     "v4",
			Short:   "Generate a UUID V4",
			Long:    "Randomly generated UUID",
			Example: "uuid v4",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(FlagNumber)
				if err != nil {
					return err
				}

				return writeMany(int(number), cmd.OutOrStdout(), func() (string, error) {
					value, genErr := uuid.NewV4()
					if genErr != nil {
						return "", fmt.Errorf("generating UUID: %w", genErr)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V5Cmd(defaultNamespace uuid.UUID) *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyNumberFlag(),
			ApplyNamespaceFlag(defaultNamespace.String()),
		)
		cmd = &cobra.Command{
			Use:     "v5 [value]",
			Short:   "Generate a UUID V5",
			Long:    "UUID based on SHA-1 hash of the namespace UUID and value",
			Example: `uuid v5 "Hello v5"`,
			Args:    cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(FlagNumber)
				if err != nil {
					return err
				}

				namespace, err := cmd.Flags().GetString(FlagNamespace)
				if err != nil {
					return err
				}

				ns, err := uuid.FromString(namespace)
				if err != nil {
					return fmt.Errorf("invalid namespace: %w", err)
				}

				return writeMany(int(number), cmd.OutOrStdout(), func() (string, error) {
					return uuid.NewV5(ns, args[0]).String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V6Cmd() *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyNumberFlag(),
		)
		cmd = &cobra.Command{
			Use:     "v6",
			Short:   "Generate a UUID V6",
			Long:    "K-sortable UUID based on a timestamp and 48 bits of pseudorandom data",
			Example: "uuid v6",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(FlagNumber)
				if err != nil {
					return err
				}

				return writeMany(int(number), cmd.OutOrStdout(), func() (string, error) {
					value, genErr := uuid.NewV6()
					if genErr != nil {
						return "", fmt.Errorf("generating UUID: %w", genErr)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func V7Cmd() *cobra.Command {
	var (
		applyFlags = MergeAppliers(
			ApplyNumberFlag(),
			ApplyEpocTime(),
		)
		cmd = &cobra.Command{
			Use:     "v7",
			Short:   "Generate a UUID V7",
			Long:    "K-sortable UUID based on the current millisecond precision",
			Example: "uuid v7",
			RunE: func(cmd *cobra.Command, args []string) error {
				number, err := cmd.Flags().GetUint32(FlagNumber)
				if err != nil {
					return err
				}

				epochStr, err := cmd.Flags().GetString(FlagEpoch)
				if err != nil {
					return err
				}

				epoch, err := time.Parse(time.RFC3339Nano, epochStr)
				if err != nil {
					return fmt.Errorf("invalid epoch format: %w", err)
				}

				return writeMany(int(number), cmd.OutOrStdout(), func() (string, error) {
					value, genErr := uuid.NewV7AtTime(epoch)
					if genErr != nil {
						return "", fmt.Errorf("generating UUID: %w", genErr)
					}

					return value.String(), nil
				})
			},
		}
	)

	applyFlags(cmd)

	return cmd
}

func ParseCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "parse [value]",
		Short:   "Parse a UUID value",
		Long:    "Parses UUID value and outputs version details",
		Example: "uuid parse 01ebb00e-d38a-11ef-8f83-426648c33d81",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			value, err := uuid.FromString(args[0])
			if err != nil {
				return err
			}

			switch value.Version() {
			case 1:
				v1, _ := uuid.TimestampFromV1(value)
				ts, _ := v1.Time()

				cmd.Printf("version: %v\n", value.Version())
				cmd.Printf("time: %s\n", ts.Format(time.RFC3339Nano))
			case 3:
				cmd.Printf("version: %v\n", value.Version())
			case 4:
				cmd.Printf("version: %v\n", value.Version())
			case 5:
				cmd.Printf("version: %v\n", value.Version())
			case 6:
				v1, _ := uuid.TimestampFromV6(value)
				ts, _ := v1.Time()

				cmd.Printf("version: %v\n", value.Version())
				cmd.Printf("time: %s\n", ts.Format(time.RFC3339Nano))
			case 7:
				v1, _ := uuid.TimestampFromV7(value)
				ts, _ := v1.Time()

				cmd.Printf("version: %v\n", value.Version())
				cmd.Printf("time: %s\n", ts.Format(time.RFC3339Nano))
			}

			return nil
		},
	}
}

func writeMany(number int, writer io.Writer, generatorFunc func() (string, error)) error {
	var sep = "\n"

	for i := 0; i < number; i++ {
		var (
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
