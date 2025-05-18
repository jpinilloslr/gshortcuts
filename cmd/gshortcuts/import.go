package main

import (
	"log/slog"

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
	rootCmd.AddCommand(importCmd)
}

func importShortcuts(fileName string) error {
	importer := core.NewImporter(slog.Default())
	return importer.ImportFromJson(fileName)
}
