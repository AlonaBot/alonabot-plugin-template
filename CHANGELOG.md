# Changelog

本文件的格式遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/)，版本号遵循[语义化版本](https://semver.org/lang/zh-CN/)。发布时 `scripts/release.sh` 取出对应版本的一节作为 Release 说明，安装后 WebUI 也会显示它。

## [Unreleased]

## [0.1.0] - 2026-11-17

### Added

- `hello` 命令：用配置的问候语问候，省略名字时问候所有人。
- `roll` 命令（别名 `dice`）：掷骰子，面数缺省取配置的 `dice.default_sides`。
