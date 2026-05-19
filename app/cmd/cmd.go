package cmd

import (
	"fmt"
	"io"
	"os/exec"
)

type Cmd struct {
	Env    []string
	Args   []string
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
	ExitF  func()
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

	execCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	execCmd.Env = cmd.Env
	execCmd.Stdin = cmd.In
	execCmd.Stdout = cmd.Out
	execCmd.Stderr = cmd.ErrOut
	err := execCmd.Run()
	if err == nil {
		return
	}

	fmt.Fprintf(cmd.Out, "%s: command not found\n", cmd.Args[0])
}
