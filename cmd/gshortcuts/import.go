package main

import (
	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [filename]",
	Short: "Import shortcuts from a file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return importShortcuts(args[0])
	},
}

func init() {
	importCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.AddCommand(importCmd)
}

func importShortcuts(fileName string) error {
	importer := core.NewImporter()
	return importer.Import(fileName, verbose)
}
