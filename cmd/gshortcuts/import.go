package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jpinilloslr/gshortcuts/internal/core"
	"github.com/spf13/cobra"
)

var strategy string
var strategies = []string{"merge", "override"}

var importCmd = &cobra.Command{
	Use:   "import [filename]",
	Short: "Import shortcuts from a file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !slices.Contains(strategies, strategy) {
			return fmt.Errorf(
				"Invalid strategy: %s. Valid strategies are: %s",
				strategy,
				strings.Join(strategies, ", "),
			)
		}
		return importShortcuts(args[0], strategy)
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
	rootCmd.AddCommand(importCmd)
}

func importShortcuts(fileName, strategy string) error {
	importer := core.NewImporter()
	return importer.Import(fileName, strategy)
}
