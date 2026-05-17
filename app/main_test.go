package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func simulateShell(t *testing.T, in string, expectedOut string) {
	inBuf := bytes.NewBufferString(in)
	outBuf := &bytes.Buffer{}
	err := run(t.Context(), nil, inBuf, outBuf)
	if err != nil {
		t.Fatalf("error run: %v", err)
	}

	if outBuf.String() != expectedOut {
		t.Fatalf("wrong output: expected %s; actual %s", expectedOut,
			outBuf.String())
	}
}

func TestPrompt(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	fmt.Fprintf(out, "%c ", sigil)

	simulateShell(t, in.String(), out.String())
}

func TestInvalidCommand(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	fmt.Fprintf(out, "%c ", sigil)

	fmt.Fprintf(in, "invalid_raspberry_command\n")
	fmt.Fprintf(out, "invalid_raspberry_command: command not found\n")
	fmt.Fprintf(out, "%c ", sigil)

	simulateShell(t, in.String(), out.String())
}

func TestRepl(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	fmt.Fprintf(out, "%c ", sigil)

	for i := range 5 {
		fmt.Fprintf(in, "invalid_command_%d\n", i+1)
		fmt.Fprintf(out, "invalid_command_%d: command not found\n", i+1)
		fmt.Fprintf(out, "%c ", sigil)
	}

	simulateShell(t, in.String(), out.String())
}
