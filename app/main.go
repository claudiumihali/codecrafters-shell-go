package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func process(line string, w io.Writer, exitF func()) error {
	words := strings.Fields(line)
	if len(words) == 0 {
		return nil
	}

	if builtin(words, w, exitF) {
		return nil
	}

	fmt.Fprintf(w, "%s: command not found\n", words[0])

	return nil
}

func run(
	ctx context.Context, args []string, in io.Reader, out io.Writer,
	exitF func(),
) error {
	logger := slog.New(slog.NewJSONHandler(os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	prompt(out)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		process(scanner.Text(), out, exitF)
		prompt(out)
	}

	return scanner.Err()
}

func main() {
	ctx := context.Background()
	err := run(ctx, os.Args, os.Stdin, os.Stdout, func() { os.Exit(0) })
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
