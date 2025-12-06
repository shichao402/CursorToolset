# CursorToolset 包开发指南

本指南帮助你开发符合 CursorToolset 规范的工具集包。

---

## 目录结构

```
my-toolset/
├── toolset.json          # 包配置文件（必需）
├── .cursortoolset/       # AI 规则目录
│   └── rules/            # 规则文件
├── rules/                # 你的规则文件
├── .gitignore
└── README.md
```

---

## toolset.json 规范

这是包的核心配置文件，定义包的元数据和下载信息。

```json
{
  "name": "my-toolset",
  "displayName": "My Toolset",
  "version": "1.0.0",
  "description": "包的简短描述",
  "author": "你的名字",
  "license": "MIT",
  "keywords": ["keyword1", "keyword2"],
  "repository": {
    "type": "git",
    "url": "https://github.com/USERNAME/my-toolset.git"
  },
  "dist": {
    "tarball": "https://github.com/USERNAME/my-toolset/releases/download/v1.0.0/my-toolset-1.0.0.tar.gz",
    "sha256": "SHA256校验和"
  },
  "cursortoolset": {
    "minVersion": "1.0.0"
  }
}
```

### 字段说明

| 字段 | 必需 | 说明 |
|------|------|------|
| name | ✅ | 包名，小写字母、数字、连字符 |
| version | ✅ | 语义化版本号 (MAJOR.MINOR.PATCH) |
| displayName | ❌ | 显示名称 |
| description | ❌ | 包描述 |
| dist.tarball | ✅ | 下载地址 |
| dist.sha256 | ✅ | SHA256 校验和 |

---

## 版本号规范

遵循语义化版本 (SemVer)：

- **MAJOR**: 不兼容的 API 变更
- **MINOR**: 向后兼容的功能新增
- **PATCH**: 向后兼容的问题修复

示例：`1.0.0`, `1.2.3`, `2.0.0`

---

## 发布流程

### 1. 更新版本号

编辑 `toolset.json`，更新 `version` 和 `dist.tarball` 中的版本号。

### 2. 打包

使用 CursorToolset 内置命令打包：

```bash
cursortoolset pack
```

或手动打包：

```bash
tar -czvf my-toolset-1.0.0.tar.gz --exclude='.git' --exclude='*.tar.gz' *
```

### 3. 计算 SHA256

```bash
shasum -a 256 my-toolset-1.0.0.tar.gz
```

将结果更新到 `toolset.json` 的 `dist.sha256` 字段。

### 4. 创建 Git Tag

```bash
git add .
git commit -m "chore: release v1.0.0"
git tag v1.0.0
git push origin main --tags
```

### 5. 发布 GitHub Release

1. 在 GitHub 仓库创建 Release
2. Tag 选择 `v1.0.0`
3. 上传打包文件 `my-toolset-1.0.0.tar.gz`
4. 发布

### 6. 提交到包索引（可选）

如果希望包被公开发现，可以提交 PR 到 CursorToolset 的 registry。

---

## AI 规则编写

在 `.cursortoolset/rules/` 或 `rules/` 目录下创建 `.md` 文件作为 AI 规则。

### 规则文件示例

```markdown
# 项目开发规范

## 代码风格
- 使用 4 空格缩进
- 函数命名使用驼峰式

## 提交规范
- feat: 新功能
- fix: 修复
- docs: 文档
```

### 最佳实践

1. **清晰明确** - 规则应该具体、可执行
2. **分类组织** - 按主题拆分多个规则文件
3. **保持更新** - 随项目演进更新规则

---

## 常用命令

```bash
# 初始化新包
cursortoolset init my-toolset

# 打包
cursortoolset pack

# 打包并更新 SHA256
cursortoolset pack --verify

# 本地安装测试
cursortoolset install ./my-toolset
```

---

## 参考资源

- [CursorToolset 仓库](https://github.com/shichao402/CursorToolset)
- [语义化版本规范](https://semver.org/lang/zh-CN/)
