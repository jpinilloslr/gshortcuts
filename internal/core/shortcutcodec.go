package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type ShortcutCodec struct {
}

func NewShortcutCodec() *ShortcutCodec {
	return &ShortcutCodec{}
}

func (c *ShortcutCodec) Decode(fileName string) ([]Shortcut, error) {
	ext := strings.ToLower(filepath.Ext(fileName))

	switch ext {
	case ".json":
		return c.decodeJSON(fileName)
	case ".yaml", ".yml":
		return c.decodeYAML(fileName)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}
}

func (l *ShortcutCodec) decodeJSON(fileName string) ([]Shortcut, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var data []Shortcut
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *ShortcutCodec) decodeYAML(fileName string) ([]Shortcut, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var data []Shortcut
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
