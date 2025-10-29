package main

import (
    "bufio"
    "fmt"
    "os"
)

type JeanClaude struct {	
}

func (jc JeanClaude) Start() {
    fmt.Println("Hello, World! My name is Jean-Claude.")
    fmt.Println("Use prompt to interact.")
}

func (jc JeanClaude) Stop() {
    fmt.Println("Goodbye. See you soon.")
}

func (jc JeanClaude) Run() {
    var exit bool = false
    for !exit {
        // TODO: smarter prompt read. Maybe introduce a prompt object build by readPrompt or analyzePrompt
        prompt := analyzePrompt(readPrompt())
        switch prompt.Code() {
        case Exit: 
            exit = true
        case LoadWordList:
        case LoadText:
        case LoadDictionary:
        case Stats:
        default:
            fmt.Println("Could not understand. Please clarify.")
        }
    }
}

func readPrompt() string {
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

func analyzePrompt(prompt string) Prompt {
    var p Prompt = Prompt {}
    return p
}