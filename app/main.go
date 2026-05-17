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

func builtinCommand(words []string, out io.Writer) bool {
	switch words[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Fprintf(out, "%s\n", strings.Join(words[1:], " "))
	default:
		return false
	}
	return true
}

func process(in string, out io.Writer) error {
	words := strings.Fields(in)
	if len(words) == 0 {
		return nil
	}

	if builtinCommand(words, out) {
		return nil
	}

	fmt.Fprintf(out, "%s: command not found\n", words[0])

	return nil
}

func main() {
	prompt(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process(scanner.Text(), os.Stdout)
		prompt(os.Stdout)
	}
	err := scanner.Err()
	if err != nil {
		slog.Error("error reading input", "err", err)
		os.Exit(1)
	}
}
