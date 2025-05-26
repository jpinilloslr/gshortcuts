package core

import (
	"fmt"
)

type ConflictChecker struct {
	manager *ShortcutManager
}

func NewConflictChecker() *ConflictChecker {
	return &ConflictChecker{
		manager: NewShortcutManager(),
	}
}

func (dc *ConflictChecker) Check() error {
	dups, err := dc.getDuplicates()
	if err != nil {
		return err
	}

	if len(dups) == 0 {
		fmt.Println("No duplicate shortcuts found.")
		return nil
	}

	for binding, entries := range dups {
		fmt.Printf("Duplicate shortcut: %s\n", binding)
		for _, entry := range entries {
			fmt.Printf("  - %s\n", entry)
		}
	}

	return nil
}

func (dc *ConflictChecker) getDuplicates() (map[string][]string, error) {
	builtInShortcuts := dc.manager.GetBuiltInShortcuts(false)

	customShortcuts, err := dc.manager.GetCustomShortcuts()
	if err != nil {
		return nil, err
	}

	data := map[string][]string{}

	for schema, shortcuts := range builtInShortcuts {
		for _, shortcut := range shortcuts {
			for _, binding := range shortcut.Bindings {
				data[binding] = append(
					data[binding],
					fmt.Sprintf("%s.%s", schema, shortcut.Key),
				)
			}
		}
	}

	for _, shortcut := range customShortcuts {
		data[shortcut.Binding] = append(
			data[shortcut.Binding],
			fmt.Sprintf("custom.%s", shortcut.Id),
		)
	}

	dupBindings := map[string][]string{}

	for binding, dups := range data {
		if len(dups) > 1 {
			dupBindings[binding] = dups
		}
	}

	return dupBindings, nil
}
