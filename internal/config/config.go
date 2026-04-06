package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Load reads and parses the config file from the default XDG location.
func Load() (*Config, error) {
	path, err := DefaultPath()
	if err != nil {
		return nil, err
	}
	return LoadFrom(path)
}

// DefaultPath returns the default config file path at $XDG_CONFIG_HOME/dotmgr/config.toml.
func DefaultPath() (string, error) {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolving home directory: %w", err)
		}
		dir = filepath.Join(home, ".config")
	}
	return filepath.Join(dir, "dotmgr", "config.toml"), nil
}

// LoadFrom reads and parses a config file from the given path.
func LoadFrom(path string) (*Config, error) {
	cfg := &Config{
		Repositories: make(map[string]Repository),
		Entries:      make(map[string]Entry),
	}

	// metadata (unused keys, misspelled keys) are discarded for now.
	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return cfg, nil
}
