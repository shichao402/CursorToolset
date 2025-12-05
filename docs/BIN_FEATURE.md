# 可执行程序暴露功能 (bin)

## 功能概述

CursorToolset 支持包开发者在 `toolset.json` 中配置 `bin` 字段，用于暴露包中的可执行程序。安装包时，管理器会自动在 `~/.cursortoolsets/bin/` 目录中创建符号链接，让用户可以直接在终端中使用这些命令。

## 使用场景

适合以下场景的包：
- 提供 CLI 工具的包
- 包含脚本的工具集
- 需要在终端中快速执行的命令

## 配置方法

### 1. 在 toolset.json 中添加 bin 配置

```json
{
  "name": "my-toolset",
  "version": "1.0.0",
  "description": "我的工具集",
  
  "bin": {
    "mytool": "bin/mytool",
    "myhelper": "scripts/helper.sh",
    "mycli": "cli/main.py"
  },
  
  "dist": {
    "tarball": "https://github.com/user/my-toolset/releases/download/v1.0.0/my-toolset-1.0.0.tar.gz",
    "sha256": "abc123..."
  }
}
```

**bin 字段说明：**
- **键**：命令名称（用户在终端输入的命令）
- **值**：包内可执行文件的相对路径（相对于包根目录）

### 2. 包目录结构示例

```
my-toolset/
├── toolset.json           # 包配置文件
├── bin/                   # 可执行文件目录
│   └── mytool             # 主程序
├── scripts/               # 脚本目录
│   └── helper.sh          # 辅助脚本
└── cli/                   # CLI 工具
    └── main.py            # Python CLI
```

### 3. 确保文件可执行

**Unix/Linux/macOS:**
```bash
chmod +x bin/mytool
chmod +x scripts/helper.sh
chmod +x cli/main.py
```

**Windows:**
- 使用 `.exe` 文件
- 或使用 `.bat`/`.cmd` 脚本

## 安装效果

### 用户安装包时

```bash
$ cursortoolset install my-toolset

📦 安装 my-toolset@1.0.0
  ⬇️  下载: https://github.com/user/my-toolset/releases/download/v1.0.0/my-toolset-1.0.0.tar.gz
  ✅ SHA256 校验通过
  📂 解压到: ~/.cursortoolsets/repos/my-toolset
✅ my-toolset 安装完成

  🔗 创建可执行程序链接...
    ✅ mytool -> bin/mytool
    ✅ myhelper -> scripts/helper.sh
    ✅ mycli -> cli/main.py

  💡 将 bin 目录添加到 PATH:
    export PATH="/Users/username/.cursortoolsets/bin:$PATH"

💡 使用提示:
   详细文档: https://github.com/firoyang/CursorToolset/blob/main/USAGE_EXAMPLE.md
```

### 创建的符号链接

```bash
~/.cursortoolsets/bin/
├── mytool     -> ../repos/my-toolset/bin/mytool
├── myhelper   -> ../repos/my-toolset/scripts/helper.sh
└── mycli      -> ../repos/my-toolset/cli/main.py
```

## 用户使用

### 1. 将 bin 目录添加到 PATH

**Bash/Zsh (Linux/macOS):**
```bash
# 临时添加（仅当前会话）
export PATH="$HOME/.cursortoolsets/bin:$PATH"

# 永久添加（写入 ~/.bashrc 或 ~/.zshrc）
echo 'export PATH="$HOME/.cursortoolsets/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

**Fish Shell:**
```fish
# 临时添加
set -x PATH $HOME/.cursortoolsets/bin $PATH

# 永久添加（写入 ~/.config/fish/config.fish）
echo 'set -x PATH $HOME/.cursortoolsets/bin $PATH' >> ~/.config/fish/config.fish
```

**Windows (PowerShell):**
```powershell
# 临时添加
$env:PATH = "$env:USERPROFILE\.cursortoolsets\bin;$env:PATH"

# 永久添加（需要管理员权限）
[Environment]::SetEnvironmentVariable(
    "PATH",
    "$env:USERPROFILE\.cursortoolsets\bin;$env:PATH",
    "User"
)
```

### 2. 直接使用命令

```bash
# 使用安装的命令
$ mytool --help
$ myhelper process data.txt
$ mycli create --name test
```

## 卸载处理

卸载包时，管理器会自动清理符号链接：

```bash
$ cursortoolset uninstall my-toolset

🗑️  卸载 my-toolset
  🔗 清理可执行程序链接...
    ✅ 已删除 mytool
    ✅ 已删除 myhelper
    ✅ 已删除 mycli
✅ my-toolset 卸载完成
```

## 跨平台支持

### Unix/Linux/macOS

- 使用符号链接（`os.Symlink`）
- 自动设置执行权限（`chmod 0755`）
- 支持任何可执行文件（shell 脚本、二进制等）

### Windows

- 使用符号链接（需要 Windows Vista+）
- 自动添加 `.exe` 扩展名
- 支持 `.exe`、`.bat`、`.cmd` 文件

## 最佳实践

### 1. 命名规范

```json
{
  "bin": {
    "mytoolset": "bin/main",           // ✅ 推荐：使用包名作为主命令
    "mytoolset-helper": "bin/helper",  // ✅ 推荐：子命令带前缀
    "mt": "bin/main"                   // ⚠️  注意：简写可能冲突
  }
}
```

### 2. 使用 shebang

**Shell 脚本：**
```bash
#!/usr/bin/env bash
# scripts/helper.sh
echo "Hello from helper"
```

**Python 脚本：**
```python
#!/usr/bin/env python3
# cli/main.py
import sys
print("Hello from Python")
```

### 3. 提供帮助信息

```bash
#!/usr/bin/env bash
# bin/mytool

if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    echo "Usage: mytool [options]"
    echo "Options:"
    echo "  -h, --help    显示帮助信息"
    exit 0
fi

# 主逻辑
```

### 4. 错误处理

```bash
#!/usr/bin/env bash
set -e  # 遇到错误立即退出

if [ ! -f "config.json" ]; then
    echo "错误: 找不到 config.json" >&2
    exit 1
fi
```

## 常见问题

### Q1: 命令找不到

**问题：** 执行 `mytool` 提示 `command not found`

**解决：**
1. 检查 PATH 是否包含 `~/.cursortoolsets/bin`
   ```bash
   echo $PATH
   ```
2. 重新加载 shell 配置
   ```bash
   source ~/.bashrc  # 或 ~/.zshrc
   ```

### Q2: 权限被拒绝

**问题：** 执行提示 `Permission denied`

**解决：**
```bash
# 手动添加执行权限
chmod +x ~/.cursortoolsets/repos/my-toolset/bin/mytool
```

### Q3: 符号链接无效

**问题：** 链接指向的文件不存在

**解决：**
1. 检查 `bin` 配置的路径是否正确
2. 确保文件在打包时包含在 tarball 中
3. 重新安装包：
   ```bash
   cursortoolset uninstall my-toolset
   cursortoolset install my-toolset --no-cache
   ```

### Q4: Windows 上符号链接失败

**问题：** Windows 提示权限不足

**解决：**
1. 使用管理员权限运行 PowerShell
2. 或者启用开发者模式（Windows 10+）：
   - 设置 → 更新与安全 → 开发者选项 → 开发者模式

## 示例：完整的包配置

```json
{
  "name": "awesome-devtools",
  "displayName": "Awesome Dev Tools",
  "version": "2.1.0",
  "description": "一套实用的开发工具",
  "author": "Your Name",
  "license": "MIT",
  "keywords": ["dev", "tools", "cli"],
  
  "repository": {
    "type": "git",
    "url": "https://github.com/user/awesome-devtools.git"
  },
  
  "bin": {
    "devtools": "bin/devtools",
    "dt": "bin/devtools",
    "devtools-format": "scripts/format.sh",
    "devtools-lint": "scripts/lint.sh",
    "devtools-build": "scripts/build.sh"
  },
  
  "dist": {
    "tarball": "https://github.com/user/awesome-devtools/releases/download/v2.1.0/awesome-devtools-2.1.0.tar.gz",
    "sha256": "a1b2c3d4e5f6..."
  },
  
  "cursortoolset": {
    "minVersion": "1.0.0"
  }
}
```

## 相关文档

- [包开发指南](PACKAGE_DEV.md)
- [使用示例](USAGE_EXAMPLE.md)
- [安装指南](INSTALL_GUIDE.md)
