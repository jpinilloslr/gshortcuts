package core

type Shortcut struct {
	Id      string
	Name    string
	Binding string
	Command string
}

type ShortcutsConfig struct {
	UpdateMethod string
	Shortcuts    []Shortcut
}
