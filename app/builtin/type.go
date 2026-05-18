package builtin

import (
	"fmt"
	"io"
)

func type_(args []string, out io.Writer) {
	if len(args) < 2 {
		return
	}

	_, found := builtinCmds[args[1]]
	if found {
		fmt.Fprintf(out, "%s is a shell builtin\n", args[1])
		return
	}

	fmt.Fprintf(out, "%s: not found\n", args[1])
}
