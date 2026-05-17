```bash
[tester::#IZ3] Running tests for Stage #IZ3 (Implement echo)
[tester::#IZ3] Running ./your_program.sh
[your-program] $ echo mango grape
[your-program] mango grape
[tester::#IZ3] ✓ Received expected response
[your-program] $ echo grape orange raspberry
[your-program] grape orange raspberry
[tester::#IZ3] ✓ Received expected response
[your-program] $ echo banana mango apple
[your-program] banana mango apple
[tester::#IZ3] ✓ Received expected response
[your-program] $ 
[tester::#IZ3] Test passed.

[tester::#PN5] Running tests for Stage #PN5 (Implement exit)
[tester::#PN5] Running ./your_program.sh
[your-program] $ invalid_grape_command
[your-program] invalid_grape_command: command not found
[tester::#PN5] ✓ Received command not found message
[your-program] $ exit
[tester::#PN5] ✓ Program exited successfully
[tester::#PN5] ✓ No output after exit command
[tester::#PN5] Test passed.

[tester::#FF0] Running tests for Stage #FF0 (Implement a REPL)
[tester::#FF0] Running ./your_program.sh
[your-program] $ invalid_command_1
[your-program] invalid_command_1: command not found
[tester::#FF0] ✓ Received command not found message
[your-program] $ invalid_command_2
[your-program] invalid_command_2: command not found
[tester::#FF0] ✓ Received command not found message
[your-program] $ invalid_command_3
[your-program] invalid_command_3: command not found
[tester::#FF0] ✓ Received command not found message
[your-program] $ invalid_command_4
[your-program] invalid_command_4: command not found
[tester::#FF0] ✓ Received command not found message
[your-program] $ invalid_command_5
[your-program] invalid_command_5: command not found
[tester::#FF0] ✓ Received command not found message
[your-program] $ 
[tester::#FF0] Test passed.

[tester::#CZ2] Running tests for Stage #CZ2 (Handle invalid commands)
[tester::#CZ2] Running ./your_program.sh
[your-program] $ invalid_raspberry_command
[your-program] invalid_raspberry_command: command not found
[tester::#CZ2] ✓ Received command not found message
[tester::#CZ2] Test passed.

[tester::#OO8] Running tests for Stage #OO8 (Print a prompt)
[tester::#OO8] Running ./your_program.sh
[your-program] $ 
[tester::#OO8] ✓ Received prompt
[tester::#OO8] Test passed.
```
