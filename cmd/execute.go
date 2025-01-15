package cmd

import (
	"os"

	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

var (
	cliVersion = "0.0.2"

	groupUUID = "UUID"

	defaultNamespace     = uuid.NamespaceDNS
	defaultUUIDGenerator = uuid.NewV4
)

func Execute() error {
	var (
		writer  = os.Stdout
		root    = RootCmd(writer, defaultUUIDGenerator)
		version = VersionCmd(cliVersion)
		v1      = V1Cmd()
		v3      = V3Cmd(defaultNamespace)
		v4      = V4Cmd()
		v5      = V5Cmd(defaultNamespace)
		v6      = V6Cmd()
		v7      = V7Cmd()
		parse   = Parse()
	)

	v1.GroupID = groupUUID
	v3.GroupID = groupUUID
	v4.GroupID = groupUUID
	v5.GroupID = groupUUID
	v6.GroupID = groupUUID
	v7.GroupID = groupUUID
	parse.GroupID = groupUUID

	root.AddGroup(&cobra.Group{ID: groupUUID, Title: "UUIDs"})
	root.AddCommand(version, v1, v3, v4, v5, v6, v7, parse)

	return root.Execute()
}
