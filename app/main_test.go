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
	invalidCommand := "invalid_raspberry_command"

	in := &strings.Builder{}
	out := &strings.Builder{}

	fmt.Fprintf(out, "%c ", sigil)

	fmt.Fprintf(in, "%s\n", invalidCommand)
	fmt.Fprintf(out, "%s: command not found\n", invalidCommand)
	fmt.Fprintf(out, "%c ", sigil)

	simulateShell(t, in.String(), out.String())
}
