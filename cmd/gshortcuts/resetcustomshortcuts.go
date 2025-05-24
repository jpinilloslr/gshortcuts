package main

import (
	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var resetCustomCmd = &cobra.Command{
	Use:   "reset-custom",
	Short: "Reset custom shortcuts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return resetShortcuts()
	},
}

func init() {
	rootCmd.AddCommand(resetCustomCmd)
}

func resetShortcuts() error {
	importer := core.NewImporter()
	return importer.ResetCustomShortcuts()
}
