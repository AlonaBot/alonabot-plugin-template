<div align="center">

# alonabot-plugin-template

_AlonaBot 插件模板_

</div>

## 介绍

这是 [AlonaBot](https://github.com/AlonaBot/AlonaBot) 的插件模板，包含：

- 一个可以直接运行的示例插件：命令在入口集中注册，处理函数按功能分文件；`hello` 命令读取配置的问候语，带一个可选参数。
- 配置声明 `config_schema.json`、中英文案 `i18n/`、插件身份 `plugin.yaml`。
- 用 `plugintest` 写的集成测试。
- 检查与测试的 CI，以及推送版本标签后自动打包五个平台、发布 GitHub Release 的工作流。
- `CHANGELOG.md`：发布说明取自这里，安装后 WebUI 也会显示。

## 用模板开始

1. 在 GitHub 上点「Use this template」创建自己的仓库。
2. 修改插件身份：
   - `go.mod` 的 `module` 改成你的仓库路径；
   - `plugin.yaml` 的 `id`（市场 ID 以 `alonabot_plugin_` 开头）、`name`、`description`、`author`，以及 `source`（`github.com/<你的账号>/<仓库名>`，发布时必须与仓库一致）；
   - `i18n/*.json` 的 `metadata`；
   - `LICENSE` 中的版权人。
3. 把各文件 import 中的 `github.com/AlonaBot/alonabot-plugin-template` 改成新的 module，运行 `go mod tidy`。
4. 照「项目结构」删改示例命令，加入你自己的功能。

## 项目结构

```text
plugin.go                 入口：全部命令在 New 中集中注册
handler/greet.go          hello 的处理函数
plugin_test.go            plugintest 集成测试，走一遍全部命令
config_schema.json        配置按功能分组，例如 greet
i18n/                     文案键按功能加前缀，例如 greet.*
```

写法与 gin 的路由类似：

- **集中注册**：`plugin.go` 的 `New` 里一行注册一个命令，`p.Command("名字").Description(...).Handle(处理函数)`，打开这个文件就能看到插件的全部命令。
- **分散实现**：处理函数放在 `handler/`，一个功能一个文件，签名是 `func(c *sdk.Context, args 参数结构体) (message.Content, error)`；返回的内容由 Core 作为回复发送。参数结构体的字段按顺序对应命令参数，指针字段是可选参数。
- **配置**：在处理函数中用 `c.Config()` 读取，例如 `c.Config().Object("greet").String("greeting")`。
- **测试**：命令的整条路径由 `plugin_test.go` 用 `plugintest` 覆盖；功能变复杂时，可以把不依赖 SDK 的逻辑放进 `internal/` 下的包，用普通的 Go 单元测试覆盖。

新增一个命令：在 `handler/` 写处理函数，在 `config_schema.json` 与 `i18n/` 中加上它用到的配置和文案，再在 `New` 中加一行注册。

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

配置在 WebUI 的插件详情页修改，保存后插件重载生效。

## 使用

| 命令 | 参数 | 说明 |
| --- | --- | --- |
| `hello` | `[名字]` | 用问候语问候，省略名字时问候所有人 |

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
