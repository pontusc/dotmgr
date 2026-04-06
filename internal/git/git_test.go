package git_test

import (
	"dotmgr/internal/git"
	"testing"
)

const repoURL = "https://github.com/PontusC/.dotfiles"
const testDir = "/tmp/dotmgr-testing"

func TestRepoClone(t *testing.T) {
	err := git.Clone(repoURL, testDir)
	if err != nil {
		t.Error("testing repo clone got error: %w", err)
	}
}
