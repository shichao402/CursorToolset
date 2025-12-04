# 环境变量配置说明

## CURSOR_TOOLSET_ROOT

### 概述

`CURSOR_TOOLSET_ROOT` 环境变量用于指定 CursorToolset 的安装根目录。如果未设置此环境变量，将使用默认路径。

### 默认路径

- **Linux/macOS**: `~/.cursor/toolsets/CursorToolset`
- **Windows**: `%USERPROFILE%\.cursor\toolsets\CursorToolset`

### 使用方法

#### 1. 临时设置（当前会话有效）

**Linux/macOS:**
```bash
export CURSOR_TOOLSET_ROOT=/path/to/your/root
cursortoolset install
```

**Windows PowerShell:**
```powershell
$env:CURSOR_TOOLSET_ROOT = "C:\path\to\your\root"
cursortoolset install
```

**Windows CMD:**
```cmd
set CURSOR_TOOLSET_ROOT=C:\path\to\your\root
cursortoolset install
```

#### 2. 永久设置

**Linux/macOS (添加到 ~/.bashrc 或 ~/.zshrc):**
```bash
export CURSOR_TOOLSET_ROOT=/path/to/your/root
```

**Windows (系统环境变量):**
1. 打开"系统属性" > "高级" > "环境变量"
2. 添加用户变量 `CURSOR_TOOLSET_ROOT`，值为目标路径

### 开发环境

在开发 CursorToolset 项目时，Makefile 会自动设置环境变量指向项目的 `.root` 目录：

```bash
# 运行任何 make 命令时，会自动设置 CURSOR_TOOLSET_ROOT=.root
make build
make test
make install
```

这确保了开发过程中的所有操作都使用项目本地的 `.root` 目录，不会影响系统安装。

### 路径优先级

1. **环境变量 `CURSOR_TOOLSET_ROOT`**（如果设置）
2. **工作目录下的 `.cursor/toolsets`**（如果指定了工作目录）
3. **默认路径** `~/.cursor/toolsets` 或 `%USERPROFILE%\.cursor\toolsets`

### 示例

#### 示例 1: 使用自定义安装路径

```bash
# 设置环境变量
export CURSOR_TOOLSET_ROOT=$HOME/my-tools/cursortoolset

# 安装工具集（会安装到 $HOME/my-tools/cursortoolset/.cursor/toolsets）
cursortoolset install
```

#### 示例 2: 开发时使用项目本地目录

```bash
# 在项目根目录运行
cd /path/to/CursorToolset

# Makefile 会自动设置 CURSOR_TOOLSET_ROOT=.root
make build
make test

# 工具集会安装到 .root/.cursor/toolsets
./cursortoolset install
```

#### 示例 3: 在特定项目中使用

```bash
# 在项目根目录
cd /path/to/my-project

# 设置环境变量指向项目目录
export CURSOR_TOOLSET_ROOT=$(pwd)

# 安装工具集（会安装到当前项目的 .cursor/toolsets）
cursortoolset install
```

### 注意事项

1. **环境变量优先级最高**: 如果设置了 `CURSOR_TOOLSET_ROOT`，将优先使用该路径
2. **路径必须存在**: 确保指定的路径存在，或者有创建权限
3. **开发环境隔离**: 开发时使用 `.root` 目录可以避免影响系统安装
4. **相对路径支持**: 环境变量可以使用相对路径，但建议使用绝对路径

### 相关文件

- `pkg/paths/paths.go`: 路径处理逻辑
- `Makefile`: 开发环境变量设置
- `.gitignore`: `.root/` 目录已添加到忽略列表

