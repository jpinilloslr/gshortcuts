package core

import "fmt"

type ImportStrategy string

const (
	Merge   ImportStrategy = "merge"
	Replace ImportStrategy = "replace"
)

var validOptions = map[string]ImportStrategy{
	"merge":   Merge,
	"replace": Replace,
}

func ParseImportStrategy(s string) (ImportStrategy, error) {
	if t, ok := validOptions[s]; ok {
		return t, nil
	}
	return "", fmt.Errorf("invalid import strategy: %s", s)
}
