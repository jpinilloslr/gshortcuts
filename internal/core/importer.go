package core

import "fmt"

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

func (i *Importer) Import(fileName string, strategy ImportStrategy) error {
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
		fmt.Printf("Imported shortcut: %s\n", shortcut.Name)
	}

	return nil
}

func (i *Importer) Reset() error {
	if err := i.manager.DeleteAll(); err != nil {
		return err
	}
	fmt.Println("Deleted all existing shortcuts")
	return nil
}
