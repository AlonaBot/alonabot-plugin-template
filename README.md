<div align="center">

# alonabot-plugin-template

_AlonaBot 插件模板_

</div>

## 介绍

这是 [AlonaBot](https://github.com/AlonaBot/AlonaBot) 的插件模板，包含：

- 一个可以直接运行、按功能分模块的示例插件：`greet` 模块的 `hello` 命令读取配置的问候语，`dice` 模块的 `roll` 命令（别名 `dice`）带一个可选参数。
- 配置声明 `config_schema.json`、中英文案 `i18n/`、插件身份 `plugin.yaml`。
- 用 `plugintest` 写的集成测试，以及不依赖 SDK 的业务逻辑单元测试。
- 检查与测试的 CI，以及推送版本标签后自动打包五个平台、发布 GitHub Release 的工作流。
- `CHANGELOG.md`：发布说明取自这里，安装后 WebUI 也会显示。

## 用模板开始

1. 在 GitHub 上点「Use this template」创建自己的仓库。
2. 修改插件身份：
   - `go.mod` 的 `module` 改成你的仓库路径；
   - `plugin.yaml` 的 `id`（市场 ID 以 `alonabot_plugin_` 开头）、`name`、`description`、`author`，以及 `source`（`github.com/<你的账号>/<仓库名>`，发布时必须与仓库一致）；
   - `i18n/*.json` 的 `metadata`；
   - `LICENSE` 中的版权人。
3. 把 `plugin.go` 与 `internal/*` 中的 import 路径改成新的 module，运行 `go mod tidy`。
4. 照「项目结构」删改示例模块，加入你自己的功能。

## 项目结构

```text
plugin.go                 入口：New 只组装各功能模块
internal/greet/greet.go   一个功能一个包：子插件、命令与配置
internal/dice/dice.go     业务逻辑（Roll）与命令登记（Register）分开
internal/dice/dice_test.go  业务逻辑的单元测试，不启动 Core
plugin_test.go            plugintest 集成测试，走一遍各模块的命令
config_schema.json        每个模块一个同名 object：greet、dice
i18n/                     文案键按模块加前缀：greet.*、dice.*
```

每个功能模块：

- 在 `internal/` 下建一个包，导出 `Register(p *sdk.Plugin)`，在其中用 `p.Sub(ID)` 创建子插件并登记命令；子插件可以统一设置缺省权限（示例为 `sdk.User`）。
- 配置写成 `config_schema.json` 中与子插件 ID 同名的 `object`，用 `s.Config()` 读取本模块的配置，模块之间互不读取对方的配置。
- 不依赖 SDK 的业务逻辑写成普通函数，用 Go 的单元测试覆盖；命令的整条路径由 `plugin_test.go` 用 `plugintest` 覆盖。
- 在 `plugin.go` 的 `New` 中调用它的 `Register`。

新增模块时照 `internal/greet` 复制一份，改 `ID`、命令、配置与文案，再在 `New` 中登记。

## 安装

在运行 AlonaBot 的机器上：

```bash
alonabot install --source <你的账号>/<仓库名>
```

也可以在 WebUI 的「插件 → 商店」中安装。首次安装会请管理员确认信任这个来源。

## 配置

| 配置项 | 类型 | 缺省值 | 说明 |
| --- | --- | --- | --- |
| `greet.greeting` | string | `Hello` | 名字前的问候语 |
| `dice.default_sides` | int | `6` | 掷骰的默认面数，2～100 |

配置在 WebUI 的插件详情页修改，保存后插件重载生效。

## 使用

| 命令 | 参数 | 说明 |
| --- | --- | --- |
| `hello` | `[名字]` | 用问候语问候，省略名字时问候所有人 |
| `roll`、`dice` | `[面数]` | 掷一个骰子，省略面数时用 `dice.default_sides` |

命令前缀由 AlonaBot 设置，缺省为 `/`，例如 `/hello 阿洛娜`。

## 开发

需要 [Go](https://go.dev/dl/)（版本见 `go.mod`）与 AlonaBot CLI。CLI 的版本应与 `go.mod` 中的 SDK 一致：

```bash
go install "github.com/AlonaBot/AlonaBot/cmd/alonabot@$(go list -m -f '{{.Version}}' github.com/AlonaBot/AlonaBot)"
```

常用命令：

```bash
alonabot run --reload .   # 在本地启动 Core 并运行本插件，改动源码后自动重建
alonabot check --dir .    # 检查 plugin.yaml、config_schema.json 与 i18n
go test ./...             # 运行 plugintest 集成测试
alonabot test --dir .     # 在临时离线会话中加载、注册并停止插件
```

`alonabot run` 的数据写在 `data/`，已被 `.gitignore` 忽略。

## 发布

1. 修改 `plugin.yaml` 的 `version`，在 `CHANGELOG.md` 中把 `[Unreleased]` 的内容移到新版本一节，例如 `## [0.2.0] - 2026-12-01`。
2. 提交后打标签并推送：

   ```bash
   git tag v0.2.0
   git push origin v0.2.0
   ```

3. `Release` 工作流核对标签与 `version`、`source` 与仓库，为 linux/amd64、linux/arm64、darwin/amd64、darwin/arm64、windows/amd64 打包，合并出一份 `release.json`，并以 CHANGELOG 中该版本的内容创建 GitHub Release。

AlonaBot 仓库公开之前，CI 拉取 SDK 需要在仓库的 Actions secrets 中设置 `ALONABOT_SDK_TOKEN`：一个能读取 `AlonaBot/AlonaBot` 的 fine-grained token（Contents 只读）。仓库公开后可以删除这个 secret。

发布签名（`release.json.sig`）是可选的。当前 CLI 只能为单个平台的 `release.json` 签名，多平台合并后的签名等 CLI 支持后再加入工作流；没有签名时，安装界面会标明来源未签名。

## 许可

[MIT](LICENSE)
