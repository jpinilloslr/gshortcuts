package core

import (
	"fmt"
	"testing"

	"github.com/jpinilloslr/gshortcuts/internal/core"
)

func TestShortcutManager_GetAll(t *testing.T) {
	manager := core.NewShortcutManager()

	items, err := manager.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fmt.Printf("Items: %v\n", items)
}

func TestShortcutManager_Set(t *testing.T) {
	manager := core.NewShortcutManager()

	shortcut := &core.Shortcut{
		Id:      "test",
		Name:    "Test",
		Command: "notify-send 'This is a test'",
		Binding: "<Ctrl><Alt>T",
	}

	err := manager.Set(shortcut)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
