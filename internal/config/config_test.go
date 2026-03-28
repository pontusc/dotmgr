package config_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"dotmgr/internal/config"
)

func testdataPath(name string) string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "testdata", name)
}

func TestLoadFrom(t *testing.T) {
	cfg, err := config.LoadFrom(testdataPath("config.toml"))
	if err != nil {
		t.Fatalf("LoadFrom: %v", err)
	}

	// Repositories
	if len(cfg.Repositories) != 2 {
		t.Fatalf("expected 2 repositories, got %d", len(cfg.Repositories))
	}
	if cfg.Repositories["dotfiles"].URL != "git@github.com:pontusc/.dotfiles.git" {
		t.Errorf("dotfiles url = %q", cfg.Repositories["dotfiles"].URL)
	}
	if cfg.Repositories["private-infra"].URL != "git@github.com:pontusc/infra-configs.git" {
		t.Errorf("private-infra url = %q", cfg.Repositories["private-infra"].URL)
	}

	// Entry count
	if len(cfg.Entries) != 7 {
		t.Fatalf("expected 7 entries, got %d", len(cfg.Entries))
	}

	// Symlink default (true)
	yf := cfg.Entries["yamlfmt"]
	if yf.Path != "~/.config/yamlfmt/config" {
		t.Errorf("yamlfmt path = %q", yf.Path)
	}
	if yf.Source.Repository != "dotfiles" {
		t.Errorf("yamlfmt source.repository = %q", yf.Source.Repository)
	}
	if yf.Source.Path != "yamlfmt/config" {
		t.Errorf("yamlfmt source.path = %q", yf.Source.Path)
	}
	if !yf.IsSymlink() {
		t.Error("yamlfmt should default to symlink=true")
	}

	// Explicit ref
	star := cfg.Entries["starship"]
	if star.Source.Ref != "main" {
		t.Errorf("starship source.ref = %q, want %q", star.Source.Ref, "main")
	}

	// Symlink opt-out
	ssh := cfg.Entries["ssh-allowed-signers"]
	if ssh.IsSymlink() {
		t.Error("ssh-allowed-signers should have symlink=false")
	}
	if ssh.Source.Repository != "private-infra" {
		t.Errorf("ssh-allowed-signers source.repository = %q", ssh.Source.Repository)
	}
}