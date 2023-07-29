package main

import (
	"fmt"
	"os"
)

// Main program with entry point
// TODO Sort Structure
// Implement proper arguments
type Human struct {
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	cliArgument := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(cliArgument)

}
