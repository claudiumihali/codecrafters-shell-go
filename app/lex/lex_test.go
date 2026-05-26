package lex

import "testing"

func TestLex(t *testing.T) {
	tests := map[string]struct {
		input string
		args  []string
	}{
		"empty": {
			input: "",
			args:  []string{},
		},
		"space": {
			input: " ",
			args:  []string{},
		},
		"single_arg": {
			input: "grape",
			args:  []string{"grape"},
		},
		"multiple_args": {
			input: "grape orange raspberry",
			args:  []string{"grape", "orange", "raspberry"},
		},
		"args_with_spaces": {
			input: "grape   orange    raspberry",
			args:  []string{"grape", "orange", "raspberry"},
		},
		"args_with_spaces_in_the_beginning_and_end": {
			input: "  grape orange raspberry  ",
			args:  []string{"grape", "orange", "raspberry"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			args := Parse(test.input)

			if len(args) != len(test.args) {
				t.Fatalf("expected %d args, actual %d", len(test.args),
					len(args))
			}

			for i := range args {
				if args[i] != test.args[i] {
					t.Errorf("expected arg %d to be %q, actual %q", i,
						test.args[i], args[i])
				}
			}
		})
	}
}
