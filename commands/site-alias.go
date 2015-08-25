package commands

import (
  "fmt"
  "os"
	"io/ioutil"
	"strings"
  "github.com/spf13/cobra"
  "github.com/mrf/grush/utils"
)

var siteAliasCmd = &cobra.Command{
  Use: "site-alias",
  Aliases: []string{"sa"},
  Short: "Print site alias records for all known site aliases and local sites.",
  Run: func(cmd *cobra.Command, args []string) {
    siteAlias()
  },
}

func init() {
  GrushCmd.AddCommand(siteAliasCmd)
}

func siteAlias() {
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
