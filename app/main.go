package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"unicode"
)

func prompt(w io.Writer) {
	fmt.Fprintf(w, "$ ")
}

func in(r io.Reader) ([]string, error) {
	out := make([]string, 0, 1)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, scanner.Err()
}

func process(in string) error {
	for word := range strings.FieldsFuncSeq(in, unicode.IsSpace) {
		fmt.Printf("%s: command not found\n", word)
		break
	}
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
