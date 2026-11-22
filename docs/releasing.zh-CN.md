# 打包与发布

[English](releasing.md)

先按[开发指南](development.zh-CN.md)准备锁定 SDK 的 CLI。支持 linux/amd64、linux/arm64、windows/amd64、darwin/amd64、darwin/arm64 五目标。

```fish
alonabot check .
alonabot pack . --target linux/amd64 --out bin
```

维护者准备 `plugin.yaml` 与 `CHANGELOG.md`，将 `[Unreleased]` 整理到对应版本。审查后发布匹配的 `v<版本>` tag。`scripts/release.sh` 核对版本和仓库来源，构建五平台归档，提取该版本正文并跳过版本标题。发布工作流上传归档与说明。模板替换前保留 `owner/alonabot-plugin-template` 占位。

发行资产提供各平台 `.tar.gz` 归档，源码下载由 GitHub 提供。Core 从包内 `plugin.yaml` 读取组件身份与版本并绑定仓库来源。正式组件安装或更新前先升级 Core v0.2.0；替换前停止组件并备份数据，保留组件 ID、来源和数据目录。

[Core 发行维护](https://github.com/AlonaBot/AlonaBot/blob/main/docs/releasing.md)
