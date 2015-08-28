package includes

// Reference implementation: drush_get_bootstrap_candidates()
func getBootstrapCandidates() (candidates map[string]Boot) {
	// @todo: Make static
	candidates = make(map[string]Boot)
	candidates["drupal7"] = DrupalBoot7{}
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
