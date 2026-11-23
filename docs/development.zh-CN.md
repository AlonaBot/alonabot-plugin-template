# 源码开发

[English](development.md)

在本仓根目录执行。使用 `go.mod` 声明的 Go 版本，本地由 asdf 管理；SDK 以该文件固定版本为准。CLI 沿用组件工具链构建：

```fish
go mod download github.com/AlonaBot/AlonaBot
env GOTOOLCHAIN=(go env GOVERSION) go -C (go list -m -f '{{.Dir}}' github.com/AlonaBot/AlonaBot) install ./cmd/alonabot
```

将 Go bin 目录加入 PATH。跨仓开发可用本地 `go.work` 指向对应 Core checkout；发行验证使用锁定的 SDK。

## 验证改动

按影响选择检查。Go、依赖和 CI 门禁变更触发以下 Go 检查；组件资源和检查工作流变更也触发 `alonabot check .`。

```fish
gofmt -l (git ls-files '*.go')
go vet ./...
go test -race ./...
for target in linux/amd64 linux/arm64 windows/amd64 darwin/amd64 darwin/arm64
    set -l parts (string split / $target)
    env GOOS=$parts[1] GOARCH=$parts[2] CGO_ENABLED=0 go build ./...
end
alonabot check .
alonabot test .
```

本仓当前没有 `go:generate` 入口，直接维护对应资源。

文档和 GitHub 元数据检查改动链接、YAML、脚本及 `git diff --check`。CI 的 Repository 比较事件 base SHA 与 HEAD；纯文档改动跳过 Go 检查。PR 记录实际命令、平台、失败与未运行项。

## 源码导航

- `plugin.go`: 公开构造入口与示例处理函数
- `plugin.yaml`: 组件身份与打包声明

[Core 源码开发](https://github.com/AlonaBot/AlonaBot/blob/main/docs/development.zh-CN.md) · [Core 插件指南](https://github.com/AlonaBot/AlonaBot/blob/main/docs/plugin-guide.zh-CN.md)

## 注册命令与读取配置

在 `plugin.go` 的 `New` 中调用 `p.Command("名字").Description(...).Menu().Handle(处理函数)`。`.Menu()` 将命令加入自动菜单。处理函数签名为 `func(c *sdk.Context, args 参数结构体) (message.Content, error)`，返回内容由 Core 作为回复发送；参数字段按顺序对应命令参数，指针字段为可选参数。

配置经 `c.Config()` 读取，例如 `c.Config().Object("greet").String("greeting")`。新增命令同步更新配置声明和 `plugin_test.go`。

```fish
alonabot run --reload .
```

该命令启动 Core 与本插件，源码变化后重建。运行数据保存于已忽略的 `data/`。`alonabot test .` 在临时离线会话加载、注册并停止插件。`plugin_test.go` 为 plugintest 集成测试，`scripts/release.sh` 负责多平台打包，`.github/workflows/` 提供检查与发布。在生成后仓库的 **Settings → Actions → General** 启用 GitHub Actions。
