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
	shortcuts, err := i.manager.GetCustomShortcuts()
	if err != nil {
		return err
	}

	data := Shortcuts{
		Custom: shortcuts,
	}

	if err := i.codec.Encode(&data, fileName); err != nil {
		return err
	}

	if verbose {
		for _, shortcut := range shortcuts {
			fmt.Printf("Exported shortcut: %s\n", shortcut.Name)
			fmt.Printf("\tCommand: %s\n", shortcut.Command)
			fmt.Printf("\tBinding: %s\n", shortcut.Binding)
		}
		fmt.Println()
	}

	fmt.Printf("%s Exported %d shortcuts to %s\n",
		color.GreenString("✔"), len(shortcuts), fileName)

	return nil
}
