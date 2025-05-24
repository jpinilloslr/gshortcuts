package core

import (
	"fmt"
	"github.com/fatih/color"
)

type Exporter struct {
	codec   *ShortcutsCodec
	manager *ShortcutManager
}

func NewExporter() *Exporter {
	return &Exporter{
		codec:   NewShortcutsCodec(),
		manager: NewShortcutManager(),
	}
}

func (i *Exporter) Export(fileName string, verbose bool) error {
	builtInShortcuts, err := i.manager.GetBuiltInShortcuts(true)
	if err != nil {
		return err
	}

	customShortcuts, err := i.manager.GetCustomShortcuts()
	if err != nil {
		return err
	}

	data := Shortcuts{
		BuiltIn: builtInShortcuts,
		Custom:  customShortcuts,
	}

	if err := i.codec.Encode(&data, fileName); err != nil {
		return err
	}

	totalCount := len(customShortcuts)

	for schema, shortcuts := range builtInShortcuts {
		totalCount += len(shortcuts)
		if verbose {
			fmt.Printf("Exported %d shortcuts in \"%s\"\n", len(shortcuts), schema)
			for _, shortcut := range shortcuts {
				fmt.Printf("\t%s: %+v\n", shortcut.Key, shortcut.Bindings)
			}
			fmt.Println()
		}
	}

	if verbose {
		fmt.Printf("Exported %d custom shortcuts\n", len(customShortcuts))
		for _, shortcut := range customShortcuts {
			fmt.Printf("\t%s: %s\n", shortcut.Id, shortcut.Binding)
		}
		fmt.Println()
	}

	fmt.Printf("%s Exported %d total shortcuts to %s\n",
		color.GreenString("✔"), totalCount, fileName)

	return nil
}
