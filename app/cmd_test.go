package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestExit(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	fmt.Fprintf(in, "invalid_grape_command\n")
	fmt.Fprintf(out, "invalid_grape_command: command not found\n")
	prompt(out)

	fmt.Fprintf(in, "exit\n")
	prompt(out)

	exited := false

	simulateShell(t, osParams{
		in:    strings.NewReader(in.String()),
		exitF: func() { exited = true },
	}, out.String())

	if !exited {
		t.Fatal("expected exit")
	}
}

func TestEcho(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	fmt.Fprintf(in, "echo mango grape\n")
	fmt.Fprintf(out, "mango grape\n")
	prompt(out)

	fmt.Fprintf(in, "echo grape orange raspberry\n")
	fmt.Fprintf(out, "grape orange raspberry\n")
	prompt(out)

	fmt.Fprintf(in, "echo banana mango apple\n")
	fmt.Fprintf(out, "banana mango apple\n")
	prompt(out)

	fmt.Fprintf(in, "echo\n")
	fmt.Fprintf(out, "\n")
	prompt(out)

	simulateShell(t, osParams{
		in:    strings.NewReader(in.String()),
		exitF: unexpectedExit(t),
	}, out.String())
}

func TestType(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	fmt.Fprintf(in, "type echo\n")
	fmt.Fprintf(out, "echo is a shell builtin\n")
	prompt(out)

	fmt.Fprintf(in, "type exit\n")
	fmt.Fprintf(out, "exit is a shell builtin\n")
	prompt(out)

	fmt.Fprintf(in, "type type\n")
	fmt.Fprintf(out, "type is a shell builtin\n")
	prompt(out)

	fmt.Fprintf(in, "type invalid_grape_command\n")
	fmt.Fprintf(out, "invalid_grape_command: not found\n")
	prompt(out)

	fmt.Fprintf(in, "type invalid_banana_command\n")
	fmt.Fprintf(out, "invalid_banana_command: not found\n")
	prompt(out)

	fmt.Fprintf(in, "type ls\n")
	fmt.Fprintf(out, "ls is /usr/bin/ls\n")
	prompt(out)

	simulateShell(t, osParams{
		in:    strings.NewReader(in.String()),
		exitF: unexpectedExit(t),
	}, out.String())
}

func TestExec(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	fmt.Fprintf(in, "sleep 1\n")
	prompt(out)

	simulateShell(t, osParams{
		in:    strings.NewReader(in.String()),
		exitF: unexpectedExit(t),
	}, out.String())
}

func TestPwd(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	fmt.Fprintf(in, "pwd\n")
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getwd: %v", err)
	}
	fmt.Fprintf(out, "%s\n", wd)
	prompt(out)

	simulateShell(t, osParams{
		in:    strings.NewReader(in.String()),
		exitF: unexpectedExit(t),
	}, out.String())
}
