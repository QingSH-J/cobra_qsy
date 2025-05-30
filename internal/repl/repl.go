package repl

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

// startREPL starts a Read-Eval-Print Loop
func Start(rootCmd *cobra.Command) {
	fmt.Println("Starting REPL, type 'exit' to quit.")

	//config readline
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "qsy> ",
		HistoryFile:     os.TempDir() + "/repl_history.txt",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		fmt.Println("Error initializing shell:", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println("Error reading line:", err)
			break
		}
		line = strings.TrimSpace(line)
		if line == "exit" {
			fmt.Println("Exiting REPL.")
			break
		}

		if line == "" {
			continue
		}

		//execute the command
		executeCommand(rootCmd, line)
	}
}

// executeCommand
func executeCommand(rootCmd *cobra.Command, line string) {
	args := parseCommand(line)

	if len(args) == 0 {
		fmt.Println("No command entered.")
		return
	}

	cmd := *rootCmd
	cmd.SetArgs(args)

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error executing command '%s': %v\n", line, err)
	} else {
		fmt.Printf("Executed command: %s\n", line)
	}
}

// parseCommand
func parseCommand(line string) []string {
	var args []string
	var currentArg strings.Builder
	inQuotes := false

	for _, char := range line {
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ' ':
			if inQuotes {
				currentArg.WriteRune(char)
			} else if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
		default:
			currentArg.WriteRune(char)
		}
	}
	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}
	return args
}
