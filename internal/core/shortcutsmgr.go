package core

import (
	"fmt"
	"log/slog"
	"os/exec"
	"regexp"
	"slices"
	"strings"
)

type ShortcutsMgr struct {
	logger *slog.Logger
}

func NewShortcutsMgr(logger *slog.Logger) *ShortcutsMgr {
	return &ShortcutsMgr{
		logger: logger,
	}
}

func (s *ShortcutsMgr) Create(id, name, command, binding string) error {
	exists, err := s.exists(id)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("shortcut with id %s already exists", id)
	}

	if err := s.addEntry(id); err != nil {
		return err
	}
	return s.setParams(id, name, command, binding)
}

func (s *ShortcutsMgr) getEntries() ([]string, error) {
	out, err := exec.Command("gsettings",
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
		items = append(items, strings.TrimSpace(item))
	}

	return items, nil
}

func (s *ShortcutsMgr) addEntry(id string) error {
	path := s.getEntryPath(id)
	items, err := s.getEntries()
	if err != nil {
		return err
	}

	if slices.Contains(items, path) {
		return nil
	}

	items = append(items, "'"+path+"'")
	data := "[" + strings.Join(items, ", ") + "]"

	return exec.Command(
		"gsettings",
		"set",
		"org.gnome.settings-daemon.plugins.media-keys",
		"custom-keybindings",
		data,
	).Run()
}

func (s *ShortcutsMgr) exists(id string) (bool, error) {
	path := s.getEntryPath(id)
	items, err := s.getEntries()
	if err != nil {
		return false, err
	}

	if slices.Contains(items, "'"+path+"'") {
		return true, nil
	}

	return false, nil
}
func (s *ShortcutsMgr) setParams(id, name, command, binding string) error {
	path := s.getEntryPath(id)
	schema := fmt.Sprintf(
		"org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:%s",
		path,
	)
	for key, val := range map[string]string{
		"name":    name,
		"command": command,
		"binding": binding,
	} {
		if err := exec.Command("gsettings",
			"set", schema, key, val,
		).Run(); err != nil {
			return err
		}
	}
	return nil
}

func (s *ShortcutsMgr) getEntryPath(id string) string {
	return fmt.Sprintf(
		"/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/%s/",
		id)
}
