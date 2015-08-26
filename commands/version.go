package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	VERSION = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the grush version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Grush Version : %s\n", VERSION)
	},
}

func init() {
	GrushCmd.AddCommand(versionCmd)
}
