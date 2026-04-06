package config

// Config represents the top-level configuration file.
type Config struct {
	// Repositories defines named git sources that entries can reference.
	Repositories map[string]Repository `toml:"repositories"`
	// Entries are the managed dotfiles/directories, keyed by name.
	Entries map[string]Entry `toml:"entries"`
}

// Repository defines a named git repository source.
type Repository struct {
	// URL is the git clone URL (SSH or HTTPS).
	URL RepoURL `toml:"url"`
}

type RepoURL string

// Entry represents a single managed dotfile or directory.
type Entry struct {
	// Path is the local destination path (supports ~ expansion).
	Path string `toml:"path"`
	// Source defines where to fetch this entry from.
	Source Source `toml:"source"`
	// Symlink controls the placement strategy. Defaults to true (symlink).
	// Set to false for direct copy.
	DisableSymlink bool `toml:"disableSymlink"`
}

// Source references a named repository and a path within it.
type Source struct {
	// Repository is the name of a repository defined in [repositories].
	Repository string `toml:"repository"`
	// Path is the path within the repository to the file or directory.
	Path string `toml:"path"`
	// Ref is an optional git ref (branch, tag, commit). Defaults to HEAD.
	Ref string `toml:"ref,omitempty"`
}
