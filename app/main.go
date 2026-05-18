package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
)

func run(in io.Reader, out io.Writer, exitF func()) error {
	prompt(out)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		cmd := cmd.Cmd{
			Args:  words,
			In:    in,
			Out:   out,
			ExitF: exitF,
		}

		cmd.Run()

		prompt(out)
	}

	return scanner.Err()
}

func main() {
	err := run(os.Stdin, os.Stdout, func() { os.Exit(0) })
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
