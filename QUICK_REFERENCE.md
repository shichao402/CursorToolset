# Bin 功能快速参考

## 📋 一分钟速览

### 包开发者

**在 `toolset.json` 中添加：**
```json
{
  "bin": {
    "命令名": "相对路径"
  }
}
```

**示例：**
```json
{
  "name": "my-tool",
  "version": "1.0.0",
  "bin": {
    "mytool": "bin/mytool",
    "mytool-helper": "scripts/helper.sh"
  }
}
```

### 用户

**安装：**
```bash
cursortoolset install my-tool
```

**配置 PATH（一次性）：**
```bash
# Bash/Zsh
echo 'export PATH="$HOME/.cursortoolsets/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Fish
echo 'set -x PATH $HOME/.cursortoolsets/bin $PATH' >> ~/.config/fish/config.fish
```

**使用：**
```bash
mytool --help
mytool-helper process
```

## 🔗 关键链接

| 文档 | 说明 |
|------|------|
| [BIN_FEATURE.md](docs/BIN_FEATURE.md) | 完整功能文档 |
| [examples/README.md](docs/examples/README.md) | 配置示例 |
| [DEMO.md](DEMO.md) | 使用演示 |

## 🎯 常见配置模式

### 单命令
```json
{
  "bin": {
    "mycmd": "bin/mycmd"
  }
}
```

### 主命令 + 别名
```json
{
  "bin": {
    "mytool": "bin/mytool",
    "mt": "bin/mytool"
  }
}
```

### 工具集
```json
{
  "bin": {
    "tool": "bin/main",
    "tool-init": "scripts/init.sh",
    "tool-build": "scripts/build.sh",
    "tool-test": "scripts/test.sh"
  }
}
```

## ⚡ 快速命令

```bash
# 构建
make build

# 测试
make test

# 查看帮助
cursortoolset --help

# 安装包
cursortoolset install <package>

# 卸载包
cursortoolset uninstall <package>

# 列出所有包
cursortoolset list

# 查看包详情
cursortoolset info <package>
```

## 📁 目录结构

```
~/.cursortoolsets/
├── bin/              # 可执行程序符号链接
├── repos/            # 已安装的包
├── cache/            # 下载缓存
└── config/           # 配置文件
```

## ✅ 检查清单

**包开发者：**
- [ ] 创建可执行文件
- [ ] 设置执行权限 `chmod +x`
- [ ] 配置 `bin` 字段
- [ ] 测试打包
- [ ] 上传 Release
- [ ] 更新 registry

**用户：**
- [ ] 安装包 `cursortoolset install`
- [ ] 配置 PATH（一次性）
- [ ] 测试命令

## 🐛 故障排除

| 问题 | 解决方案 |
|------|---------|
| 命令找不到 | 检查 PATH 配置 |
| 权限被拒绝 | `chmod +x` 源文件 |
| 链接无效 | 重新安装包 |
| 冲突 | 使用包名前缀 |

## 💡 最佳实践

✅ 使用包名作为命令前缀  
✅ 提供 `--help` 选项  
✅ 使用 shebang (`#!/usr/bin/env bash`)  
✅ 避免常见系统命令名  
✅ 提供清晰的错误信息  

## 📊 实现状态

- ✅ 核心功能完成
- ✅ 跨平台支持
- ✅ 文档齐全
- ✅ 测试通过
- ✅ 向后兼容

## 🚀 版本信息

- **功能版本**: v1.2.0
- **实现日期**: 2025-12-05
- **状态**: ✅ 完成

---

**快速帮助**: 查看 [docs/BIN_FEATURE.md](docs/BIN_FEATURE.md) 获取完整文档
