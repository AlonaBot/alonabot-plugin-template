# Packaging and releases

[简体中文](releasing.zh-CN.md)

Prepare the locked SDK CLI following [development](development.md). The five supported targets are linux/amd64, linux/arm64, windows/amd64, darwin/amd64, and darwin/arm64.

```fish
alonabot check .
alonabot pack . --target linux/amd64 --out bin
```

Maintainers prepare `plugin.yaml` and `CHANGELOG.md`, moving `[Unreleased]` into the matching version section. After review, publish the matching `v<version>` tag. `scripts/release.sh` checks version and source against the repository, builds five archives, and extracts that version section without its heading. The release workflow publishes the archives and notes. Keep the `owner/alonabot-plugin-template` placeholders until template setup.

Release assets contain platform `.tar.gz` archives; GitHub provides source downloads. Core reads component identity and version from the packaged `plugin.yaml` and binds the repository source. Upgrade Core to v0.2.0 before installing or updating formal components; stop the component and back up its data before replacement, retaining its ID, source, and data directory.

[Core release maintenance](https://github.com/AlonaBot/AlonaBot/blob/main/docs/releasing.md)
