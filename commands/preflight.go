package commands

// File is similar to:
// https://github.com/drush-ops/drush/blob/master/includes/preflight.inc

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func PreFlight(cmd *cobra.Command, args []string) {
	InitializeConfig()

	preFlightRoot()
}

func preFlightRoot() {
	// Locate the actual drupal directory given the selectedRoot
	viper.GetString("SelectedRoot")
}
