package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * Check to see if we are in what looks like a Drupal install.
 */
func CheckDrupal() {
	currentdir, err := ioutil.ReadDir(".")
	for _, fileInfo := range currentdir {
		if !fileInfo.IsDir() {
			if fileInfo.Name() == "index.php" {
				indexFile, err := os.Open("./index.php")
				Check(err)
				fileScanner := bufio.NewScanner(indexFile)
				for fileScanner.Scan() {
					if strings.Contains(fileScanner.Text(), "Drupal") {
						fmt.Printf("Looks like a Drupal Site")
						return
					}
				}
				fmt.Printf("Not a Drupal Site Try Somewhere Else")
				defer indexFile.Close()
				return
			}
		}
	}
	Check(err)
}
