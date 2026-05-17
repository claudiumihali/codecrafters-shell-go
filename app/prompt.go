package main

import (
	"fmt"
	"io"
)

func prompt(w io.Writer) {
	fmt.Fprintf(w, "$ ")
}
