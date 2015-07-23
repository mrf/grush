package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {

		currentcommand := os.Args[1]
		fmt.Printf("Main command %s \n", currentcommand)
		if currentcommand == "alias" {
			alias()
		} else if currentcommand == "list" {
			list()
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
	readSettings()
}

func readSettings() {
	// see if we have alias and read from alias file
	// current directory and sites/default/files deep
	currentdir, err := ioutil.ReadDir(".")
	for _, fileInfo := range currentdir {
		if fileInfo.IsDir() {
			fmt.Printf("File %s", fileInfo)
			ioutil.ReadFile("index.php")
		}
	}
	if err != nil {
		fmt.Printf("ReadDir %s: %v", currentdir, err)
	}
}
