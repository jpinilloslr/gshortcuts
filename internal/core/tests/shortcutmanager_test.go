package core

import (
	"fmt"
	"testing"

	"github.com/jpinilloslr/gshortcuts/internal/core"
)

func TestShortcutManager_GetCustomShortcuts(t *testing.T) {
	manager := core.NewShortcutManager()

	items, err := manager.GetCustomShortcuts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fmt.Printf("Items: %+v\n", items)
}

func TestShortcutManager_GetBuiltInShortcuts(t *testing.T) {
	manager := core.NewShortcutManager()

	items, err := manager.GetBuiltInShortcuts(true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fmt.Printf("Items: %+v\n", items)
}
