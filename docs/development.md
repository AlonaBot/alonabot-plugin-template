# Development

[简体中文](development.zh-CN.md)

Run these commands from this repository root. Use the Go version declared in `go.mod` (managed with asdf locally); the SDK stays pinned there. Build the CLI with the same component toolchain:

```fish
go mod download github.com/AlonaBot/AlonaBot
env GOTOOLCHAIN=(go env GOVERSION) go -C (go list -m -f '{{.Dir}}' github.com/AlonaBot/AlonaBot) install ./cmd/alonabot
```

Keep the Go bin directory on PATH. For cross-repository development, a local `go.work` may select the corresponding Core checkout; verify releases using the locked SDK.

## Verify a change

Select checks by the affected behavior. CI runs the following Go checks for Go, dependency, and CI gate changes; component resources and checking workflows also trigger `alonabot check .`.

```fish
gofmt -l (git ls-files '*.go')
go vet ./...
go test -race ./...
for target in linux/amd64 linux/arm64 windows/amd64 darwin/amd64 darwin/arm64
    set -l parts (string split / $target)
    env GOOS=$parts[1] GOARCH=$parts[2] CGO_ENABLED=0 go build ./...
end
alonabot check .
alonabot test .
```

This repository has no `go:generate` entry; edit the owned resources directly.

For documentation and GitHub metadata, check the changed links, YAML and scripts, plus `git diff --check`. In CI, Repository compares the event base SHA with HEAD; docs-only changes skip Go checks. Record actual commands, platform, failures, and unrun checks in the PR.

## Source map

- `plugin.go`: public construction entry and example handlers
- `plugin.yaml`: component identity and packaging

[Core development](https://github.com/AlonaBot/AlonaBot/blob/main/docs/development.md) · [Core plugin guide](https://github.com/AlonaBot/AlonaBot/blob/main/docs/plugin-guide.md)

## Register commands and read configuration

In `plugin.go`, call `p.Command("name").Description(...).Menu().Handle(handler)` from `New`. `.Menu()` adds the command to automatic menus. A handler has signature `func(c *sdk.Context, args Args) (message.Content, error)`; Core sends returned content as a reply. Fields map to command arguments in order, and pointer fields are optional.

Read configuration with `c.Config()`, for example `c.Config().Object("greet").String("greeting")`. Update configuration declarations and `plugin_test.go` with new commands.

```fish
alonabot run --reload .
```

This starts Core with the plugin and rebuilds it after source changes. Runtime data goes to ignored `data/`. `alonabot test .` loads, registers, and stops the plugin in a temporary offline session. `plugin_test.go` contains a plugintest integration test. `scripts/release.sh` handles multi-platform packaging; `.github/workflows/` runs checks and releases. Enable GitHub Actions under **Settings → Actions → General** in the generated repository.
