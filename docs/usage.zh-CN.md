# 使用生成后的插件

[English](usage.md)

先完成[模板替换](template-setup.zh-CN.md)。在 Core v0.2.0 或更高版本的运行机器上执行：

```fish
alonabot install --source owner/alonabot-plugin-template
```

也可从 WebUI「插件 → 商店」安装，首次安装由管理员确认来源信任。

示例提供 `hello [名字]` 命令，读取配置问候语。省略名字回复 `Hello, everyone!`，`/hello Alona` 回复 `Hello, Alona!`。前缀由 Core 配置，默认 `/`，例如 `/hello 阿洛娜`。

在 WebUI 插件详情页保存配置后，插件重载生效：

| 配置 | 类型 | 默认值 | 用途 |
| --- | --- | --- | --- |
| `greet.greeting` | string | `Hello` | 名字前的问候语 |

将本页配置与命令说明替换为生成后插件的实际功能。开发与打包见[开发指南](development.zh-CN.md)和[发布指南](releasing.zh-CN.md)。
