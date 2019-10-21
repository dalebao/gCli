package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Reader() string {
	fmt.Print("\033[36m gli: \033[0m")
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return ""
	}
	return strings.TrimSuffix(cmdString, "\n")
}

func Print(output string) {
	if output != "" {
		_, _ = fmt.Println(output)
	}
}
