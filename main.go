package main

import (
  "os"
	"fmt"

	"bufio"
	"io/ioutil"
	"strings"

	"github.com/mrf/grush/commands"
)

func main() {
  commands.Execute()
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
	checkDrupal()
	// current directory and sites/default/files deep
}

/**
 * Check to see if we are in what looks like a Drupal install.
 */
func checkDrupal() {
	currentdir, err := ioutil.ReadDir(".")
	for _, fileInfo := range currentdir {
		if !fileInfo.IsDir() {
			if fileInfo.Name() == "index.php" {
				indexFile, err := os.Open("./index.php")
				check(err)
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
	check(err)
}
