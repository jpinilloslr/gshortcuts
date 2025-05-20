package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ShortcutCodec struct {
}

func NewShortcutCodec() *ShortcutCodec {
	return &ShortcutCodec{}
}

func (c *ShortcutCodec) Decode(fileName string) ([]Shortcut, error) {
	extension := strings.ToLower(filepath.Ext(fileName))

	switch extension {
	case ".json":
		return c.decodeJSON(fileName)
	case ".yaml", ".yml":
		return c.decodeYAML(fileName)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", extension)
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
	return nil, fmt.Errorf("yaml format not supported yet")
}
