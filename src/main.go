package main

import (
	"jean-claude/commandline"
	"bufio"
	"fmt"
	"os"
)

// Will instantiate JeanClaude and read user input
func main() {
	jc := JeanClaude {}
	jc.Start()
	
	exit := false
    for !exit {
        // TODO: smarter prompt read. Maybe introduce a prompt object build by readLine or analyzeCommand
        command := commandline.NewCommandLine(readLine())
        switch command.Code() {
        case commandline.Exit: 
            exit = true
        case commandline.LoadWordList:
        case commandline.LoadText:
        case commandline.LoadDictionary:
        case commandline.Stats:
        default:
            fmt.Println("Could not understand. Please clarify.")
        }
    }

	jc.Stop()
}

func readLine() string {
    fmt.Print("JC> ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() 
    if scanner.Err() != nil {
        // TODO: raise exception
        return "IO Error"
    } else {
        return scanner.Text()
    }
}
