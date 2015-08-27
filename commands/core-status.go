package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mrf/grush/includes"
)

var coreStatusCmd = &cobra.Command{
	Use:     "core-status",
	Aliases: []string{"status", "st"},
	Short:   "Provides a birds-eye view of the current Drupal installation, if any.",
	Run:     coreStatus,
}

func init() {
	GrushCmd.AddCommand(coreStatusCmd)
}

func coreStatus(cmd *cobra.Command, args []string) {
	drupalRoot := includes.LocateDrupalRoot(viper.GetString("SelectedRoot"))
	if drupalRoot == "" {
		// Temporary output
		fmt.Printf("Not a Drupal root.\n")
	} else {
		fmt.Printf("Drupal root : %s\n", drupalRoot)
	}
}
