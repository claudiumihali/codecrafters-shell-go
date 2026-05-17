package main

import (
	"fmt"
	"io"
	"strings"
)

func builtin(words []string, w io.Writer, exitF func()) bool {
	switch words[0] {
	case "exit":
		exitF()
	case "echo":
		fmt.Fprintf(w, "%s\n", strings.Join(words[1:], " "))
	default:
		return false
	}
	return true
}
