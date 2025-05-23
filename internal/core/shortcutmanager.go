package core

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/jpinilloslr/gshortcuts/internal/gsettings"
)

const customKeyBindings = "custom-keybindings"
const baseSchema = "org.gnome.settings-daemon.plugins.media-keys"
const basePath = "/org/gnome/settings-daemon/plugins/media-keys"

type ShortcutManager struct {
}

func NewShortcutManager() *ShortcutManager {
	return &ShortcutManager{}
}

func (s *ShortcutManager) GetAll() ([]Shortcut, error) {
	settings, err := gsettings.New(baseSchema)
	if err != nil {
		return nil, err
	}
	defer settings.Close()

	paths := settings.GetStringArray(customKeyBindings)

	shortcuts := []Shortcut{}
	for _, path := range paths {
		current, err := s.getShortcut(path)
		if err != nil {
			return nil, err
		}
		shortcuts = append(shortcuts, *current)
	}

	return shortcuts, nil
}

func (s *ShortcutManager) Set(shortcut *Shortcut) error {
	settings, err := gsettings.New(baseSchema)
	if err != nil {
		return err
	}
	defer settings.Close()

	newPath := fmt.Sprintf("%s/%s/", basePath, shortcut.Id)
	paths := settings.GetStringArray(customKeyBindings)

	if !slices.Contains(paths, newPath) {
		paths = append(paths, newPath)

		if err := settings.SetStringArray(customKeyBindings, paths); err != nil {
			return err
		}
	}

	if err := s.setParams(newPath, shortcut); err != nil {
		return err
	}

	settings.Sync()
	return nil
}

func (s *ShortcutManager) DeleteAll() error {
	settings, err := gsettings.New(baseSchema)
	if err != nil {
		return err
	}
	defer settings.Close()

	settings.Reset(customKeyBindings)
	settings.Sync()
	return nil
}

func (s *ShortcutManager) setParams(path string, shortcut *Shortcut) error {
	schema := fmt.Sprintf("%s.%s", baseSchema, "custom-keybinding")

	settings, err := gsettings.NewWithPath(schema, path)
	if err != nil {
		return err
	}
	defer settings.Close()

	if err := settings.SetString("name", shortcut.Name); err != nil {
		return err
	}

	if err := settings.SetString("command", shortcut.Command); err != nil {
		return err
	}

	if err := settings.SetString("binding", shortcut.Binding); err != nil {
		return err
	}

	return nil
}

func (s *ShortcutManager) getShortcut(path string) (*Shortcut, error) {
	schema := fmt.Sprintf("%s.%s", baseSchema, "custom-keybinding")

	settings, err := gsettings.NewWithPath(schema, path)
	if err != nil {
		return nil, err
	}
	defer settings.Close()

	id, err := s.getIdFromPath(path)
	if err != nil {
		return nil, err
	}

	return &Shortcut{
		Id:      id,
		Name:    settings.GetString("name"),
		Command: settings.GetString("command"),
		Binding: settings.GetString("binding"),
	}, nil
}

func (s *ShortcutManager) getIdFromPath(path string) (string, error) {
	re := regexp.MustCompile(`/media-keys/(.*)/$`)
	matches := re.FindStringSubmatch(path)

	if len(matches) < 1 {
		return "", nil
	}

	return matches[1], nil
}
