package cmd

import (
	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

func Execute(version string) error {
	var (
		cliVersion           = version
		defaultNamespace     = uuid.NamespaceDNS
		defaultUUIDGenerator = uuid.NewV4

		root       = RootCmd(defaultUUIDGenerator)
		versionCmd = VersionCmd(cliVersion)
		v1         = V1Cmd()
		v3         = V3Cmd(defaultNamespace)
		v4         = V4Cmd()
		v5         = V5Cmd(defaultNamespace)
		v6         = V6Cmd()
		v7         = V7Cmd()
		parse      = ParseCmd()
		null       = NullCmd()

		uuidGroup = &cobra.Group{
			ID:    "UUID_GROUP",
			Title: "UUIDs Commands",
		}
	)

	v1.GroupID = uuidGroup.ID
	v3.GroupID = uuidGroup.ID
	v4.GroupID = uuidGroup.ID
	v5.GroupID = uuidGroup.ID
	v6.GroupID = uuidGroup.ID
	v7.GroupID = uuidGroup.ID
	parse.GroupID = uuidGroup.ID
	null.GroupID = uuidGroup.ID

	root.AddGroup(uuidGroup)
	root.AddCommand(versionCmd, v1, v3, v4, v5, v6, v7, parse, null)

	return root.Execute()
}
