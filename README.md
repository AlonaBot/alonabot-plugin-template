<div align="center">

# alonabot-plugin-template

_AlonaBot 插件简单描述_

[![CI](https://github.com/owner/alonabot-plugin-template/actions/workflows/ci.yml/badge.svg)](https://github.com/owner/alonabot-plugin-template/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/owner/alonabot-plugin-template)](https://github.com/owner/alonabot-plugin-template/releases)
[![Go](https://img.shields.io/badge/Go-1.27.3-00ADD8?logo=go)](https://go.dev/)
[![AlonaBot](https://img.shields.io/badge/AlonaBot-plugin-blue)](https://github.com/AlonaBot/AlonaBot)

</div>

这是 [AlonaBot](https://github.com/AlonaBot/AlonaBot) 的插件模板，包含 `hello` 示例命令、配置、集成测试和多平台发布工作流。

<details open>
<summary>模板库使用方法</summary>

1. 点击 [使用此模板](https://github.com/new?template_owner=AlonaBot&template_name=alonabot-plugin-template) 创建仓库。
2. 在新仓库中按下表全局替换占位名称。填写仓库所有者、仓库名与插件 ID，插件 ID 只使用小写字母、数字与连字符。

   | 占位内容 | 替换为 | 示例 |
   | --- | --- | --- |
   | `owner` | 你的 GitHub 用户名或组织名 | `alice` |
   | `alonabot-plugin-template` | 你的仓库名与插件 ID，保留 `alonabot-plugin-` 前缀 | `alonabot-plugin-weather` |

   替换后，`go.mod` 的 module、`plugin.yaml` 的 source、README 的徽章、安装与仓库链接应共同指向你的仓库。保留 SDK 的 `github.com/AlonaBot/AlonaBot` 路径与模板创建链接。
3. 修改 `plugin.yaml` 的 `name`、`description`、`author`，以及本页标题、简介、配置和指令说明；Go 徽章版本与 `go.mod` 保持一致。
4. 在仓库首页 **About → 设置 → Topics** 添加 `alonabot`、`alonabot-plugin` 标签。
5. 处理函数直接写在 `plugin.go`，在 `New` 中注册命令。回复直接构造文本，配置说明写在 `config_schema.json`。
6. 运行 `go mod tidy` 和下方开发命令，更新 `CHANGELOG.md`，然后删去这段模板使用说明。

需要开源许可证时，在新仓库通过 **Add file → Create new file → LICENSE → Choose a license template** 自行选择。

</details>

## 介绍

示例提供 `hello` 命令：读取配置的问候语，带一个可选的名字参数。省略名字时回复 `Hello, everyone!`，发送 `/hello Alona` 时回复 `Hello, Alona!`。

按插件实际功能替换本节介绍。

## 安装

在运行 AlonaBot 的机器上执行：

```sh
alonabot install --source owner/alonabot-plugin-template
```

也可以在 WebUI 的「插件 → 商店」中安装。首次安装会请管理员确认信任这个来源。

## 配置

在 WebUI 的插件详情页修改配置，保存后插件重载生效。

| 配置项 | 类型 | 默认值 | 说明 |
| --- | --- | --- | --- |
| `greet.greeting` | string | `Hello` | 名字前的问候语 |

## 使用

| 命令 | 参数 | 说明 |
| --- | --- | --- |
| `hello` | `[名字]` | 用问候语问候，省略名字时问候所有人 |

命令前缀由 AlonaBot 设置，默认为 `/`，例如 `/hello 阿洛娜`。

## 开发

需要 Go（版本见 `go.mod`）与 AlonaBot CLI。CLI 的版本应与 `go.mod` 中的 SDK 一致：

```fish
go mod download github.com/AlonaBot/AlonaBot
go -C (go list -m -f '{{.Dir}}' github.com/AlonaBot/AlonaBot) install ./cmd/alonabot
```

### 项目结构

```text
plugin.go                 插件入口、参数结构体与命令处理函数
plugin_test.go            plugintest 集成测试
plugin.yaml               插件身份、版本与来源
config_schema.json        配置声明与默认值
CHANGELOG.md              发布说明，安装后 WebUI 也会显示
scripts/release.sh        多平台打包脚本
.github/workflows/        检查与发布工作流
```

在 `plugin.go` 的 `New` 中调用 `p.Command("名字").Description(...).Handle(处理函数)` 注册命令。处理函数签名为 `func(c *sdk.Context, args 参数结构体) (message.Content, error)`，返回内容由 Core 作为回复发送；参数字段按顺序对应命令参数，指针字段为可选参数。

配置通过 `c.Config()` 读取，例如 `c.Config().Object("greet").String("greeting")`。新增命令时，同步修改配置声明与 `plugin_test.go`。

```sh
alonabot run --reload .   # 启动 Core 并运行本插件，改动源码后自动重建
alonabot check .          # 检查插件声明与配置
go test ./...            # 运行 plugintest 集成测试
alonabot test .           # 在临时离线会话中加载、注册并停止插件
```

`alonabot run` 的数据写在 `data/`，已被 `.gitignore` 忽略。

## 发布

<details>
<summary>配置 GitHub Actions</summary>

工作流会检查插件、运行测试，并在推送版本标签后打包发布。仓库的 **Settings → Actions → General** 中应启用 GitHub Actions。

</details>

1. 修改 `plugin.yaml` 的 `version`，在 `CHANGELOG.md` 中把 `[Unreleased]` 的内容移到新版本一节，例如 `## [0.2.0] - 2026-12-01`。
2. 提交修改后，为该版本打标签并推送：

   ```sh
   git tag v0.2.0
   git push origin v0.2.0
   ```

3. `Release` 工作流核对标签与 `version`、`source` 与仓库，为 linux/amd64、linux/arm64、darwin/amd64、darwin/arm64、windows/amd64 打包，一次生成共享的 `release.json`，并以 CHANGELOG 中该版本的内容创建 [GitHub Release](https://github.com/owner/alonabot-plugin-template/releases)。

当前工作流发布未签名包，安装界面会标明来源未签名。CLI 的 `pack --key` 可为五个平台共享的 `release.json` 统一签名；签名密钥由发布者保管。
