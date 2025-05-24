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

func (i *Importer) Import(fileName string, strategy ImportStrategy, verbose bool) error {
	shortcuts, err := i.codec.Decode(fileName)
	if err != nil {
		return err
	}

	if strategy == Replace {
		if !confirm("This will delete all existing shortcuts. Do you want to continue?") {
			return fmt.Errorf("Aborded")
		}

		if err := i.manager.ResetCustomShortcuts(); err != nil {
			return err
		}
		fmt.Println("Deleted all existing shortcuts")
	}

	for _, shortcut := range shortcuts.Custom {
		if err := i.manager.SetCustomShortcut(&shortcut); err != nil {
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
		color.GreenString("✔"), len(shortcuts.Custom), fileName)

	return nil
}

func (i *Importer) Reset() error {
	if !confirm("This will delete all existing shortcuts. Do you want to continue?") {
		return fmt.Errorf("Aborded")
	}

	if err := i.manager.ResetCustomShortcuts(); err != nil {
		return err
	}
	fmt.Println("Deleted all existing shortcuts")
	return nil
}
