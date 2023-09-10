package cmd

import (
	"github.com/spf13/cobra"
)

func VersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("version is: %s\n", version)
		},
	}
}
