# dotmgr

Stateless dotfiles manager. Retrieves files from git repositories and places them at configured local paths.

## Learning Project

This project exists for the user to learn Go. **Do not write implementation code.** Instead:

- **Explain** concepts, patterns, and Go idioms when asked
- **Review** code the user has written — point out issues, suggest improvements
- **Advise** on design decisions, architecture, and idiomatic Go approaches
- **Guide** debugging by asking questions and narrowing down the problem, not by writing the fix

Chores and cleanup tasks (formatting, linting, CI config, Makefile updates, CLAUDE.md updates) are fine to implement directly.

Agents (expert, scout, infra) should follow the same rule: provide explanations and guidance, not implementations.

## Architecture

- **CLI**: Cobra subcommands (`sync`, `diff`, `status`)
- **Config**: TOML at `$XDG_CONFIG_HOME/dotmgr/config.toml` (defaults to `~/.config/dotmgr/config.toml`)
- **Git**: go-git (pure Go, no external git binary required). SSH URLs (`git@host:path.git`) for private repos, HTTPS for public repos.
- **Placement**: Two strategies controlled by `symlink` bool (default `true`):
  - **Symlink** (default): clone repo to `$XDG_DATA_HOME/dotmgr/`, symlink targets to cache paths. Enables editing and pushing back to source.
  - **Direct copy** (`symlink = false`): fetch file content from remote, write to local path. Stateless.

## Project Layout

```
cmd/dotmgr/        # entrypoint
cmd/sync/           # sync subcommand
cmd/diff/           # diff subcommand
cmd/status/         # status subcommand
internal/config/    # TOML config parsing
internal/git/       # go-git repository operations
internal/placement/ # file placement (copy / symlink)
internal/cache/     # local repo cache for clone strategy
```

## Build

All builds and tests run in Docker — no local Go toolchain required.

```sh
make build   # docker buildx bake → tests then compiles, outputs bin/dotmgr
make clean   # remove bin/
```

## Config Format

Repositories are defined once and referenced by name from entries. See `internal/config/testdata/config.toml` for the canonical test fixture.

```toml
# Named repository sources
[repositories.dotfiles]
url = "git@github.com:pontusc/.dotfiles.git"

[repositories.private-infra]
url = "git@github.com:pontusc/infra-configs.git"

# Entries are top-level keys. Symlink is default (true).
[yamlfmt]
path = "~/.config/yamlfmt/config"
source = { repository = "dotfiles", path = "yamlfmt/config" }

# With explicit git ref
[starship]
path = "~/.config/starship.toml"
source = { repository = "dotfiles", path = "starship/starship.toml", ref = "main" }

# Opt out of symlink with symlink = false
[ssh-allowed-signers]
path = "~/.ssh/allowed_signers"
symlink = false
source = { repository = "private-infra", path = "ssh/allowed_signers" }
```

## Conventions

- Go 1.26, module path `dotmgr`
- All packages under `internal/` — no exported API
- Tests next to source (`_test.go` in same package or `_test` package)
- Test fixtures in `testdata/` directories
- Builds and tests run in Docker via `docker-bake.hcl` — no local toolchain deps
- No CGO — pure Go for easy cross-compilation
- **Keep this file in sync with implementation** — update config format, commands, and build instructions here when they change
