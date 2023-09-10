package cmd

import (
	"github.com/spf13/cobra"
)

const (
	flagCopy      = "copy"
	flagNamespace = "namespace"
	flagNumber    = "number"
)

type FlagApplier func(cmd *cobra.Command)

func ApplyCopyClipboardFlag() FlagApplier {
	return func(cmd *cobra.Command) {
		cmd.Flags().BoolP(
			flagCopy,
			"c",
			false,
			"copy value to clipboard",
		)
	}
}

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

func MergeAppliers(appliers ...FlagApplier) FlagApplier {
	return func(cmd *cobra.Command) {
		for _, a := range appliers {
			a(cmd)
		}
	}
}
