package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var colors = []string{
	"01;31",
	"01;32",
	"01;33",
	"01;34",
	"01;35",
	"01;36",
	"01;37",
	"01;91",
	"01;92",
	"01;93",
	"01;94",
	"01;95",
	"01;96",
	"01;97",
}

func main() {
	grepCmdTemplate := "GREP_COLOR='%s' egrep --color=always '%s|$'"
	if len(os.Args) < 3 {
		panic(fmt.Sprintf("Usage: %s filepath pattern[, pattern2[, pattern3[, ...]]]", os.Args[0]))
	}
	args := os.Args[1:]
	filePath := args[0]
	args = args[1:]
	if len(args) > len(colors) {
		fmt.Fprintln(os.Stderr, "More patterns provided than available colors, some might not be highlighted properly")
	}

	cmdParts := make([]string, len(args))
	for i, pattern := range args {
		var color string
		if i < len(colors) {
			color = colors[i]
		}
		cmdParts[i] = fmt.Sprintf(grepCmdTemplate, color, pattern)
	}

	cmdParts[0] += fmt.Sprintf(" %s", filePath)
	cmd := strings.Join(cmdParts, " | ")
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic(fmt.Sprintf("Failed to execute command %v: %s", cmd, err.Error()))
	}
	fmt.Println(string(out))

}
