package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
)

func run(
	ctx context.Context, args []string, in io.Reader, out io.Writer,
	exitF func(),
) error {
	prompt(out)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		cmdArgs := strings.Fields(scanner.Text())
		if len(cmdArgs) == 0 {
			prompt(out)
			continue
		}

		if builtin.Run(cmdArgs, in, out, exitF) {
			prompt(out)
			continue
		}

		fmt.Fprintf(out, "%s: command not found\n", cmdArgs[0])
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
