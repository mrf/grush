package commands

import (
	"bytes"

	// CLI library
	"github.com/spf13/cobra"
	// Viper is used to store configuration similar to drush_set_context()
	"github.com/spf13/viper"
)

// Command struct details: https://github.com/spf13/cobra/blob/master/command.go
var GrushCmd = &cobra.Command{
	Use:              "grush",
	Short:            "grush is drush in go",
	Long:             "Grush is a command line shell and Unix scripting interface for Drupal",
	PersistentPreRun: PreFlight,
}

// Called by main()
func Execute() {
	GrushCmd.Execute()
}

// Flags added to commands
var Debug, AssumeNo, Simulate, Verbose, AssumeYes bool
var Root, Uri string

func init() {
	// POSIX/GNU-style flags
	// More details: github.com/spf13/pflag
	GrushCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display even more information, including internal messages.")
	GrushCmd.PersistentFlags().BoolVarP(&AssumeNo, "no", "n", false, "Assume 'no' as answer to all prompts.")
	GrushCmd.PersistentFlags().StringVarP(&Root, "root", "r", ".", "Drupal root directory to use (default: current directory).")
	GrushCmd.PersistentFlags().BoolVarP(&Simulate, "simulate", "s", false, "Simulate all relevant actions (don't actually change the system).")
	GrushCmd.PersistentFlags().StringVarP(&Uri, "uri", "l", "", "URI of the drupal site to use (only needed in multisite environments or when running on an alternate port).")
	GrushCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display extra information about the command.")
	GrushCmd.PersistentFlags().BoolVarP(&AssumeYes, "yes", "y", false, "Assume 'yes' as answer to all prompts.")
}

/**
 * Called by PreFlight()
 */
func InitializeConfig() {
	// Viper requires a configuration "file" so we provide an empty one here.
	// @todo: Expand to load grush.yml from standard drush locations, plus
	//        optionally all of the actual drush config files.
	viper.SetConfigType("yaml")
	var yamlConfig = []byte("")
	viper.ReadConfig(bytes.NewBuffer(yamlConfig))

	loadDefaultSettings()
}

/**
 * Set all Grush viper settings to the Flag values
 */
func loadDefaultSettings() {
	viper.SetDefault("Debug", Debug)
	viper.SetDefault("AssumeNo", AssumeNo)
	viper.SetDefault("SelectedRoot", Root)
	viper.SetDefault("Simulate", Simulate)
	viper.SetDefault("Uri", Uri)
	viper.SetDefault("Verbose", Verbose)
	viper.SetDefault("AssumeYes", AssumeYes)
}
