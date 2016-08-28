package main

import "fmt"
import "os"
import "strings"
import "github.com/mattmoore/novo/console"

func main() {
	fmt.Println(console.Parse(strings.Join(os.Args[1:], " ")))
}
