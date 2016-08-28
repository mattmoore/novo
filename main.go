package main

import "fmt"
import "os"
import "strings"
import "github.com/mattmoore/novo/interpreter"

func main() {
	fmt.Println(interpreter.Parse(strings.Join(os.Args[1:], " ")))
}
