package main

import (
	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete all shortcuts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return resetShortcuts()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

func resetShortcuts() error {
	importer := core.NewImporter()
	return importer.Reset()
}
