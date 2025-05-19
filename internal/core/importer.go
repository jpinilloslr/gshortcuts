package core

import (
	"log/slog"
)

type Importer struct {
	logger  *slog.Logger
	manager *ShortcutsMgr
	loader  *ShortcutsLoader
}

func NewImporter(
	logger *slog.Logger,
) *Importer {
	return &Importer{
		logger:  logger,
		manager: NewShortcutsMgr(logger),
		loader:  NewShortcutsLoader(logger),
	}
}

func (i *Importer) ImportFromJson(fileName string) error {
	config, err := i.loader.LoadJson(fileName)
	if err != nil {
		return err
	}

	for _, shortcut := range config.Shortcuts {
		if err := i.manager.Set(
			shortcut.Id,
			shortcut.Name,
			shortcut.Command,
			shortcut.Binding,
		); err != nil {
			i.logger.Error("Failed to set shortcut", "id", shortcut.Id, "err", err)
			continue
		}
		i.logger.Info("Installed shortcut", "id", shortcut.Id, "binding", shortcut.Binding)
	}

	return nil
}
