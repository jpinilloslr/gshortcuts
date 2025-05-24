package main

import (
	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var exportAll bool

var exportCmd = &cobra.Command{
	Use:   "export [filename]",
	Short: "Export shortcuts to a file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return exportShortcuts(args[0])
	},
}

func init() {
	exportCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	exportCmd.Flags().BoolVarP(
		&exportAll,
		"all",
		"a",
		false,
		"Export all shortcuts, only modified ones will be exported by default",
	)
	rootCmd.AddCommand(exportCmd)
}

func exportShortcuts(fileName string) error {
	exporter := core.NewExporter()
	return exporter.Export(fileName, verbose, !exportAll)
}
