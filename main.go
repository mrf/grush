package main

import (
	"fmt"
	"io/ioutil"
	"os"
) func main() {
	cmd := os.Args[0]
	fmt.Printf("Program name: %s\n", cmd)
	for i, a := range os.Args[2:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}

	command := os.Args[1]

	fmt.Printf("Main command %s", command)
	if command == "sa" {
		fmt.Printf("SITEN ALIASE")
		aliases, err := ioutil.ReadDir("/home/mark/.drush")
		//fmt.Printf("alieases: %s", aliases)
		for i, alias := range aliases {
			if alias.IsDir {
				fmt.Printf("Alias: %s", alias.Name, i+1)
			}
		}
		if err != nil {
			fmt.Printf("ReadDir %s: %v", aliases, err)
		}
	}
}
