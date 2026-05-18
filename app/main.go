package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
)

type osParams struct {
	environ []string
	in      io.Reader
	out     io.Writer
	exitF   func()
}

func run(os osParams) error {
	prompt(os.out)

	scanner := bufio.NewScanner(os.in)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		cmd := cmd.Cmd{
			Environ: os.environ,
			Args:    words,
			In:      os.in,
			Out:     os.out,
			ExitF:   os.exitF,
		}

		cmd.Run()

		prompt(os.out)
	}

	return scanner.Err()
}

func main() {
	err := run(osParams{
		environ: os.Environ(),
		in:      os.Stdin,
		out:     os.Stdout,
		exitF:   func() { os.Exit(0) },
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
