package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func builtin(words []string, w io.Writer) bool {
	switch words[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Fprintf(w, "%s\n", strings.Join(words[1:], " "))
	default:
		return false
	}
	return true
}
