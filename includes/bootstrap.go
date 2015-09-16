package includes

import (
	"fmt"
	"os"
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
	// todo: check if php exists and error out before execution see: exec.LookPath

	cmd := exec.Command("php", "-r print 'we have php';")

	output, err := cmd.CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
	fmt.Printf("==> Output: %s\n", string(output))
	return nil
}
