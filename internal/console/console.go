package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Confirm(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(
			"%s%s [y/n]: ",
			color.YellowString("❓"),
			prompt,
		)
		input, _ := reader.ReadString('\n')
		switch strings.ToLower(strings.TrimSpace(input)) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Please type y or n and hit enter.")
		}
	}
}

func PrintWarning(format string, a ...any) {
	fmt.Printf(
		"%s %s\n",
		color.YellowString("⚠️ "),
		fmt.Sprintf(format, a...),
	)
}

func PrintError(format string, a ...any) {
	fmt.Printf(
		"%s %s\n",
		color.RedString("❌ "),
		fmt.Sprintf(format, a...),
	)
}
