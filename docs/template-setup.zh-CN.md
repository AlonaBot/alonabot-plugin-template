# 从模板创建插件

[English](template-setup.md)

1. 点击[使用此模板](https://github.com/new?template_owner=AlonaBot&template_name=alonabot-plugin-template)创建仓库。
2. 在新仓全局替换下表占位，插件 ID 使用小写字母、数字和连字符。

| 占位 | 替换内容 | 示例 |
| --- | --- | --- |
| `owner` | GitHub 用户名或组织名 | `alice` |
| `alonabot-plugin-template` | 仓库名与插件 ID，保留 `alonabot-plugin-` 前缀 | `alonabot-plugin-weather` |

替换后，`go.mod` 的 module、`plugin.yaml` 的 source、中英 README 徽章、安装与仓库链接应共同指向新仓；Issue 与 triage 的 owner 指派同步替换。保留 SDK 的 `github.com/AlonaBot/AlonaBot` 路径与本页模板创建链接。

3. 修改 `plugin.yaml` 的 `name`、`description`、`author`，以及双语 README 的标题、简介、配置与命令说明；如使用 Go 版本徽章，与 `go.mod` 保持一致。
4. 在仓库首页 **About → 设置 → Topics** 添加 `alonabot` 和 `alonabot-plugin`。
5. 处理函数写在 `plugin.go`，通过 `New` 注册命令；直接构造回复文本，配置说明放在 `config_schema.json`。见[源码开发](development.zh-CN.md)。
6. 运行 `go mod tidy` 及开发检查，更新生成后仓库的 `CHANGELOG.md`，再从首页移除模板专用引导，保留插件使用入口。

如需开源许可证，在新仓通过 **Add file → Create new file → LICENSE → Choose a license template** 选择，并更新双语许可说明。模板自身保持 owner/组件身份占位。
