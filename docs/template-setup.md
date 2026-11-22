# Create a plugin from the template

[简体中文](template-setup.zh-CN.md)

1. Create a repository with [Use this template](https://github.com/new?template_owner=AlonaBot&template_name=alonabot-plugin-template).
2. Replace the placeholders below throughout the generated repository. Plugin IDs use lowercase letters, digits, and hyphens.

| Placeholder | Replace with | Example |
| --- | --- | --- |
| `owner` | Your GitHub username or organization | `alice` |
| `alonabot-plugin-template` | Repository name and plugin ID, retaining the `alonabot-plugin-` prefix | `alonabot-plugin-weather` |

The module in `go.mod`, source in `plugin.yaml`, bilingual README badges, installation commands, and repository links must point to the new repository. Replace owner assignments in Issue forms and triage too. Retain the SDK path `github.com/AlonaBot/AlonaBot` and this page's template creation link.

3. Update `name`, `description`, and `author` in `plugin.yaml`, along with both README titles, summaries, configuration, and command descriptions. If you use a Go version badge, keep it aligned with `go.mod`.
4. Add `alonabot` and `alonabot-plugin` under **About → Settings → Topics** on the repository homepage.
5. Write handlers in `plugin.go`, register them in `New`, construct response text directly, and describe configuration in `config_schema.json`. See [development](development.md).
6. Run `go mod tidy` and the development checks, update the generated repository's `CHANGELOG.md`, then remove template-only onboarding from its homepage while retaining plugin usage links.

To add an open-source license, use **Add file → Create new file → LICENSE → Choose a license template** in the generated repository and update both license descriptions. The template itself retains owner/component placeholders.
