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

func (s *ShortcutManager) Test() error {
	schema := "org.gnome.desktop.wm.keybindings"
	key := "switch-to-workspace-9"
	settings, err := gsettings.New(schema)
	if err != nil {
		return err
	}
	defer settings.Close()

	values := settings.GetStringArray(key)
	fmt.Printf("%s keybindings: %v\n", key, values)

	mod, err := settings.IsKeyModified(schema, key)
	if err != nil {
		return err
	}

	if mod {
		fmt.Printf("The key '%s' is modified.\n", key)
	} else {
		fmt.Printf("The key '%s' is not modified.\n", key)
	}

	return nil
}

func (s *ShortcutManager) GetCustomShortcuts() ([]CustomShortcut, error) {
	settings, err := gsettings.New(baseSchema)
	if err != nil {
		return nil, err
	}
	defer settings.Close()

	paths := settings.GetStringArray(customKeyBindings)

	shortcuts := []CustomShortcut{}
	for _, path := range paths {
		current, err := s.getCustomShortcut(path)
		if err != nil {
			return nil, err
		}
		shortcuts = append(shortcuts, *current)
	}

	return shortcuts, nil
}

func (s *ShortcutManager) SetCustomShortcut(shortcut *CustomShortcut) error {
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

func (s *ShortcutManager) ResetCustomShortcuts() error {
	settings, err := gsettings.New(baseSchema)
	if err != nil {
		return err
	}
	defer settings.Close()

	settings.Reset(customKeyBindings)
	settings.Sync()
	return nil
}

func (s *ShortcutManager) setParams(path string, shortcut *CustomShortcut) error {
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

func (s *ShortcutManager) getCustomShortcut(path string) (*CustomShortcut, error) {
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

	return &CustomShortcut{
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
