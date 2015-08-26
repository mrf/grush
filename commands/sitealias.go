package commands

import (
	"fmt"
	"github.com/mrf/grush/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
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
	aliases, err := ioutil.ReadDir(os.Getenv("HOME") + "/.drush")
	for _, fileInfo := range aliases {
		if strings.Contains(fileInfo.Name(), ".drushrc.php") {
			aliasFile, err := ioutil.ReadFile(os.Getenv("HOME") + "/.drush/" + fileInfo.Name())
			utils.Check(err)
			fmt.Printf("Alias File: %s \n", aliasFile)
		}
	}
	utils.Check(err)
}
