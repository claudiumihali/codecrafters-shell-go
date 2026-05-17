package main

import (
	"fmt"
	"io"
)

const sigil = '$'

func prompt(w io.Writer) {
	fmt.Fprintf(w, "%c ", sigil)
}
