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

type builtinCmdRunF func(Cmd) error

var builtinCmds map[string]builtinCmdRunF

// init is required here because the [type_] function refers to the
// [builtinCmds] map, so the map cannot be initialized and defined in a single
// statement.
func init() {
	builtinCmds = map[string]builtinCmdRunF{
		"exit": exit,
		"echo": echo,
		"type": type_,
		"pwd":  pwd,
		"cd":   cd,
	}
}

func (cmd Cmd) Run() error {
	if len(cmd.Args) == 0 {
		return nil
	}

	f, found := builtinCmds[cmd.Args[0]]
	if found {
		return f(cmd)
	}

	execCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	execCmd.Env = cmd.Env
	execCmd.Stdin = cmd.In
	execCmd.Stdout = cmd.Out
	execCmd.Stderr = cmd.ErrOut
	err := execCmd.Run()
	if err == nil {
		return nil
	}

	fmt.Fprintf(cmd.Out, "%s: command not found\n", cmd.Args[0])
	return nil
}
