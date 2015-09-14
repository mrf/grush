package commands

import (
	"github.com/spf13/cobra"
)

var siteAliasCmd = &cobra.Command{
	Use:     "site-alias",
	Aliases: []string{"sa"},
	Short:   "Print site alias records for all known site aliases and local sites.",
	Run:     siteAlias,
}

func init() {
	GrushCmd.AddCommand(siteAliasCmd)
}

func siteAlias(cmd *cobra.Command, args []string) {
	// @todo swap this for loading aliases already in config
}
