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
	schema := fmt.Sprintf(
		"org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:%s",
		path,
	)

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
	return value, nil
}

func (s *ShortcutManager) getEntryPath(id string) string {
	return fmt.Sprintf(
		"/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/%s/",
		id)
}
