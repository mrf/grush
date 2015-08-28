package includes

import (
	"os"
	"path"
	"path/filepath"
)

// LocateDrupalRoot performs an exhaustive depth-first search to try and locate the
// Drupal root directory. This makes it possible to run Grush from a
// subdirectory of the drupal root.
// Reference implementation: drush_locate_root()
func LocateDrupalRoot(startPath string) string {
	// Empty string assumes current working directory
	if startPath == "" || startPath == "." {
		// Convert to abspath
		resultPath, err := os.Getwd()
		if err != nil {
			return ""
		}
		startPath = resultPath
	}

	// Clean up the path
	startPath = path.Clean(startPath)
	resultPath, err := filepath.EvalSymlinks(startPath)
	if err != nil {
		return ""
	}
	startPath = resultPath

	drupalRoot := ""
	if validDrupalRoot(startPath) {
		drupalRoot = startPath
	}
	// @todo: This only checks the current directory right now.

	return drupalRoot
}

// validDrupalRoot checks whether given path qualifies as a Drupal root.
func validDrupalRoot(path string) bool {
	return FindBootstrapForRoot(path) != nil
}
