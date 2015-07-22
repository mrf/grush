package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	program := os.Args[0]
	fmt.Printf("Program name: %s\n", program)

	command := os.Args[1]

	fmt.Printf("Main command %s", command)
	if command == "sa" {
		aliases, err := ioutil.ReadDir("/home/mark/.drush")
		for _, fileInfo := range aliases {
			fmt.Printf("\n File Name : %s \n", fileInfo.Name())
		}
		if err != nil {
			fmt.Printf("ReadDir %s: %v", aliases, err)
		}
	} else {
		fmt.Print("Command Not Found")
	}
}
