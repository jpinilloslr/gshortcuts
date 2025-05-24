package core

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jpinilloslr/gshortcuts/internal/console"
)

type Importer struct {
	codec   *ShortcutsCodec
	manager *ShortcutManager
}

func NewImporter() *Importer {
	return &Importer{
		codec:   NewShortcutsCodec(),
		manager: NewShortcutManager(),
	}
}

func (i *Importer) Import(fileName string, verbose bool) error {
	shortcuts, err := i.codec.Decode(fileName)
	if err != nil {
		return err
	}

	totalCount := len(shortcuts.Custom)

	for schema, entries := range shortcuts.BuiltIn {
		processedCount := i.manager.SetBuiltInShortcuts(schema, entries)
		totalCount += processedCount

		if verbose {
			fmt.Printf("Imported %d shortcuts in \"%s\"\n", processedCount, schema)
			for _, shortcut := range entries {
				fmt.Printf("  %s: %+v\n", shortcut.Key, shortcut.Bindings)
			}
		}
	}

	if err := i.manager.SetCustomShortcuts(shortcuts.Custom); err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Imported %d custom shortcuts\n", len(shortcuts.Custom))
		for _, shortcut := range shortcuts.Custom {
			fmt.Printf("  %s: %+v\n", shortcut.Id, shortcut.Binding)
		}
	}

	fmt.Printf("%s Imported %d total shortcuts from %s\n",
		color.GreenString("✔"), totalCount, fileName)

	return nil
}

func (i *Importer) ResetCustomShortcuts() error {
	if !console.Confirm("This will delete all existing custom shortcuts. Do you want to continue?") {
		return fmt.Errorf("Aborded")
	}

	if err := i.manager.ResetCustomShortcuts(); err != nil {
		return err
	}
	fmt.Println("Reset all custom shortcuts")
	return nil
}
