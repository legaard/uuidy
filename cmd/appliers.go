package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	flagNamespace = "namespace"
	flagNumber    = "number"
	flagEpoch     = "epoch"
)

type FlagApplier func(cmd *cobra.Command)

func ApplyNamespaceFlag(defaultNs string) FlagApplier {
	return func(cmd *cobra.Command) {
		cmd.Flags().String(
			flagNamespace,
			defaultNs,
			"namespace used when generating values",
		)
	}
}

func ApplyNumberFlag() FlagApplier {
	return func(cmd *cobra.Command) {
		cmd.Flags().Uint32P(
			flagNumber,
			"n",
			1,
			"number of values to generate",
		)
	}
}

func ApplyEpocTime() FlagApplier {
	return func(cmd *cobra.Command) {
		cmd.Flags().StringP(
			flagEpoch,
			"e",
			time.Now().Format(time.RFC3339Nano),
			"epoch time used to generate (format: RFC3339 nano)",
		)
	}
}

func MergeAppliers(appliers ...FlagApplier) FlagApplier {
	return func(cmd *cobra.Command) {
		for _, a := range appliers {
			a(cmd)
		}
	}
}
