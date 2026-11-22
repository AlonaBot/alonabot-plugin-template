# Changelog

本文件的格式遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/)，版本号遵循[语义化版本](https://semver.org/lang/zh-CN/)。发布时 `scripts/release.sh` 取出对应版本的一节作为 Release 说明，安装后 WebUI 也会显示它。

## [Unreleased]

### Changed

- SDK 与 CLI 固定使用正式版本 `v0.2.0`，工作流直接获取公开 SDK。
- `hello` 命令通过 `.Menu()` 加入自动菜单，开发示例同步菜单声明方式。
- 源码安装 CLI 时沿用组件的 Go 工具链，使声明扫描器与组件导出数据匹配。

## [0.1.0] - 2026-11-17

### Added

- `hello` 命令：用配置的问候语问候，省略名字时问候所有人。
