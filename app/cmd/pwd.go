package cmd

import (
	"fmt"
	"os"
)

func pwd(cmd Cmd) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Fprintf(cmd.Out, "%s\n", wd)
	return nil
}
