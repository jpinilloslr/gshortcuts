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

func (c *ShortcutCodec) Encode(shortcuts []Shortcut, fileName string) error {
	ext := strings.ToLower(filepath.Ext(fileName))

	switch ext {
	case ".json":
		return c.encodeJSON(shortcuts, fileName)
	default:
		return c.encodeYAML(shortcuts, fileName)
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

func (l *ShortcutCodec) encodeJSON(shortcuts []Shortcut, fileName string) error {
	data, err := json.MarshalIndent(shortcuts, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
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

func (l *ShortcutCodec) encodeYAML(shortcuts []Shortcut, fileName string) error {
	data, err := yaml.Marshal(shortcuts)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
