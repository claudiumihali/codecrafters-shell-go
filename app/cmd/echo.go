package cmd

import (
	"fmt"
	"strings"
)

func echo(cmd Cmd) {
	fmt.Fprintf(cmd.Out, "%s\n", strings.Join(cmd.Args[1:], " "))
}
