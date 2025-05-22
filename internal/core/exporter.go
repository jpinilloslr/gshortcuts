package core

import (
	"fmt"
	"github.com/fatih/color"
)

type Exporter struct {
	codec   *ShortcutCodec
	manager *ShortcutManager
}

func NewExporter() *Exporter {
	return &Exporter{
		codec:   NewShortcutCodec(),
		manager: NewShortcutManager(),
	}
}

func (i *Exporter) Export(fileName string, verbose bool) error {
	shortcuts, err := i.manager.GetAll()
	if err != nil {
		return err
	}

	if err := i.codec.Encode(shortcuts, fileName); err != nil {
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
