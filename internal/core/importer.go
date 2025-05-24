package core

import (
	"fmt"

	"github.com/fatih/color"
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

	for schema, entries := range shortcuts.BuiltIn {
		if err := i.manager.SetBuiltInShortcuts(schema, entries); err != nil {
			return err
		}
	}

	if err := i.manager.SetCustomShortcuts(shortcuts.Custom); err != nil {
		return err
	}

	totalCount := len(shortcuts.Custom)

	for schema, shortcuts := range shortcuts.BuiltIn {
		totalCount += len(shortcuts)
		if verbose {
			fmt.Printf("Imported %d shortcuts in \"%s\"\n", len(shortcuts), schema)
			for _, shortcut := range shortcuts {
				fmt.Printf("\t%s: %+v\n", shortcut.Key, shortcut.Bindings)
			}
			fmt.Println()
		}
	}

	if verbose {
		fmt.Printf("Imported %d custom shortcuts\n", len(shortcuts.Custom))
		for _, shortcut := range shortcuts.Custom {
			fmt.Printf("\t%s: %s\n", shortcut.Id, shortcut.Binding)
		}
		fmt.Println()
	}

	fmt.Printf("%s Imported %d total shortcuts from %s\n",
		color.GreenString("✔"), totalCount, fileName)

	return nil
}

func (i *Importer) ResetCustomShortcuts() error {
	if !confirm("This will delete all existing custom shortcuts. Do you want to continue?") {
		return fmt.Errorf("Aborded")
	}

	if err := i.manager.ResetCustomShortcuts(); err != nil {
		return err
	}
	fmt.Println("Reset all custom shortcuts")
	return nil
}
