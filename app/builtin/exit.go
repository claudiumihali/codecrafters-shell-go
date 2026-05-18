package builtin

func exit(exitF func()) {
	exitF()
}
