package core

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/jpinilloslr/gshortcuts/internal/console"
	"github.com/jpinilloslr/gshortcuts/internal/gsettings"
)

const customKeyBindings = "custom-keybindings"
const customBaseSchema = "org.gnome.settings-daemon.plugins.media-keys"
const customBasePath = "/org/gnome/settings-daemon/plugins/media-keys"

var builtInSchemas = [...]string{
	"org.gnome.desktop.wm.keybindings",
	"org.gnome.mutter.keybindings",
	"org.gnome.mutter.wayland.keybindings",
	"org.gnome.shell.keybindings",
}

type ShortcutManager struct {
}

func NewShortcutManager() *ShortcutManager {
	return &ShortcutManager{}
}

func (s *ShortcutManager) GetBuiltInShortcuts(
	modifiedOnly bool,
) map[string][]BuiltInShortcut {
	result := map[string][]BuiltInShortcut{}

	for _, schema := range builtInSchemas {
		settings, err := gsettings.New(schema)
		if err != nil {
			console.PrintWarning("Couldn't create gsettings for schema %s: %v", schema, err)
			continue
		}
		defer settings.Close()

		shortcuts, err := s.getBuiltInShortcutsFromSchema(
			settings,
			modifiedOnly,
		)
		if err != nil {
			console.PrintWarning("Couldn't read shortcuts from schema %s: %v\n", schema, err)
			continue
		}

		if len(shortcuts) > 0 {
			result[schema] = shortcuts
		}
	}

	return result
}

func (s *ShortcutManager) SetBuiltInShortcuts(
	schema string,
	shortcuts []BuiltInShortcut,
) int {
	imported := 0
	settings, err := gsettings.New(schema)
	if err != nil {
		console.PrintWarning("Couldn't create gsettings for schema %s: %v", schema, err)
		return imported
	}
	defer settings.Close()

	for _, current := range shortcuts {
		err := settings.SetStringArray(current.Key, current.Bindings)
		if err != nil {
			console.PrintWarning(
				"Couldn't write shortcut %s for schema %s: %v",
				current.Key,
				schema,
				err,
			)
			continue
		}
		imported++
	}

	settings.Sync()
	return imported
}

func (s *ShortcutManager) GetCustomShortcuts() ([]CustomShortcut, error) {
	settings, err := gsettings.New(customBaseSchema)
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

func (s *ShortcutManager) SetCustomShortcuts(shortcuts []CustomShortcut) error {
	settings, err := gsettings.New(customBaseSchema)
	if err != nil {
		return err
	}
	defer settings.Close()

	for _, current := range shortcuts {
		if err := s.setCustomShortcut(settings, &current); err != nil {
			return err
		}
	}

	settings.Sync()
	return nil
}

func (s *ShortcutManager) ResetCustomShortcuts() error {
	settings, err := gsettings.New(customBaseSchema)
	if err != nil {
		return err
	}
	defer settings.Close()

	settings.Reset(customKeyBindings)
	settings.Sync()
	return nil
}

func (s *ShortcutManager) getBuiltInShortcutsFromSchema(
	settings *gsettings.GSettings,
	modifiedOnly bool,
) ([]BuiltInShortcut, error) {
	keys, err := settings.ListKeys()
	if err != nil {
		return nil, err
	}

	shortcuts := []BuiltInShortcut{}
	for _, key := range keys {
		if modifiedOnly {
			mod, err := settings.IsKeyModified(key)
			if err != nil {
				return nil, err
			}
			if !mod {
				continue
			}
		}

		shortcuts = append(shortcuts, BuiltInShortcut{
			Key:      key,
			Bindings: settings.GetStringArray(key),
		})
	}

	return shortcuts, nil
}

func (s *ShortcutManager) setCustomShortcut(
	settings *gsettings.GSettings,
	shortcut *CustomShortcut,
) error {
	newPath := fmt.Sprintf("%s/%s/", customBasePath, shortcut.Id)
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
	return nil
}

func (s *ShortcutManager) setParams(path string, shortcut *CustomShortcut) error {
	schema := fmt.Sprintf("%s.%s", customBaseSchema, "custom-keybinding")

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
	schema := fmt.Sprintf("%s.%s", customBaseSchema, "custom-keybinding")

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
