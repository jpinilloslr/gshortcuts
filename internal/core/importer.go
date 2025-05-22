package core

import (
	"fmt"

	"github.com/fatih/color"
)

type Importer struct {
	codec   *ShortcutCodec
	manager *ShortcutManager
}

func NewImporter() *Importer {
	return &Importer{
		codec:   NewShortcutCodec(),
		manager: NewShortcutManager(),
	}
}

func (i *Importer) Import(fileName string, strategy ImportStrategy, verbose bool) error {
	shortcuts, err := i.codec.Decode(fileName)
	if err != nil {
		return err
	}

	if strategy == Replace {
		if err := i.manager.DeleteAll(); err != nil {
			return err
		}
		fmt.Println("Deleted all existing shortcuts")
	}

	for _, shortcut := range shortcuts {
		if err := i.manager.Set(&shortcut); err != nil {
			return err
		}
		if verbose {
			fmt.Printf("Imported shortcut: %s\n", shortcut.Name)
			fmt.Printf("\tCommand: %s\n", shortcut.Command)
			fmt.Printf("\tBinding: %s\n", shortcut.Binding)
		}
	}

	if verbose {
		fmt.Println()
	}

	fmt.Printf("%s Imported %d shortcuts from %s\n",
		color.GreenString("✔"), len(shortcuts), fileName)

	return nil
}

func (i *Importer) Reset() error {
	if err := i.manager.DeleteAll(); err != nil {
		return err
	}
	fmt.Println("Deleted all existing shortcuts")
	return nil
}
