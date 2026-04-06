// Package git provides repository operations using go-git.
package git

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func SetupRepoCache(url string, repoPath string) error {
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		// Case: destination does not exist, clone repo
		_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
			URL: url,
		})
		if err != nil {
			return fmt.Errorf("attempt to clone %v: %w", url, err)
		}
	} else {
		// Case: destination exists, load
		_, err := git.PlainOpen(repoPath)
		if err != nil {
			return fmt.Errorf("attempt to load %v, %w:", repoPath, err)
		}
	}
	return nil
}
