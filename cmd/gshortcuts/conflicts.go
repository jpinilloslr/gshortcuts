package main

import (
	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var conflictsCmd = &cobra.Command{
	Use:   "conflicts",
	Short: "Check for duplicate shortcuts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return checkForConflicts()
	},
}

func init() {
	rootCmd.AddCommand(conflictsCmd)
}

func checkForConflicts() error {
	checker := core.NewConflictChecker()
	return checker.Check()
}
