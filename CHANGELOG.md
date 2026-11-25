# Changelog

本文件的格式遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/)，版本号遵循[语义化版本](https://semver.org/lang/zh-CN/)。发布时 `scripts/release.sh` 取出对应版本的一节作为 Release 说明，安装后 WebUI 也会显示它。条目说明开发者能看到的变化，PR 编号链接用括号放在句末标点前。

## [Unreleased]

### Changed

- `hello` 命令会显示在自动命令菜单中；模板演示了如何用 `.Menu()` 声明菜单项。
- 模板使用公开的 AlonaBot v0.2.0 SDK 和 CLI，开发者可以直接用正式版本构建和打包。

## [0.1.0] - 2026-11-17

### Added

- `hello` 命令：用配置的问候语问候，省略名字时问候所有人。
