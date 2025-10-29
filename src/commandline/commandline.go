package commandline

import (
	"slices"
	"strings"
)

// The CommandLine struct will contain the command code as well as any parameters 
// relevant for the considered command.
type CommandLine struct {
	code Code
	parameters []string
}

// Returns command's code i.e. a predefined action to be performed based on user input.
func (c CommandLine) Code() Code {
	return c.code
}

// Returns command's parameters i.e. parameters required to perform the required action.
func (c CommandLine) Parameters() string {
	return strings.Join(c.parameters, ", ")
}

// Returns a string representation of the command.
func (c CommandLine) String() string {
	return c.Code().String() + "-" + c.Parameters()
}

// Builds a new command based on user input
func NewCommandLine(input string) CommandLine {
	// Perform basic checks and tokenize input string
	if len(input) == 0 {
    	return CommandLine { code: Unknown }
	}

	// Analyze input string to return proper CommandLine
	arguments := strings.Split(input, " ")
	mainCommand := strings.ToLower(arguments[0])

	// exit command takes no argument
	if mainCommand == "exit" {
		return CommandLine { code: Exit }
	}
	// load command can be:
	// - load wordlist <path>
	// - load text <path>
	// - load dictionary <path>
	if mainCommand == "load" {
		// Ensure we have 3 arguments with subcommand in proper available commands
		if len(arguments) < 3 || !slices.Contains([]string{"wordlist", "text", "dictionary"}, strings.ToLower(arguments[1])) {
			return CommandLine { code: Unknown }
		}
		subCommand := strings.ToLower(arguments[1])
		switch subCommand {
		case "wordlist": 
			return CommandLine { code: LoadWordList, parameters: []string{arguments[2]} }
		case "text": 
			return CommandLine { code: LoadText, parameters: []string{arguments[2]} }
		case "dictionary": 
			return CommandLine { code: LoadDictionary, parameters: []string{arguments[2]} }
		}
		return CommandLine { code: Unknown }
	}
	// Return unknown as a default
	return CommandLine { code: Unknown }
}