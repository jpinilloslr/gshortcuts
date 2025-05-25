package main

import (
	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var skipConfirmation bool

var resetCustomCmd = &cobra.Command{
	Use:   "reset-custom",
	Short: "Reset custom shortcuts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return resetShortcuts()
	},
}

func init() {
	resetCustomCmd.Flags().BoolVarP(
		&skipConfirmation,
		"assumeyes",
		"y",
		false,
		"Skip confirmation prompt",
	)
	rootCmd.AddCommand(resetCustomCmd)
}

func resetShortcuts() error {
	importer := core.NewImporter()
	return importer.ResetCustomShortcuts(skipConfirmation)
}
