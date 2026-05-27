package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/lex"
)

type osParams struct {
	env    []string
	in     io.Reader
	out    io.Writer
	errOut io.Writer
	exitF  func()
}

func run(os osParams) error {
	prompt(os.out)

	scanner := bufio.NewScanner(os.in)
	for scanner.Scan() {
		words := lex.Split(scanner.Text())

		cmd := cmd.Cmd{
			Env:    os.env,
			Args:   words,
			In:     os.in,
			Out:    os.out,
			ErrOut: os.errOut,
			ExitF:  os.exitF,
		}

		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.errOut, "%v\n", err)
		}

		prompt(os.out)
	}

	return scanner.Err()
}

func main() {
	err := run(osParams{
		env:    os.Environ(),
		in:     os.Stdin,
		out:    os.Stdout,
		errOut: os.Stderr,
		exitF:  func() { os.Exit(0) },
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
