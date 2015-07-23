package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func alias() {
	aliases, err := ioutil.ReadDir(os.Getenv("HOME") + "/.drush")
	for _, fileInfo := range aliases {
		if strings.Contains(fileInfo.Name(), ".drushrc.php") {
			aliasFile, err := ioutil.ReadFile(os.Getenv("HOME") + "/.drush/" + fileInfo.Name())
			check(err)
			fmt.Printf("Alias File: %s \n", aliasFile)
		}
	}
	check(err)
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
	check(err)
}
