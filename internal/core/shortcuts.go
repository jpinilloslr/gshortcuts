package core

type Shortcuts struct {
	BuiltIn map[string][]BuiltInShortcut `json:"builtIn,omitempty"`
	Custom  []CustomShortcut             `json:"custom,omitempty"`
}
