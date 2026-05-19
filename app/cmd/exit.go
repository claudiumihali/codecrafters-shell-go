package cmd

func exit(cmd Cmd) error {
	cmd.ExitF()
	return nil
}
