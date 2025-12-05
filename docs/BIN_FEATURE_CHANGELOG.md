# Bin 功能更新日志

## 版本：v1.2.0+

## 新增功能：可执行程序暴露 (bin)

### 概述

新增了包可执行程序暴露功能，允许包开发者在 `toolset.json` 中配置 `bin` 字段，管理器会在安装时自动创建符号链接到 `~/.cursortoolsets/bin/` 目录。

### 核心改动

#### 1. 类型定义 (`pkg/types/toolset.go`)

**新增字段：**
```go
type Manifest struct {
    // ... 其他字段
    
    // 可执行程序配置（可选）
    Bin map[string]string `json:"bin,omitempty"` // {"命令名": "相对路径"}
    
    // ... 其他字段
}
```

**字段说明：**
- `Bin`: 键值对映射，键为命令名，值为包内可执行文件的相对路径
- 可选字段，不配置则不影响现有包

#### 2. 安装器 (`pkg/installer/installer.go`)

**新增功能：**

1. **linkBinaries()** - 创建可执行程序符号链接
   - 读取 manifest 的 `bin` 配置
   - 验证源文件是否存在
   - 在 `~/.cursortoolsets/bin/` 创建符号链接
   - 设置可执行权限（Unix 系统）
   - 支持 Windows 平台（.exe 扩展名）

2. **unlinkBinaries()** - 清理可执行程序符号链接
   - 读取 manifest 的 `bin` 配置
   - 删除对应的符号链接

3. **loadPackageManifest()** - 从已安装包加载 manifest
   - 读取包目录中的 `toolset.json`
   - 用于卸载时获取 bin 配置

**修改的函数：**

- `Install()`: 添加了符号链接创建步骤
- `Uninstall()`: 添加了符号链接清理步骤

#### 3. 路径管理 (`pkg/paths/paths.go`)

**已有支持：**
- `GetBinDir()` - 获取 bin 目录路径
- `EnsureAllDirs()` - 已包含 bin 目录的创建

### 用户体验改进

#### 安装时输出

```
📦 安装 my-toolset@1.0.0
  ⬇️  下载: https://...
  ✅ SHA256 校验通过
  📂 解压到: ~/.cursortoolsets/repos/my-toolset
✅ my-toolset 安装完成

  🔗 创建可执行程序链接...
    ✅ mytool -> bin/mytool
    ✅ myhelper -> scripts/helper.sh

  💡 将 bin 目录添加到 PATH:
    export PATH="/Users/username/.cursortoolsets/bin:$PATH"
```

#### 卸载时输出

```
🗑️  卸载 my-toolset
  🔗 清理可执行程序链接...
    ✅ 已删除 mytool
    ✅ 已删除 myhelper
✅ my-toolset 卸载完成
```

### 兼容性

#### 向后兼容

- ✅ 不影响现有未配置 `bin` 字段的包
- ✅ `bin` 字段为可选字段（`omitempty`）
- ✅ 安装/卸载流程保持兼容

#### 跨平台支持

| 平台 | 符号链接 | 扩展名处理 | 权限设置 |
|------|---------|-----------|---------|
| Linux | ✅ | - | chmod 0755 |
| macOS | ✅ | - | chmod 0755 |
| Windows | ✅ | 自动添加 .exe | - |

### 错误处理

#### 宽松策略

安装时：
- 源文件不存在：⚠️ 警告并跳过，不阻断安装
- 创建链接失败：⚠️ 警告并继续，不阻断安装
- 设置权限失败：⚠️ 警告并继续

卸载时：
- 读取 manifest 失败：继续删除包目录
- 删除链接失败：⚠️ 警告并继续

**设计理由：**
确保即使 bin 配置有问题，也不影响包的基本安装和卸载功能。

### 安全考虑

#### 符号链接安全

- ✅ 只能链接到包自身目录内的文件
- ✅ 使用相对路径验证，防止路径遍历
- ✅ 符号链接目标由管理器控制

#### 路径隔离

- ✅ 所有内容在 `~/.cursortoolsets/` 下
- ✅ 不修改系统目录
- ✅ 用户需手动添加 bin 到 PATH

### 使用示例

#### 包开发者配置

```json
{
  "name": "awesome-cli",
  "version": "1.0.0",
  "bin": {
    "awesome": "bin/awesome",
    "aws": "bin/awesome",
    "awesome-config": "scripts/config.sh"
  },
  "dist": {
    "tarball": "https://github.com/user/awesome-cli/releases/download/v1.0.0/awesome-cli-1.0.0.tar.gz",
    "sha256": "abc123..."
  }
}
```

#### 用户使用

```bash
# 安装包
cursortoolset install awesome-cli

# 添加到 PATH（一次性）
export PATH="$HOME/.cursortoolsets/bin:$PATH"

# 使用命令
awesome --help
awesome-config init
```

### 测试建议

#### 单元测试

- [ ] 测试符号链接创建
- [ ] 测试符号链接删除
- [ ] 测试错误处理（文件不存在）
- [ ] 测试 Windows 平台扩展名处理

#### 集成测试

- [ ] 完整安装流程测试
- [ ] 完整卸载流程测试
- [ ] 多个可执行程序的包
- [ ] 不同文件类型（脚本、二进制）

#### 平台测试

- [ ] Linux 测试
- [ ] macOS 测试
- [ ] Windows 测试

### 文档更新

已更新的文档：

1. ✅ **BIN_FEATURE.md** - 完整的功能文档
2. ✅ **cursortoolset-development.md** - MCP 规则更新
3. ✅ **BIN_FEATURE_CHANGELOG.md** - 本文档

需要更新的文档：

- [ ] README.md - 添加 bin 功能说明
- [ ] PACKAGE_DEV.md - 添加 bin 配置指南
- [ ] USAGE_EXAMPLE.md - 添加使用示例

### 后续优化

#### 可能的增强功能

1. **bin 冲突检测**
   - 安装前检查命令名是否已被其他包使用
   - 提供冲突解决策略

2. **bin 版本管理**
   - 支持同一包的多版本共存
   - 命令别名（如 `mytool@1.0`）

3. **PATH 自动配置**
   - 检测 shell 类型
   - 自动写入配置文件（需用户确认）

4. **bin 列表命令**
   ```bash
   cursortoolset bin list
   cursortoolset bin which mytool
   ```

### 迁移指南

#### 对于包开发者

**旧版本（无 bin 支持）：**
```json
{
  "name": "my-tool",
  "version": "1.0.0"
}
```

用户需要手动创建链接或添加 PATH。

**新版本（使用 bin）：**
```json
{
  "name": "my-tool",
  "version": "2.0.0",
  "bin": {
    "mytool": "bin/mytool"
  }
}
```

管理器自动处理，用户体验更好。

#### 对于用户

- 无需任何操作
- 旧包照常使用
- 新包自动享受 bin 功能

### 贡献者

- 功能设计与实现
- 文档编写
- 测试用例

### 相关 Issue/PR

- Issue: #xxx - 支持可执行程序暴露
- PR: #xxx - 实现 bin 功能

---

**更新日期**: 2025-12-05
**相关版本**: v1.2.0+
