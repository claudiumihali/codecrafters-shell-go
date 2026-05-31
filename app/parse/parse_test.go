package parse

import (
	"errors"
	"testing"
)

type testCase struct {
	input string
	args  []string
	err   error
}

func runTests(t *testing.T, tests map[string]testCase) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			args, err := Split(test.input)
			if !errors.Is(err, test.err) {
				t.Fatalf("expected error: %v, actual: %v", test.err, err)
			}

			if len(args) != len(test.args) {
				t.Fatalf("expected args no: %d, actual: %d", len(test.args),
					len(args))
			}

			for i := range args {
				if args[i] != test.args[i] {
					t.Errorf("expected arg %d to be: %q, actual: %q", i,
						test.args[i], args[i])
				}
			}
		})
	}
}

func TestParseArg(t *testing.T) {
	tests := map[string]testCase{
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
	runTests(t, tests)
}

func TestParseSingleQuotes(t *testing.T) {
	tests := map[string]testCase{
		"spaces_within_single_quotes": {
			input: "echo 'hello    world'",
			args:  []string{"echo", "hello    world"},
		},
		"empty_single_quotes": {
			input: "echo hello '' world",
			args:  []string{"echo", "hello", "", "world"},
		},
		"single_quotes_in_arg": {
			input: "echo he'llo'wo'rld'",
			args:  []string{"echo", "helloworld"},
		},
		"adjacent_single_quotes_in_arg": {
			input: "echo hello''world",
			args:  []string{"echo", "helloworld"},
		},
		"adjacent_single_quotes_in_single_quoted_arg": {
			input: "echo 'he  llo''world'",
			args:  []string{"echo", "he  lloworld"},
		},
		"single_quotes_not_closed": {
			input: "echo 'hello world",
			err:   unterminatedSingleQuoteErr,
		},
	}
	runTests(t, tests)
}

func TestParseDoubleQuotes(t *testing.T) {
	tests := map[string]testCase{
		"spaces_within_double_quotes": {
			input: `echo "hello    world"`,
			args:  []string{"echo", "hello    world"},
		},
		"empty_double_quotes": {
			input: `echo hello "" world`,
			args:  []string{"echo", "hello", "", "world"},
		},
		"double_quotes_in_arg": {
			input: `echo he"llo"wo"rld"`,
			args:  []string{"echo", "helloworld"},
		},
		"adjacent_double_quotes_in_arg": {
			input: `echo hello""world`,
			args:  []string{"echo", "helloworld"},
		},
		"adjacent_double_quotes_in_double_quoted_arg": {
			input: `echo "he  llo""world"`,
			args:  []string{"echo", "he  lloworld"},
		},
		"double_quotes_not_closed": {
			input: `echo "hello world`,
			err:   unterminatedDoubleQuoteErr,
		},
	}
	runTests(t, tests)
}

func TestParseBackslash(t *testing.T) {
	tests := map[string]testCase{
		"backslash_before_space": {
			input: `echo three\ \ \ spaces and\  one`,
			args:  []string{"echo", "three   spaces", "and ", "one"},
		},
		"backslash_before_backslash": {
			input: `echo hello\\world`,
			args:  []string{"echo", `hello\world`},
		},
		"backslash_before_quote": {
			input: `echo \'hello\'`,
			args:  []string{"echo", "'hello'"},
		},
	}
	runTests(t, tests)
}
