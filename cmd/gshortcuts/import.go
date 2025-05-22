package main

import (
	"fmt"
	"strings"

	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var strategy string
var strategies = []string{"merge", "replace"}

var importCmd = &cobra.Command{
	Use:   "import [filename]",
	Short: "Import shortcuts from a file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		strat, err := core.ParseImportStrategy(strategy)
		if err != nil {
			return fmt.Errorf(
				"Invalid strategy: %s. Valid strategies are: %s",
				strategy,
				strings.Join(strategies, ", "),
			)
		}
		return importShortcuts(args[0], strat)
	},
}

func init() {
	importCmd.Flags().StringVarP(
		&strategy,
		"strategy",
		"s",
		strategies[0],
		fmt.Sprintf("Import strategy (%s)", strings.Join(strategies, ", ")),
	)

	importCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.AddCommand(importCmd)
}

func importShortcuts(fileName string, strategy core.ImportStrategy) error {
	importer := core.NewImporter()
	return importer.Import(fileName, strategy, verbose)
}
