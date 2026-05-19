package cmd

import (
	"fmt"
	"strings"
)

func echo(cmd Cmd) error {
	fmt.Fprintf(cmd.Out, "%s\n", strings.Join(cmd.Args[1:], " "))
	return nil
}
