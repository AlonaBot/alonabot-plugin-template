#!/usr/bin/env bash
# 为五个目标平台打包，合并 release.json，并从 CHANGELOG.md 取出本版本的说明。
# 用法：scripts/release.sh v0.1.0 owner/repo
# 输出：dist/ 下的归档、release.json 与 notes.md。
set -euo pipefail

tag="$1"
repo="$2"
targets=(linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64)

id="$(sed -n 's/^id: *//p' plugin.yaml)"
version="$(sed -n 's/^version: *//p' plugin.yaml)"
source="$(sed -n 's/^source: *//p' plugin.yaml)"

if [[ "$tag" != "v$version" ]]; then
  echo "tag $tag does not match plugin.yaml version $version" >&2
  exit 1
fi
if [[ "$source" != "github.com/$repo" ]]; then
  echo "plugin.yaml source $source does not match github.com/$repo" >&2
  exit 1
fi

rm -rf build dist
mkdir -p dist
for target in "${targets[@]}"; do
  os="${target%/*}"
  arch="${target#*/}"
  name="$id-$version-$os-$arch.tar.gz"
  alonabot pack --dir . --goos "$os" --goarch "$arch" --out "build/$os-$arch" \
    --asset-url "https://github.com/$repo/releases/download/$tag/$name"
  mv "build/$os-$arch/$name" dist/
done

# 每次 pack 只写一个平台的 release.json，这里把各平台的 assets 合并成一份。
jq -s '.[0] + {assets: (map(.assets) | add)}' build/*/release.json > dist/release.json
rm -rf build

# 取出 CHANGELOG.md 中「## [版本]」到下一个「## [」之间的内容作为发布说明。
awk -v v="$version" '
  $0 ~ "^## \\[" v "\\]" { found = 1; next }
  found && /^## \[/ { exit }
  found { print }
' CHANGELOG.md > dist/notes.md
if [[ ! -s dist/notes.md ]]; then
  echo "CHANGELOG.md has no section for $version" >&2
  exit 1
fi
