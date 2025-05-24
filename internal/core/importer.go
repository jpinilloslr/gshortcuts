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

	if err := i.manager.SetCustomShortcuts(shortcuts.Custom); err != nil {
		return err
	}

	if verbose {
		for _, shortcut := range shortcuts.Custom {
			fmt.Printf("Imported custom shortcut: %s\n", shortcut.Name)
		}
		fmt.Println()
	}

	fmt.Printf("%s Imported %d shortcuts from %s\n",
		color.GreenString("✔"), len(shortcuts.Custom), fileName)

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
