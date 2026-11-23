# Using the generated plugin

[简体中文](usage.zh-CN.md)

Complete [template setup](template-setup.md), then run this on the machine running Core v0.2.0 or later:

```fish
alonabot install --source owner/alonabot-plugin-template
```

You can also install from **Plugins → Store** in WebUI. An administrator confirms source trust on first installation.

The example supplies `hello [name]` and reads the configured greeting. Omitting the name replies `Hello, everyone!`; `/hello Alona` replies `Hello, Alona!`. Core sets the prefix, which defaults to `/`, for example `/hello 阿洛娜`.

Save configuration in the WebUI plugin details page to reload the plugin:

| Setting | Type | Default | Purpose |
| --- | --- | --- | --- |
| `greet.greeting` | string | `Hello` | Greeting before the name |

Replace these configuration and command descriptions with the generated plugin's actual behavior. See [development](development.md) and [releases](releasing.md).
