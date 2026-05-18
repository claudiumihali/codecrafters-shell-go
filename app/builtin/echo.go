package builtin

import (
	"fmt"
	"io"
	"strings"
)

func echo(args []string, out io.Writer) {
	fmt.Fprintf(out, "%s\n", strings.Join(args[1:], " "))
}
