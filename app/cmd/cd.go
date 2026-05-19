package cmd

import (
	"fmt"
	"os"
)

func cd(cmd Cmd) error {
	if len(cmd.Args) < 2 {
		return nil
	}

	if cmd.Args[1] == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		cmd.Args[1] = homeDir
	}

	err := os.Chdir(cmd.Args[1])
	if err != nil {
		fmt.Fprintf(cmd.Out, "cd: %s: No such file or directory\n", cmd.Args[1])
	}

	return nil
}
