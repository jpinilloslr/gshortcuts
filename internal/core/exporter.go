package core

import "fmt"

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

func (i *Exporter) Export(fileName string) error {
	shortcuts, err := i.manager.GetAll()
	if err != nil {
		return err
	}

	if err := i.codec.Encode(shortcuts, fileName); err != nil {
		return err
	}

	for _, shortcut := range shortcuts {
		fmt.Printf("Exported shortcut: %s\n", shortcut.Name)
		fmt.Printf("\tCommand: %s\n", shortcut.Command)
		fmt.Printf("\tBinding: %s\n", shortcut.Binding)
	}

	return nil
}
