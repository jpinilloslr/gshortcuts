package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type ShortcutsCodec struct {
}

func NewShortcutsCodec() *ShortcutsCodec {
	return &ShortcutsCodec{}
}

func (c *ShortcutsCodec) Decode(fileName string) (*Shortcuts, error) {
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

func (c *ShortcutsCodec) Encode(shortcuts *Shortcuts, fileName string) error {
	ext := strings.ToLower(filepath.Ext(fileName))

	switch ext {
	case ".json":
		return c.encodeJSON(shortcuts, fileName)
	default:
		return c.encodeYAML(shortcuts, fileName)
	}
}

func (l *ShortcutsCodec) decodeJSON(fileName string) (*Shortcuts, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var data Shortcuts
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (l *ShortcutsCodec) encodeJSON(shortcuts *Shortcuts, fileName string) error {
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	err := enc.Encode(shortcuts)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (l *ShortcutsCodec) decodeYAML(fileName string) (*Shortcuts, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var data Shortcuts
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (l *ShortcutsCodec) encodeYAML(shortcuts *Shortcuts, fileName string) error {
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
