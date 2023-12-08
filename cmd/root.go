package cmd

import (
	"os"

	"github.com/manyids2/tui-builder/components"
	"github.com/spf13/cobra"
)

var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tui-builder",
	Short: "Interactively build `tview` grids",
	Long: `Interactively build tview grids, by 
	specifying columns and rows in yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		components.RunApp(path)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "./layouts/two-column.yaml", "Path to yaml config")
}
