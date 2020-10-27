package main

import "fmt"

func main() {
	var w Writer = ConsolePrettyWriter{}
	var cpw ConsolePrettyWriter = w.(ConsolePrettyWriter)
	cpw.setPrettyCharacter("***")
	cpw.Write([]byte("Hello!"))
}

// Write interface
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter struct
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// ConsolePrettyWriter struct
type ConsolePrettyWriter struct {
	prettyCharacter string
}

func (cpw ConsolePrettyWriter) Write(data []byte) (int, error) {
	n, err := fmt.Printf("%v %v %v", cpw.prettyCharacter, string(data), cpw.prettyCharacter)
	return n, err
}

func (cpw *ConsolePrettyWriter) setPrettyCharacter(pretty string) {
	cpw.prettyCharacter = pretty
}
