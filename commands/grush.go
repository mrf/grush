package commands

import (
  "github.com/spf13/cobra"
)

var GrushCmd = &cobra.Command{
  Use: "grush",
  Short: "grush is drush in go",
  Long: "Grush is a command line shell and Unix scripting interface for Drupal",
  Run: func(cmd *cobra.Command, args []string) {
    // Init Config
  },
}

func Execute() {
  GrushCmd.Execute()
}
