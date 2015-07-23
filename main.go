package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {

		commands := map[string]string{
			"sa": "site-alias",
			"ls": "pm-list",
		}
		for _, command := range commands {
			currentcommand := os.Args[1]
			fmt.Printf("Main command %s \n", currentcommand)
			if currentcommand == "site-alias" {
				sitealias()
			} else if command == "" {
				fmt.Print("Command Not Found")
			}
		}
	} else {
		fmt.Printf("Grush Version 0.0.0")
	}
}

func sitealias() {
	aliases, err := ioutil.ReadDir("/home/mark/.drush")
	for _, fileInfo := range aliases {
		fmt.Printf("\n File Name : %s \n", fileInfo.Name())
	}
	if err != nil {
		fmt.Printf("ReadDir %s: %v", aliases, err)
	}
}
