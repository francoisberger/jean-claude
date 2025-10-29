package main

import (
    "fmt"
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

}
