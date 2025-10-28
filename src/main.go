package main

import (
	"jeanclaude/jeanclaude"
)

func main() {
	jc := jeanclaude.JeanClaude {}
	jc.Start()
	jc.Run()
	jc.Stop()
}