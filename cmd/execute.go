package cmd

import (
	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
	"log"
	"uuid/internal/capture"
	"uuid/internal/clipboard"
)

var (
	cliVersion = "0.0.1"

	groupGenerator = "UUID_GENERATOR"

	defaultNamespace     = uuid.NamespaceDNS
	defaultUUIDGenerator = uuid.NewV4
)

func Execute() error {
	var (
		writer          = log.Writer()
		clipboardWriter = clipboard.NewWriter()
		captureWriter   = capture.NewWriter(writer)

		root    = RootCmd(captureWriter, clipboardWriter, defaultUUIDGenerator)
		version = VersionCmd(cliVersion)
		v1      = V1Cmd(captureWriter)
		v3      = V3Cmd(captureWriter, defaultNamespace)
		v4      = V4Cmd(captureWriter)
		v5      = V5Cmd(captureWriter, defaultNamespace)
		v6      = V6Cmd(captureWriter)
		v7      = V7Cmd(captureWriter)
	)

	root.AddCommand(version)

	root.AddGroup(&cobra.Group{ID: groupGenerator, Title: "Generate UUIDs"})
	v1.GroupID = groupGenerator
	v3.GroupID = groupGenerator
	v4.GroupID = groupGenerator
	v5.GroupID = groupGenerator
	v6.GroupID = groupGenerator
	v7.GroupID = groupGenerator
	root.AddCommand(v1, v3, v4, v5, v6, v7)

	return root.Execute()
}
