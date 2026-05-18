package builtin

import (
	"io"
)

type builtinCmdRunF func(args []string, in io.Reader, out io.Writer, exitF func())

var builtinCmds map[string]builtinCmdRunF

func init() {
	builtinCmds = map[string]builtinCmdRunF{
		"exit": func(args []string, in io.Reader, out io.Writer, exitF func()) {
			exit(exitF)
		},
		"echo": func(args []string, in io.Reader, out io.Writer, exitF func()) {
			echo(args, out)
		},
		"type": func(args []string, in io.Reader, out io.Writer, exitF func()) {
			type_(args, out)
		},
	}
}

func Run(args []string, in io.Reader, out io.Writer, exitF func()) bool {
	f, found := builtinCmds[args[0]]
	if !found {
		return false
	}

	f(args, in, out, exitF)

	return true
}
