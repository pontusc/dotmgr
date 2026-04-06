// Package cache manages the local repository cache for clone-and-symlink sources.
package cache

import (
	"fmt"
	"os"
	"path/filepath"
)

// Holds the path where are all git repos should be cached
type Cache struct {
	cacheDir string
}

// Standard constructor
func New() (*Cache, error) {
	dir, err := DefaultDir()
	if err != nil {
		return nil, err
	}

	return &Cache{cacheDir: dir}, nil
}

// This constructor is solely used for testing
func NewWithDir(overrideDir string) *Cache {
	return &Cache{cacheDir: overrideDir}
}

// Based on given repository name, returns the full path of where it should be stored
func (c *Cache) RepoPath(name string) string {
	return filepath.Join(c.cacheDir, name)
}

// DefaultDir returns the directory where git repos should be stored
func DefaultDir() (string, error) {
	const cacheDir string = "dotmgr"
	dir := os.Getenv("XDG_DATA_HOME")
	if dir == "" {
		// If XDG_DATA_HOME isnt set, default to $HOME/.local/share/dotmgr
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolving home directory: %w", err)
		}
		dir = filepath.Join(home, ".local", "share")
	}

	return filepath.Join(dir, cacheDir), nil
}
