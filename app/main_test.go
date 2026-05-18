package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func simulateShell(t *testing.T, in string, expectedOut string, exitF func()) {
	inBuf := bytes.NewBufferString(in)
	outBuf := &bytes.Buffer{}
	err := run(inBuf, outBuf, exitF)
	if err != nil {
		t.Fatalf("error run: %v", err)
	}

	if outBuf.String() != expectedOut {
		t.Fatalf("wrong output: expected %s; actual %s", expectedOut,
			outBuf.String())
	}
}

func unexpectedExit(t *testing.T) func() {
	return func() {
		t.Fatal("unexpected exit")
	}
}

func TestPrompt(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	simulateShell(t, in.String(), out.String(), unexpectedExit(t))
}

func TestInvalidCommand(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	fmt.Fprintf(in, "invalid_raspberry_command\n")
	fmt.Fprintf(out, "invalid_raspberry_command: command not found\n")
	prompt(out)

	simulateShell(t, in.String(), out.String(), unexpectedExit(t))
}

func TestRepl(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	prompt(out)

	for i := range 5 {
		fmt.Fprintf(in, "invalid_command_%d\n", i+1)
		fmt.Fprintf(out, "invalid_command_%d: command not found\n", i+1)
		prompt(out)
	}

	simulateShell(t, in.String(), out.String(), unexpectedExit(t))
}
