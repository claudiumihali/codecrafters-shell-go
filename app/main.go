package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func prompt(w io.Writer) {
	fmt.Fprintf(w, "$ ")
}

func builtinCommand(word string) bool {
	switch word {
	case "exit":
		os.Exit(0)
	default:
		return false
	}
	return true
}

func process(in string) error {
	words := strings.Fields(in)
	if len(words) == 0 {
		return nil
	}

	if builtinCommand(words[0]) {
		return nil
	}

	fmt.Printf("%s: command not found\n", words[0])

	return nil
}

func main() {
	prompt(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process(scanner.Text())
		prompt(os.Stdout)
	}
	err := scanner.Err()
	if err != nil {
		slog.Error("error reading input", "err", err)
		os.Exit(1)
	}
}
