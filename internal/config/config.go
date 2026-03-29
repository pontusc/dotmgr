package config

import (
	"bytes"
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

// LoadFrom reads and parses a config file from the given path.
func LoadFrom(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	// Decode into a raw map to separate repositories from entries.
	var raw map[string]any
	if _, err := toml.Decode(string(data), &raw); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	cfg := &Config{
		Repositories: make(map[string]Repository),
		Entries:      make(map[string]Entry),
	}

	// Extract repositories via re-encode/decode of just that section.
	if repos, ok := raw["repositories"]; ok {
		var buf bytes.Buffer
		if err := toml.NewEncoder(&buf).Encode(map[string]any{"repositories": repos}); err != nil {
			return nil, fmt.Errorf("re-encoding repositories: %w", err)
		}
		var wrapper struct {
			Repositories map[string]Repository `toml:"repositories"`
		}
		if _, err := toml.Decode(buf.String(), &wrapper); err != nil {
			return nil, fmt.Errorf("parsing repositories: %w", err)
		}
		cfg.Repositories = wrapper.Repositories
	}

	// Everything else is an entry — re-encode each and decode as Entry.
	for key, val := range raw {
		if key == "repositories" {
			continue
		}
		var buf bytes.Buffer
		if err := toml.NewEncoder(&buf).Encode(val); err != nil {
			return nil, fmt.Errorf("re-encoding entry %q: %w", key, err)
		}
		var entry Entry
		if _, err := toml.Decode(buf.String(), &entry); err != nil {
			return nil, fmt.Errorf("parsing entry %q: %w", key, err)
		}
		cfg.Entries[key] = entry
	}

	return cfg, nil
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

// IsSymlink returns whether this entry uses symlink placement (default true).
func (e Entry) IsSymlink() bool {
	if e.Symlink == nil {
		return true
	}
	return *e.Symlink
}
