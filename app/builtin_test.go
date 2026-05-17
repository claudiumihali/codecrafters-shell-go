package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExit(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	fmt.Fprintf(out, "%c ", sigil)

	fmt.Fprintf(in, "invalid_grape_command\n")
	fmt.Fprintf(out, "invalid_grape_command: command not found\n")
	fmt.Fprintf(out, "%c ", sigil)

	fmt.Fprintf(in, "exit\n")
	fmt.Fprintf(out, "%c ", sigil)

	exited := false

	simulateShell(t, in.String(), out.String(), func() { exited = true })

	if !exited {
		t.Fatal("expected exit")
	}
}

func TestEcho(t *testing.T) {
	in := &strings.Builder{}
	out := &strings.Builder{}

	fmt.Fprintf(out, "%c ", sigil)

	fmt.Fprintf(in, "echo mango grape\n")
	fmt.Fprintf(out, "mango grape\n")
	fmt.Fprintf(out, "%c ", sigil)

	fmt.Fprintf(in, "echo grape orange raspberry\n")
	fmt.Fprintf(out, "grape orange raspberry\n")
	fmt.Fprintf(out, "%c ", sigil)

	fmt.Fprintf(in, "echo banana mango apple\n")
	fmt.Fprintf(out, "banana mango apple\n")
	fmt.Fprintf(out, "%c ", sigil)

	simulateShell(t, in.String(), out.String(), unexpectedExit(t))
}
