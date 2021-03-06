package includes

import "os"

// Boot interface describes the functions available for booting a CMS.
type Boot interface {
	ValidRoot(path string) bool
	//ContribModulesPaths()
	//ContribThemesPaths()
}

// DrupalBoot stores functions Drupal specific boot functions.
type DrupalBoot struct {
}

// DrupalBoot7 is for Drupal 7 specific functions.
type DrupalBoot7 struct {
	DrupalBoot
}

// DrupalBoot8 is for Drupal8 specific functions.
type DrupalBoot8 struct {
	DrupalBoot
}

// ValidRoot returns true when a Drupal7 root is located.
// Reference implementation: drush/lib/Drush/Boot/DrupalBoot7.php
func (d DrupalBoot7) ValidRoot(path string) bool {
	if !fileExists(path + "/index.php") {
		return false
	}

	const Candidate = "includes/common.inc"
	// Uses the same file_exist tests used in Drush
	if fileExists(path+"/"+Candidate) &&
		fileExists(path+"/misc/drupal.js") &&
		fileExists(path+"/modules/field/field.module") {
		return true
	}

	return false
}

// ValidRoot returns true when a Drupal8 root is located.
// Reference implementation: drush/lib/Drush/Boot/DrupalBoot8.php
func (d DrupalBoot8) ValidRoot(path string) bool {
	if !fileExists(path + "/index.php") {
		return false
	}

	const Candidate = "core/includes/common.inc"
	// Uses the same file_exist tests used in Drush
	if fileExists(path+"/"+Candidate) &&
		fileExists(path+"/core/core.services.yml") {
		if fileExists(path+"/core/misc/drupal.js") ||
			fileExists(path+"/core/assets/js/drupal.js") {
			return true
		}
	}
	return false
}

func fileExists(path string) bool {
	finfo, err := os.Stat(path)
	if err != nil {
		// Doesn't exist.
		return false
	}
	if finfo.IsDir() {
		// Not a file.
		return false
	}
	return true
}
