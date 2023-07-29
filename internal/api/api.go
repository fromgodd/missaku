package api

import (
	"fmt"
	"net/http"
	"os"
)

func init() {

	client := &http.Client{}

}

func checkFile() {
	if _, err := os.Stat("token.txt"); err == nil {
		token, err := os.ReadFile("token.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func fetchToken() {
	// DL13GP8EH5CewI9XbDDyPOle3YzbYH1B
	//Initial connection

}
