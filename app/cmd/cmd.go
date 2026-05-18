package cmd

import (
	"fmt"
	"io"
)

type Cmd struct {
	Environ map[string]string
	Args    []string
	In      io.Reader
	Out     io.Writer
	ExitF   func()
}

var builtinCmds map[string]func(Cmd)

func init() {
	builtinCmds = map[string]func(Cmd){
		"exit": exit,
		"echo": echo,
		"type": type_,
	}
}

func (cmd Cmd) Run() {
	if len(cmd.Args) == 0 {
		return
	}

	f, found := builtinCmds[cmd.Args[0]]
	if found {
		f(cmd)
		return
	}

	fmt.Fprintf(cmd.Out, "%s: command not found\n", cmd.Args[0])
}
