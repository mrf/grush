package includes

import (
	"fmt"
	"log"
	"os/exec"
)

// Reference implementation: drush_get_bootstrap_candidates()
func getBootstrapCandidates() (candidates map[string]Boot) {
	// @todo: Make static
	candidates = make(map[string]Boot)
	candidates["drupal8"] = DrupalBoot8{}
	return
}

// FindBootstrapForRoot checks all candidates if ValidRoot returns true.
// Reference implementation: drush_bootstrap_class_for_root()
func FindBootstrapForRoot(path string) Boot {
	candidates := getBootstrapCandidates()
	for _, c := range candidates {
		if c.ValidRoot(path) {
			return c
		}
	}
	return nil
}

func Boot8(level string) Boot {
	// todo: include boot level
	app := "php"

	flags := "-r"
	code := "echo getcwd();"

	cmd := exec.Command(app, flags, code)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PHP Says:", stdout)
	return nil
}
