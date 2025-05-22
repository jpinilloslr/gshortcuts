package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func confirm(prompt string) bool {
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
