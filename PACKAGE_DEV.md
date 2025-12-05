# CursorToolset 包开发指南

本文档介绍如何开发和发布 CursorToolset 包。

## 设计理念

CursorToolset 采用类似 Homebrew/pip 的简洁设计：

- **只下载，不执行** - 管理器只负责下载和解压，不执行任何脚本
- **声明式配置** - 包通过 `toolset.json` 自描述
- **中心化索引** - 通过 Registry 统一管理包列表

---

## 快速开始

### 1. 初始化包项目

```bash
cursortoolset init my-toolset --author "Your Name"
cd my-toolset
```

生成的目录结构：

```
my-toolset/
├── toolset.json          # 包配置文件（必需）
├── README.md             # 包说明文档
└── .gitignore
```

### 2. 添加内容

```
my-toolset/
├── toolset.json
├── README.md
├── rules/                # AI 规则文件
│   ├── coding-style.md
│   └── best-practices.md
└── prompts/              # 提示词模板（可选）
    └── code-review.md
```

### 3. 配置 toolset.json

```json
{
  "name": "my-toolset",
  "displayName": "My Toolset",
  "version": "1.0.0",
  "description": "一个 AI 工具集",
  "author": "Your Name",
  "license": "MIT",
  "keywords": ["ai", "rules"],

  "repository": {
    "type": "git",
    "url": "https://github.com/yourname/my-toolset.git"
  },

  "dist": {
    "tarball": "https://github.com/yourname/my-toolset/releases/download/v1.0.0/my-toolset-1.0.0.tar.gz",
    "sha256": ""
  },

  "cursortoolset": {
    "minVersion": "1.0.0"
  }
}
```

### 4. 打包

```bash
# 打包（排除不需要的文件）
tar -czvf my-toolset-1.0.0.tar.gz \
  --exclude='.git' \
  --exclude='.DS_Store' \
  --exclude='*.tar.gz' \
  .

# 计算 SHA256
shasum -a 256 my-toolset-1.0.0.tar.gz
# 输出: abc123...  my-toolset-1.0.0.tar.gz

# 将 SHA256 填入 toolset.json 的 dist.sha256 字段
```

### 5. 发布到 GitHub

```bash
# 提交代码
git add .
git commit -m "Release v1.0.0"
git push origin main

# 创建 Tag
git tag v1.0.0
git push origin v1.0.0
```

然后在 GitHub 创建 Release，上传 `my-toolset-1.0.0.tar.gz`。

### 6. 提交到 Registry

Fork 本项目，编辑 `registry.json` 添加你的包：

```json
{
  "version": "1",
  "packages": [
    {
      "name": "my-toolset",
      "manifestUrl": "https://raw.githubusercontent.com/yourname/my-toolset/main/toolset.json"
    }
  ]
}
```

提交 Pull Request 即可。

---

## toolset.json 规范

### 字段说明

| 字段 | 必需 | 类型 | 说明 |
|------|:----:|------|------|
| `name` | ✅ | string | 包名，小写字母、数字、连字符 |
| `version` | ✅ | string | 语义化版本号 |
| `dist.tarball` | ✅ | string | 下载地址 |
| `dist.sha256` | ✅ | string | SHA256 校验和 |
| `displayName` | | string | 显示名称 |
| `description` | | string | 描述 |
| `author` | | string | 作者 |
| `license` | | string | 许可证 |
| `keywords` | | string[] | 关键词（用于搜索） |
| `repository.url` | | string | 仓库地址 |
| `dependencies` | | string[] | 依赖的包名 |
| `cursortoolset.minVersion` | | string | 最低管理器版本 |

### 完整示例

```json
{
  "name": "github-action-toolset",
  "displayName": "GitHub Action AI 工具集",
  "version": "1.0.0",
  "description": "帮助 AI 助手完成 GitHub Actions CI/CD 任务",
  "author": "shichao402",
  "license": "MIT",
  "keywords": ["github", "actions", "ci", "cd"],

  "repository": {
    "type": "git",
    "url": "https://github.com/shichao402/GithubActionAISelfBuilder.git"
  },

  "dist": {
    "tarball": "https://github.com/shichao402/GithubActionAISelfBuilder/releases/download/v1.0.0/github-action-toolset-1.0.0.tar.gz",
    "sha256": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
    "size": 102400
  },

  "cursortoolset": {
    "minVersion": "1.0.0"
  },

  "dependencies": []
}
```

---

## 版本号规范

使用 [语义化版本](https://semver.org/lang/zh-CN/)：

```
MAJOR.MINOR.PATCH
```

| 版本 | 何时增加 |
|------|----------|
| MAJOR | 不兼容的变更 |
| MINOR | 向下兼容的新功能 |
| PATCH | 向下兼容的修复 |

### 更新版本流程

1. 修改 `toolset.json` 中的 `version`
2. 更新 `dist.tarball` 中的版本号
3. 重新打包，计算新的 SHA256
4. 更新 `dist.sha256`
5. 创建新 Tag 和 Release

---

## 包内容建议

### 推荐目录结构

```
my-toolset/
├── toolset.json          # 必需
├── README.md             # 使用说明
├── rules/                # AI 规则
│   ├── general.md
│   └── specific.md
├── prompts/              # 提示词模板
│   └── templates.md
└── examples/             # 示例
    └── demo.md
```

### 规则文件编写建议

```markdown
# 规则标题

简要描述这个规则的用途。

## 适用场景

- 场景 1
- 场景 2

## 规则内容

### 1. 具体规则

详细说明...

### 2. 示例

好的做法：
```code
...
```

不好的做法：
```code
...
```
```

---

## 本地测试

```bash
# 打包
tar -czvf test.tar.gz --exclude='.git' --exclude='*.tar.gz' .

# 解压测试
mkdir -p /tmp/test-pkg
tar -xzvf test.tar.gz -C /tmp/test-pkg
ls -la /tmp/test-pkg

# 验证 SHA256
shasum -a 256 test.tar.gz
```

---

## GitHub Actions 自动发布（可选）

在你的包仓库中添加 `.github/workflows/release.yml`：

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Get version
        id: version
        run: echo "VERSION=${GITHUB_REF#refs/tags/v}" >> $GITHUB_OUTPUT

      - name: Package
        run: |
          PKG_NAME="${{ github.event.repository.name }}"
          VERSION="${{ steps.version.outputs.VERSION }}"
          tar -czvf "${PKG_NAME}-${VERSION}.tar.gz" \
            --exclude='.git' \
            --exclude='.github' \
            --exclude='*.tar.gz' \
            .

      - name: Calculate SHA256
        id: sha256
        run: |
          PKG_NAME="${{ github.event.repository.name }}"
          VERSION="${{ steps.version.outputs.VERSION }}"
          SHA256=$(shasum -a 256 "${PKG_NAME}-${VERSION}.tar.gz" | cut -d' ' -f1)
          echo "SHA256=${SHA256}" >> $GITHUB_OUTPUT
          echo "SHA256: ${SHA256}"

      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: "*.tar.gz"
          body: |
            ## SHA256
            ```
            ${{ steps.sha256.outputs.SHA256 }}
            ```
```

---

## 常见问题

### 包名有什么限制？

- 只能包含：小写字母 `a-z`、数字 `0-9`、连字符 `-`
- 不能以连字符开头或结尾
- 建议：`my-awesome-toolset`

### SHA256 怎么计算？

```bash
# macOS / Linux
shasum -a 256 package.tar.gz

# 或
openssl dgst -sha256 package.tar.gz
```

### 如何处理依赖？

```json
{
  "dependencies": ["base-toolset", "common-rules"]
}
```

管理器会在安装时自动安装依赖。

### 如何更新已发布的包？

1. 修改内容
2. 更新版本号
3. 重新打包
4. 创建新 Release
5. 更新 `toolset.json`

---

## 获取帮助

- 示例包: [GithubActionAISelfBuilder](https://github.com/shichao402/GithubActionAISelfBuilder)
- 提交 Issue: [CursorToolset Issues](https://github.com/shichao402/CursorToolset/issues)
