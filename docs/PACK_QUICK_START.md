# 📦 打包快速开始

5 分钟学会标准化打包你的工具集包。

## 前置条件

- ✅ 已安装 CursorToolset
- ✅ 有一个包项目（或使用 `cursortoolset init` 创建）

## 完整流程

### 步骤 1: 准备项目

如果还没有项目，先初始化：

```bash
cursortoolset init my-awesome-toolset
cd my-awesome-toolset
```

### 步骤 2: 开发你的工具集

```bash
# 创建规则文件
mkdir -p rules
cat > rules/coding-style.md << 'EOF'
# Coding Style Rules

Always use camelCase for variables.
EOF

# 创建 README
cat > README.md << 'EOF'
# My Awesome Toolset

This is an awesome toolset for development.

## Installation

\`\`\`bash
cursortoolset install my-awesome-toolset
\`\`\`
EOF
```

### 步骤 3: 配置包信息

编辑 `toolset.json`：

```json
{
  "name": "my-awesome-toolset",
  "displayName": "My Awesome Toolset",
  "version": "1.0.0",
  "description": "An awesome toolset for developers",
  "author": "Your Name",
  "license": "MIT",
  "keywords": ["development", "toolset", "cursor"],
  
  "repository": {
    "type": "git",
    "url": "https://github.com/yourusername/my-awesome-toolset.git"
  },
  
  "dist": {
    "tarball": "https://github.com/yourusername/my-awesome-toolset/releases/download/v1.0.0/my-awesome-toolset-1.0.0.tar.gz",
    "sha256": ""
  },
  
  "cursortoolset": {
    "minVersion": "1.0.0"
  }
}
```

### 步骤 4: 标准化打包 🎯

```bash
# 一键打包并自动更新 SHA256
cursortoolset pack --verify
```

输出示例：
```
📦 标准化打包工具集包
   目录: /Users/dev/my-awesome-toolset

✅ 验证通过: my-awesome-toolset v1.0.0

📋 收集到 6 个文件

🔨 创建压缩包: my-awesome-toolset-1.0.0.tar.gz

🔐 计算 SHA256...

✅ 打包完成！

📦 文件: my-awesome-toolset-1.0.0.tar.gz
📏 大小: 2.1 KB
🔐 SHA256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

🔄 更新 toolset.json 中的 sha256...
✅ 已更新 toolset.json

💡 下一步：
   1. 在 GitHub 创建 Release (v1.0.0)
   2. 上传 my-awesome-toolset-1.0.0.tar.gz 到 Release
```

### 步骤 5: 发布到 GitHub

```bash
# 1. 提交更改
git add .
git commit -m "Release v1.0.0"

# 2. 创建标签
git tag v1.0.0

# 3. 推送
git push origin main
git push origin v1.0.0
```

### 步骤 6: 创建 GitHub Release

1. 访问你的 GitHub 仓库
2. 点击 "Releases" → "Create a new release"
3. 选择标签 `v1.0.0`
4. 填写发布说明
5. 上传 `my-awesome-toolset-1.0.0.tar.gz`
6. 点击 "Publish release"

### 步骤 7: 更新下载地址

复制 Release 中的文件下载地址，更新 `toolset.json`：

```json
{
  "dist": {
    "tarball": "https://github.com/yourusername/my-awesome-toolset/releases/download/v1.0.0/my-awesome-toolset-1.0.0.tar.gz",
    "sha256": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
  }
}
```

提交更新：

```bash
git add toolset.json
git commit -m "Update download URL for v1.0.0"
git push
```

## 完成！🎉

现在其他用户可以通过以下方式安装你的工具集：

```bash
# 如果已添加到 Registry
cursortoolset install my-awesome-toolset

# 或直接指定 URL
cursortoolset install https://github.com/yourusername/my-awesome-toolset/releases/download/v1.0.0/my-awesome-toolset-1.0.0.tar.gz
```

## 常用命令速查

```bash
# 打包（自动命名）
cursortoolset pack

# 打包并自动更新 SHA256
cursortoolset pack --verify

# 指定输出文件名
cursortoolset pack --output custom-name.tar.gz

# 排除额外文件
cursortoolset pack --exclude "*.tmp" --exclude "test-data"

# 打包指定目录
cursortoolset pack ./my-project
```

## 发布检查清单

- [ ] 更新 `toolset.json` 中的版本号
- [ ] 确保 README.md 完整
- [ ] 运行 `cursortoolset pack --verify`
- [ ] 测试生成的包（解压验证）
- [ ] 创建 Git Tag
- [ ] 在 GitHub 创建 Release
- [ ] 上传 tar.gz 文件
- [ ] 更新 dist.tarball 地址
- [ ] 提交最终版本

## 下次更新流程

```bash
# 1. 修改代码...

# 2. 更新版本号
vim toolset.json  # 改为 1.0.1 或 1.1.0

# 3. 打包
cursortoolset pack --verify

# 4. 发布
git add .
git commit -m "Release v1.1.0"
git tag v1.1.0
git push origin main v1.1.0

# 5. 在 GitHub 创建新的 Release
```

## 疑难解答

### 问题 1: "包名格式不正确"

**原因**: 包名包含大写字母、下划线或特殊字符

**解决**: 只使用小写字母、数字和连字符
```
✅ my-toolset
✅ dev-utils
❌ My-Toolset (大写)
❌ my_toolset (下划线)
```

### 问题 2: "版本号格式不正确"

**原因**: 版本号不符合语义化版本规范

**解决**: 使用 `MAJOR.MINOR.PATCH` 格式
```
✅ 1.0.0
✅ 2.1.3
❌ 1.0 (缺少 PATCH)
❌ v1.0.0 (不要加 v 前缀)
```

### 问题 3: "打包文件太大"

**原因**: 包含了不必要的文件

**解决**: 使用 `--exclude` 排除
```bash
cursortoolset pack --exclude "*.mp4" --exclude "large-data"
```

### 问题 4: SHA256 不匹配

**原因**: 文件被修改或下载不完整

**解决**: 重新打包
```bash
cursortoolset pack --verify
```

## 高级技巧

### 技巧 1: 配置可执行程序

如果你的包包含可执行程序：

```json
{
  "bin": {
    "mytool": "bin/mytool.sh",
    "helper": "scripts/helper"
  }
}
```

### 技巧 2: CI/CD 自动化

创建 `.github/workflows/release.yml`：

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
          curl -sSL https://raw.githubusercontent.com/shichao402/CursorToolset/main/scripts/install.sh | bash
      
      - name: Pack
        run: |
          $HOME/.cursortoolsets/bin/cursortoolset pack --verify
          
      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: |
            *.tar.gz
```

### 技巧 3: 本地测试

在发布前测试包：

```bash
# 1. 打包
cursortoolset pack

# 2. 创建测试目录
mkdir /tmp/test-pack
cd /tmp/test-pack

# 3. 解压
tar -xzf ~/my-awesome-toolset/my-awesome-toolset-1.0.0.tar.gz

# 4. 验证内容
ls -la
cat toolset.json
```

## 相关文档

- [完整打包文档](PACK_FEATURE.md) - 详细功能说明
- [包开发指南](PACKAGE_DEV.md) - 包开发最佳实践
- [Bin 功能](BIN_FEATURE.md) - 可执行程序配置
- [配置示例](examples/README.md) - 各种配置示例

## 获取帮助

```bash
# 查看帮助
cursortoolset pack --help

# 查看版本
cursortoolset version

# 提交 Issue
# https://github.com/shichao402/CursorToolset/issues
```

---

**Happy Packaging! 🎉**
