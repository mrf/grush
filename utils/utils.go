package utils

import (
  "os"
  "fmt"
  "io/ioutil"
  "bufio"
  "strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadSettings() {
	// see if we have alias and read from alias file
	CheckDrupal()
	// current directory and sites/default/files deep
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
