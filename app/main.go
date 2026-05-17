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

func process(in string, out io.Writer) error {
	words := strings.Fields(in)
	if len(words) == 0 {
		return nil
	}

	if builtin(words, out) {
		return nil
	}

	fmt.Fprintf(out, "%s: command not found\n", words[0])

	return nil
}

func run(ctx context.Context, args []string) error {
	logger := slog.New(slog.NewJSONHandler(os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	prompt(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process(scanner.Text(), os.Stdout)
		prompt(os.Stdout)
	}

	return scanner.Err()
}

func main() {
	ctx := context.Background()
	err := run(ctx, os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
