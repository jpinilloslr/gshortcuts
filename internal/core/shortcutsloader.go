package core

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type ShortcutsLoader struct {
	logger *slog.Logger
}

func NewShortcutsLoader(logger *slog.Logger) *ShortcutsLoader {
	return &ShortcutsLoader{
		logger: logger,
	}
}

func (l *ShortcutsLoader) LoadJson(fileName string) (*ShortcutsConfig, error) {
	l.logger.Info("Loading actions from file", "file", fileName)
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var data ShortcutsConfig
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (l *ShortcutsLoader) LoadYaml(fileName string) (*ShortcutsConfig, error) {
	return nil, fmt.Errorf("yaml format not supported yet")
}
