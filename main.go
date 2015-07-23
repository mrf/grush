package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {

		commands := map[string]string{
			"sa": "alias",
			"ls": "list",
		}
		for _, command := range commands {
			currentcommand := os.Args[1]
			fmt.Printf("Main command %s \n", currentcommand)
			if currentcommand == "site-alias" {
				alias()
			} else if command == "list" {
				fmt.Print("Command Not Found")
			}
		}
	} else {
		fmt.Printf("Grush Version 0.0.0")
	}
}

func alias() {
	aliases, err := ioutil.ReadDir(os.Getenv("HOME") + "/.drush")
	for _, fileInfo := range aliases {
		fmt.Printf("\n File Name : %s \n", fileInfo.Name())
	}
	if err != nil {
		fmt.Printf("ReadDir %s: %v", aliases, err)
	}
}

func list() {
	fmt.Printf("Command not implemented")
}

func readSettings() {
	// see if we have alias and read from alias file
	// current directory and sites/default/files deep
	currentdir, err := ioutil.ReadDir(".")
	for _, fileInfo := range currentdir {
		fmt.Printf("File %s", fileInfo.Name)
		ioutil.ReadFile("index.php")
	}
	if err != nil {
		fmt.Printf("ReadDir %s: %v", currentdir, err)
	}
}
