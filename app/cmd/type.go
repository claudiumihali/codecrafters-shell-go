package cmd

import (
	"fmt"
)

func type_(cmd Cmd) {
	if len(cmd.Args) < 2 {
		return
	}

	_, found := builtinCmds[cmd.Args[1]]
	if found {
		fmt.Fprintf(cmd.Out, "%s is a shell builtin\n", cmd.Args[1])
		return
	}

	fmt.Fprintf(cmd.Out, "%s: not found\n", cmd.Args[1])
}
