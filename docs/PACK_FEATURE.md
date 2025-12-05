# 📦 标准化打包功能文档

## 概述

`cursortoolset pack` 命令提供了标准化的工具集包打包功能，确保生成的包符合 CursorToolset 规范。

## 功能特性

- ✅ **配置验证** - 自动验证 `toolset.json` 的完整性和正确性
- ✅ **智能排除** - 自动排除不需要的文件（.git、node_modules 等）
- ✅ **标准格式** - 生成符合规范的 tar.gz 压缩包
- ✅ **SHA256 计算** - 自动计算并显示文件校验和
- ✅ **自动更新** - 可选自动更新 toolset.json 中的 sha256 字段
- ✅ **友好提示** - 清晰的打包过程和发布指南

## 使用方法

### 基本用法

```bash
# 打包当前目录
cursortoolset pack

# 打包指定目录
cursortoolset pack ./my-toolset

# 指定输出文件名
cursortoolset pack --output my-toolset-1.0.0.tar.gz

# 打包并自动更新 toolset.json 中的 sha256
cursortoolset pack --verify
```

### 命令选项

| 选项 | 简写 | 说明 | 默认值 |
|------|------|------|--------|
| `--output` | `-o` | 输出文件名 | `<name>-<version>.tar.gz` |
| `--verify` | `-v` | 验证并更新 toolset.json | false |
| `--exclude` | `-e` | 额外排除的文件/目录 | [] |

### 示例

#### 1. 基本打包

```bash
$ cursortoolset pack
📦 标准化打包工具集包
   目录: /Users/dev/my-toolset

✅ 验证通过: my-toolset v1.0.0

📋 收集到 15 个文件

🔨 创建压缩包: my-toolset-1.0.0.tar.gz

🔐 计算 SHA256...

✅ 打包完成！

📦 文件: my-toolset-1.0.0.tar.gz
📏 大小: 2.3 MB
🔐 SHA256: abc123def456...

💡 下一步：
   1. 在 GitHub 创建 Release (v1.0.0)
   2. 上传 my-toolset-1.0.0.tar.gz 到 Release
   3. 复制 SHA256 到 toolset.json 的 dist.sha256 字段

   或使用 --verify 自动更新: cursortoolset pack --verify
```

#### 2. 自动更新配置

```bash
$ cursortoolset pack --verify
📦 标准化打包工具集包
   目录: /Users/dev/my-toolset

✅ 验证通过: my-toolset v1.0.0

📋 收集到 15 个文件

🔨 创建压缩包: my-toolset-1.0.0.tar.gz

🔐 计算 SHA256...

✅ 打包完成！

📦 文件: my-toolset-1.0.0.tar.gz
📏 大小: 2.3 MB
🔐 SHA256: abc123def456...

🔄 更新 toolset.json 中的 sha256...
✅ 已更新 toolset.json

💡 下一步：
   1. 在 GitHub 创建 Release (v1.0.0)
   2. 上传 my-toolset-1.0.0.tar.gz 到 Release
```

#### 3. 自定义排除

```bash
# 排除额外的文件
cursortoolset pack --exclude "*.tmp" --exclude "test-data"
```

## 打包流程

### 1. 验证阶段

系统会验证以下内容：

- ✅ `toolset.json` 文件存在
- ✅ JSON 格式正确
- ✅ 必填字段完整（name、version）
- ✅ 包名格式正确（小写字母、数字、连字符）
- ✅ 版本号符合语义化版本规范（MAJOR.MINOR.PATCH）

### 2. 文件收集阶段

自动收集目录下的所有文件，同时排除：

**默认排除的文件/目录：**
- `.git` - Git 仓库目录
- `.gitignore` - Git 配置文件
- `.DS_Store` - macOS 系统文件
- `Thumbs.db` - Windows 系统文件
- `node_modules` - Node.js 依赖目录
- `.idea`、`.vscode` - IDE 配置目录
- `*.swp`、`*.swo` - Vim 临时文件
- `*.log` - 日志文件
- `dist` - 构建产物目录
- `*.tar.gz`、`*.zip` - 压缩包文件

**可自定义排除：**
```bash
cursortoolset pack --exclude "docs/draft" --exclude "*.bak"
```

### 3. 压缩打包阶段

- 使用 `tar.gz` 格式压缩
- 保持文件权限和目录结构
- 使用相对路径（不包含父目录路径）

### 4. 校验和计算

- 使用 SHA256 算法
- 计算整个压缩包的校验和
- 用于后续的安全验证

### 5. 配置更新（可选）

使用 `--verify` 选项时：
- 自动将计算的 SHA256 写入 `toolset.json`
- 更新 `dist.sha256` 字段
- 保持原有的 JSON 格式

## 完整发布流程

### 准备阶段

```bash
# 1. 确保 toolset.json 配置正确
cat toolset.json

# 2. 更新版本号（如果需要）
vim toolset.json  # 修改 version 字段

# 3. 测试你的工具集
# ... 运行测试 ...
```

### 打包阶段

```bash
# 标准化打包
cursortoolset pack --verify

# 输出示例：my-toolset-1.0.0.tar.gz
```

### 发布阶段

```bash
# 1. 提交代码（如果有修改）
git add .
git commit -m "Release v1.0.0"

# 2. 创建 Git Tag
git tag v1.0.0
git push origin v1.0.0

# 3. 在 GitHub 创建 Release
# - 标签：v1.0.0
# - 标题：v1.0.0
# - 上传：my-toolset-1.0.0.tar.gz

# 4. 更新 toolset.json 中的 dist.tarball
vim toolset.json
# 修改为：https://github.com/USER/REPO/releases/download/v1.0.0/my-toolset-1.0.0.tar.gz

# 5. 提交 toolset.json 的更新
git add toolset.json
git commit -m "Update download URL for v1.0.0"
git push
```

### 提交到 Registry（可选）

如果要让其他用户可以通过 `cursortoolset install` 安装：

```bash
# 提交 PR 到 CursorToolset Registry
# 在 registry.json 中添加你的包信息
```

## 配置文件要求

### toolset.json 必填字段

```json
{
  "name": "my-toolset",        // 包名（必须）
  "version": "1.0.0",          // 版本号（必须）
  "description": "描述",        // 描述（推荐）
  "author": "作者",            // 作者（推荐）
  "dist": {
    "tarball": "https://...",  // 下载地址（发布后填写）
    "sha256": "abc123..."      // SHA256（pack --verify 自动生成）
  }
}
```

### 版本号规范

必须遵循语义化版本规范（SemVer）：

- **格式**: `MAJOR.MINOR.PATCH`
- **示例**: `1.0.0`、`2.1.3`、`0.9.0`
- **说明**:
  - MAJOR: 不兼容的 API 修改
  - MINOR: 向下兼容的功能新增
  - PATCH: 向下兼容的问题修复

### 包名规范

- 只能包含小写字母、数字和连字符
- 不能以连字符开头或结尾
- 推荐使用有意义的名称

**合法示例：**
- `my-toolset`
- `cursor-rules`
- `dev-utils`

**不合法示例：**
- `My-Toolset` ❌ (包含大写)
- `-my-toolset` ❌ (以连字符开头)
- `my_toolset` ❌ (包含下划线)

## 目录结构建议

```
my-toolset/
├── toolset.json              # 包配置文件（必需）
├── README.md                 # 说明文档（推荐）
├── LICENSE                   # 许可证（推荐）
├── .cursortoolset/           # 开发规则目录
│   └── rules/
│       └── dev-guide.md
├── rules/                    # AI 规则文件（核心内容）
│   ├── coding-style.md
│   └── best-practices.md
├── bin/                      # 可执行程序（可选）
│   └── mytool
└── scripts/                  # 辅助脚本（可选）
    └── helper.sh
```

## 常见问题

### Q1: 打包后文件太大怎么办？

**A**: 检查是否包含了不必要的文件：

```bash
# 查看压缩包内容
tar -tzf my-toolset-1.0.0.tar.gz

# 如果包含不必要的文件，使用 --exclude 排除
cursortoolset pack --exclude "large-data" --exclude "*.mp4"
```

### Q2: 如何验证打包结果？

**A**: 解压并检查内容：

```bash
# 创建测试目录
mkdir test-unpack
cd test-unpack

# 解压
tar -xzf ../my-toolset-1.0.0.tar.gz

# 检查内容
ls -la
cat toolset.json
```

### Q3: SHA256 不匹配怎么办？

**A**: 重新打包并更新：

```bash
# 重新打包并自动更新 sha256
cursortoolset pack --verify

# 或手动更新 toolset.json
# 复制显示的 SHA256 到 dist.sha256 字段
```

### Q4: 可以打包二进制文件吗？

**A**: 可以！包中可以包含任何类型的文件：

```json
{
  "bin": {
    "mytool": "bin/mytool",
    "helper": "bin/helper.exe"
  }
}
```

### Q5: 如何跨平台发布？

**A**: 为不同平台创建不同的包：

```bash
# Linux 版本
cursortoolset pack --output my-toolset-1.0.0-linux.tar.gz

# macOS 版本
cursortoolset pack --output my-toolset-1.0.0-darwin.tar.gz

# Windows 版本
cursortoolset pack --output my-toolset-1.0.0-windows.tar.gz
```

在 toolset.json 中配置多个平台：

```json
{
  "dist": {
    "linux": {
      "tarball": "https://.../my-toolset-1.0.0-linux.tar.gz",
      "sha256": "..."
    },
    "darwin": {
      "tarball": "https://.../my-toolset-1.0.0-darwin.tar.gz",
      "sha256": "..."
    }
  }
}
```

## 最佳实践

### 1. 版本管理

- ✅ 每次发布前更新版本号
- ✅ 使用 Git Tag 标记版本
- ✅ 遵循语义化版本规范
- ✅ 在 CHANGELOG.md 中记录变更

### 2. 文件组织

- ✅ 保持目录结构清晰
- ✅ 将规则文件放在 `rules/` 目录
- ✅ 可执行程序放在 `bin/` 目录
- ✅ 辅助脚本放在 `scripts/` 目录

### 3. 文档完整

- ✅ README.md 包含使用说明
- ✅ LICENSE 明确许可证
- ✅ 规则文件有清晰的注释

### 4. 测试验证

- ✅ 打包前测试功能
- ✅ 解压验证内容完整性
- ✅ 在干净环境中测试安装

### 5. 发布流程

```bash
# 1. 准备发布
vim toolset.json  # 更新版本号
make test         # 运行测试

# 2. 打包
cursortoolset pack --verify

# 3. 提交代码
git add .
git commit -m "Release v1.0.0"
git tag v1.0.0
git push origin v1.0.0

# 4. 创建 Release
# 在 GitHub 上创建 Release 并上传压缩包

# 5. 更新下载地址
vim toolset.json  # 更新 dist.tarball
git commit -am "Update download URL"
git push
```

## 集成到 CI/CD

### GitHub Actions 示例

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
      - uses: actions/checkout@v3
      
      - name: Install CursorToolset
        run: |
          # 安装 cursortoolset
          curl -sSL https://... | bash
      
      - name: Pack
        run: |
          cursortoolset pack --verify
          
      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: |
            *.tar.gz
```

## 安全注意事项

1. **不要包含敏感信息**
   - 密钥、密码、Token
   - 个人身份信息
   - 内部 API 地址

2. **验证 SHA256**
   - 发布后验证下载的文件
   - 确保 SHA256 匹配

3. **权限控制**
   - 不要在包中包含高权限脚本
   - 二进制文件应该来自可信源

## 技术细节

### 压缩格式

- **格式**: tar.gz (gzip 压缩的 tar 归档)
- **压缩级别**: 默认（level 6）
- **兼容性**: 所有 Unix-like 系统和 Windows (with tools)

### SHA256 计算

- **算法**: SHA256 (SHA-2 family)
- **输出**: 64 位十六进制字符串
- **用途**: 验证文件完整性和安全性

### 文件路径

- **格式**: 使用相对路径（不包含父目录）
- **分隔符**: 统一使用 `/`（跨平台兼容）
- **权限**: 保持原有文件权限

## 相关命令

- `cursortoolset init` - 初始化新包项目
- `cursortoolset install` - 安装包
- `cursortoolset publish` - 发布到 Registry（规划中）

## 参考资料

- [包开发指南](PACKAGE_DEV.md)
- [配置规范](../README.md#toolsetjson-规范)
- [语义化版本规范](https://semver.org/lang/zh-CN/)
