package core

import (
	"fmt"
	"os/exec"
	"regexp"
	"slices"
	"strings"
)

type ShortcutManager struct {
}

func NewShortcutManager() *ShortcutManager {
	return &ShortcutManager{}
}

func (s *ShortcutManager) GetAll() ([]Shortcut, error) {
	entries, err := s.getEntries()
	if err != nil {
		return nil, err
	}

	shortcuts := make([]Shortcut, 0, len(entries))
	for _, entry := range entries {
		shortcut, err := s.getShortcut(entry)
		if err != nil {
			return nil, err
		}
		shortcuts = append(shortcuts, *shortcut)
	}

	return shortcuts, nil
}

func (s *ShortcutManager) Set(shortcut *Shortcut) error {
	exists, err := s.exists(shortcut.Id)
	if err != nil {
		return err
	}

	if !exists {
		if err := s.addEntry(shortcut.Id); err != nil {
			return err
		}
	}

	return s.setParams(shortcut)
}

func (s *ShortcutManager) DeleteAll() error {
	return exec.Command(
		"gsettings",
		"reset",
		"org.gnome.settings-daemon.plugins.media-keys",
		"custom-keybindings",
	).Run()
}

func (s *ShortcutManager) getEntries() ([]string, error) {
	out, err := exec.Command(
		"gsettings",
		"get",
		"org.gnome.settings-daemon.plugins.media-keys",
		"custom-keybindings",
	).Output()
	if err != nil {
		return nil, err
	}

	data := strings.TrimSpace(string(out))
	re := regexp.MustCompile(`\[(.*)\]`)
	matches := re.FindStringSubmatch(data)

	if len(matches) < 1 {
		return []string{}, nil
	}

	untrimmed := strings.Split(matches[1], ",")
	items := make([]string, 0, len(untrimmed))
	for _, item := range untrimmed {
		value := strings.TrimSpace(
			strings.Trim(
				strings.TrimSpace(item),
				"'",
			),
		)
		if value != "" {
			items = append(items, value)
		}
	}

	return items, nil
}

func (s *ShortcutManager) addEntry(id string) error {
	path := s.getEntryPath(id)
	items, err := s.getEntries()
	if err != nil {
		return err
	}

	if slices.Contains(items, path) {
		return nil
	}

	items = append(items, path)
	quotedItems := make([]string, 0, len(items))
	for _, item := range items {
		quotedItems = append(quotedItems, "'"+item+"'")
	}

	data := "[" + strings.Join(quotedItems, ", ") + "]"

	return exec.Command(
		"gsettings",
		"set",
		"org.gnome.settings-daemon.plugins.media-keys",
		"custom-keybindings",
		data,
	).Run()
}

func (s *ShortcutManager) exists(id string) (bool, error) {
	path := s.getEntryPath(id)
	items, err := s.getEntries()
	if err != nil {
		return false, err
	}

	if slices.Contains(items, path) {
		return true, nil
	}

	return false, nil
}

func (s *ShortcutManager) setParams(shortcut *Shortcut) error {
	path := s.getEntryPath(shortcut.Id)
	schema := s.getSchema(path)

	for key, val := range map[string]string{
		"name":    shortcut.Name,
		"command": shortcut.Command,
		"binding": shortcut.Binding,
	} {
		if err := exec.Command("gsettings",
			"set", schema, key, val,
		).Run(); err != nil {
			return err
		}
	}
	return nil
}

func (s *ShortcutManager) getShortcut(path string) (*Shortcut, error) {
	schema := s.getSchema(path)

	var (
		name    string
		command string
		binding string
	)

	for key, val := range map[string]*string{
		"name":    &name,
		"command": &command,
		"binding": &binding,
	} {
		value, err := s.getParam(schema, key)
		if err != nil {
			return nil, err
		}
		*val = value
	}

	id, err := s.getIdFromSchema(schema)
	if err != nil {
		return nil, err
	}

	return &Shortcut{
		Id:      id,
		Name:    name,
		Command: command,
		Binding: binding,
	}, nil
}

func (s *ShortcutManager) getParam(schema, key string) (string, error) {
	data, err := exec.Command(
		"gsettings",
		"get",
		schema,
		key,
	).Output()
	if err != nil {
		return "", err
	}

	value := strings.TrimSpace(
		strings.Trim(
			strings.TrimSpace(string(data)),
			"'",
		),
	)

	return s.unquote(value), nil
}

func (s *ShortcutManager) getSchema(path string) string {
	return fmt.Sprintf(
		"org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:%s",
		path,
	)
}

func (s *ShortcutManager) getEntryPath(id string) string {
	return fmt.Sprintf(

		"/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/%s/",
		id)
}

func (s *ShortcutManager) getIdFromSchema(schema string) (string, error) {
	re := regexp.MustCompile(`custom-keybindings/(.*)/$`)
	matches := re.FindStringSubmatch(schema)

	if len(matches) < 1 {
		return "", nil
	}

	return matches[1], nil
}

func (s *ShortcutManager) unquote(value string) string {
	// TODO: We need a proper way to unescape GVariant strings
	return strings.ReplaceAll(value, `\\`, `\`)
}
