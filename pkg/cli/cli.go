package main

import (
	"fmt"
)

type ApiDef struct {
}

func main() {
	fmt.Println("Missaku CLI - easily manage API tokens and settings")
	var userInput string
	stat := 0
	for stat == 0 {
		fmt.Print("$ ")
		fmt.Scanln(&userInput)
		switch userInput {
		case "api":
			fmt.Println("Enter API key: ")
		case "version":
			fmt.Println("Missaku V0.1")

		}
		//realize normal STD::in, out, err buffer

		// when api typed, parsing should be done
		// it will read text data. Example: api apikey.txt, api -t token.txt
		//$ api token.txt
		//$ nf - new file
		//
	}
}
