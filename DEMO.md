# Bin 功能演示

## 快速演示

### 1. 包开发者视角

创建一个带有可执行程序的包：

```bash
# 创建包目录
mkdir my-awesome-tool
cd my-awesome-tool

# 创建可执行文件
mkdir -p bin scripts

cat > bin/awesome << 'EOF'
#!/usr/bin/env bash
echo "🚀 Awesome Tool v1.0.0"
echo "Usage: awesome [command]"
EOF
chmod +x bin/awesome

cat > scripts/init.sh << 'EOF'
#!/usr/bin/env bash
echo "🎉 Initializing project..."
mkdir -p .awesome
echo "Done!"
EOF
chmod +x scripts/init.sh

# 创建 toolset.json
cat > toolset.json << 'EOF'
{
  "name": "awesome-tool",
  "displayName": "Awesome Tool",
  "version": "1.0.0",
  "description": "一个超棒的开发工具",
  "author": "Your Name",
  "license": "MIT",
  "keywords": ["tool", "dev", "awesome"],
  
  "repository": {
    "type": "git",
    "url": "https://github.com/user/awesome-tool.git"
  },
  
  "bin": {
    "awesome": "bin/awesome",
    "awesome-init": "scripts/init.sh"
  },
  
  "dist": {
    "tarball": "https://github.com/user/awesome-tool/releases/download/v1.0.0/awesome-tool-1.0.0.tar.gz",
    "sha256": "待填写"
  }
}
EOF

# 打包
tar -czf awesome-tool-1.0.0.tar.gz *

# 计算 SHA256
shasum -a 256 awesome-tool-1.0.0.tar.gz
```

### 2. 用户视角

安装和使用包：

```bash
# 安装包
$ cursortoolset install awesome-tool

📦 安装 awesome-tool@1.0.0
  ⬇️  下载: https://github.com/.../awesome-tool-1.0.0.tar.gz
  ✅ SHA256 校验通过
  📂 解压到: ~/.cursortoolsets/repos/awesome-tool
✅ awesome-tool 安装完成

  🔗 创建可执行程序链接...
    ✅ awesome -> bin/awesome
    ✅ awesome-init -> scripts/init.sh

  💡 将 bin 目录添加到 PATH:
    export PATH="/Users/username/.cursortoolsets/bin:$PATH"

# 添加到 PATH
$ export PATH="$HOME/.cursortoolsets/bin:$PATH"

# 使用命令
$ awesome
🚀 Awesome Tool v1.0.0
Usage: awesome [command]

$ awesome-init
🎉 Initializing project...
Done!

# 查看链接
$ ls -la ~/.cursortoolsets/bin/
awesome -> ../repos/awesome-tool/bin/awesome
awesome-init -> ../repos/awesome-tool/scripts/init.sh

# 卸载
$ cursortoolset uninstall awesome-tool

🗑️  卸载 awesome-tool
  🔗 清理可执行程序链接...
    ✅ 已删除 awesome
    ✅ 已删除 awesome-init
✅ awesome-tool 卸载完成
```

## 实际应用示例

### 示例 1: DevOps 工具集

```json
{
  "name": "devops-toolkit",
  "version": "2.0.0",
  "bin": {
    "deploy": "bin/deploy",
    "rollback": "bin/rollback",
    "monitor": "scripts/monitor.py",
    "backup": "scripts/backup.sh"
  }
}
```

使用：
```bash
$ deploy production
$ monitor --service api
$ backup --full
$ rollback --to v1.9.0
```

### 示例 2: 项目模板生成器

```json
{
  "name": "project-generator",
  "version": "1.5.0",
  "bin": {
    "gen": "bin/generator",
    "gen-react": "templates/react/create.sh",
    "gen-vue": "templates/vue/create.sh",
    "gen-node": "templates/node/create.sh"
  }
}
```

使用：
```bash
$ gen --help
$ gen-react my-app
$ gen-vue my-project
$ gen-node api-server
```

### 示例 3: 代码质量工具

```json
{
  "name": "code-quality",
  "version": "1.0.0",
  "bin": {
    "cq": "bin/cq",
    "cq-lint": "scripts/lint.sh",
    "cq-format": "scripts/format.sh",
    "cq-test": "scripts/test.sh",
    "cq-coverage": "scripts/coverage.sh"
  }
}
```

使用：
```bash
$ cq check
$ cq-lint --fix
$ cq-format
$ cq-test --watch
$ cq-coverage --html
```

## 对比：之前 vs 现在

### 之前（无 bin 功能）

**包开发者：**
```markdown
# 安装说明
1. 克隆仓库
2. 添加到 PATH：
   export PATH="$PATH:~/.cursortoolsets/repos/my-tool/bin"
3. 或创建符号链接：
   ln -s ~/.cursortoolsets/repos/my-tool/bin/mytool ~/bin/mytool
```

**用户操作：**
```bash
# 安装
cursortoolset install my-tool

# 手动配置（麻烦！）
export PATH="$PATH:$HOME/.cursortoolsets/repos/my-tool/bin"
# 或
ln -s ~/.cursortoolsets/repos/my-tool/bin/mytool ~/bin/mytool

# 使用
mytool --help
```

### 现在（有 bin 功能）

**包开发者：**
```json
{
  "bin": {
    "mytool": "bin/mytool"
  }
}
```

**用户操作：**
```bash
# 安装（自动处理）
cursortoolset install my-tool

# 一次性配置 PATH
export PATH="$HOME/.cursortoolsets/bin:$PATH"

# 直接使用
mytool --help
```

## 高级用法

### 命令别名

```json
{
  "bin": {
    "mytool": "bin/mytool",
    "mt": "bin/mytool"        // 简写别名
  }
}
```

### 多平台支持

```json
{
  "bin": {
    "mytool": "bin/mytool.sh",           // Unix
    "mytool-windows": "bin/mytool.exe"   // Windows
  }
}
```

### 复杂工具链

```json
{
  "bin": {
    "tool": "bin/main",
    "tool-dev": "scripts/dev.sh",
    "tool-prod": "scripts/prod.sh",
    "tool-init": "scripts/init.sh",
    "tool-config": "config/setup.py",
    "tool-doctor": "scripts/doctor.sh"
  }
}
```

## 完整工作流演示

```bash
# === 包开发者 ===

# 1. 初始化包
cursortoolset init my-cli-tool
cd my-cli-tool

# 2. 创建可执行文件
mkdir bin
cat > bin/mycli << 'EOF'
#!/usr/bin/env bash
echo "My CLI Tool"
EOF
chmod +x bin/mycli

# 3. 配置 toolset.json
cat > toolset.json << 'EOF'
{
  "name": "my-cli-tool",
  "version": "1.0.0",
  "bin": {
    "mycli": "bin/mycli"
  },
  "dist": {
    "tarball": "https://github.com/user/my-cli-tool/releases/download/v1.0.0/my-cli-tool-1.0.0.tar.gz",
    "sha256": "..."
  }
}
EOF

# 4. 打包发布
tar -czf my-cli-tool-1.0.0.tar.gz *
# 上传到 GitHub Release

# 5. 提交到 Registry
# 编辑 registry.json，提交 PR

# === 用户 ===

# 1. 更新索引
cursortoolset registry update

# 2. 搜索包
cursortoolset search cli

# 3. 查看详情
cursortoolset info my-cli-tool

# 4. 安装
cursortoolset install my-cli-tool

# 5. 配置 PATH（一次性）
echo 'export PATH="$HOME/.cursortoolsets/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# 6. 使用
mycli --help
mycli command --option value

# 7. 更新
cursortoolset update my-cli-tool

# 8. 卸载
cursortoolset uninstall my-cli-tool
```

## 常见场景

### 场景 1: 快速原型工具

```bash
# 创建简单的脚本工具
echo '#!/usr/bin/env bash' > quick-script.sh
echo 'echo "Quick tool"' >> quick-script.sh
chmod +x quick-script.sh

# 配置 bin
{
  "bin": {
    "quick": "quick-script.sh"
  }
}

# 安装后立即使用
$ quick
Quick tool
```

### 场景 2: 团队共享工具

```bash
# 团队成员 A 创建工具
# 团队成员 B、C、D 只需：
$ cursortoolset install team-toolkit
$ export PATH="$HOME/.cursortoolsets/bin:$PATH"
$ team-tool --help
```

### 场景 3: 持续集成

```bash
# CI 脚本中
export PATH="$HOME/.cursortoolsets/bin:$PATH"
cursortoolset install ci-tools
ci-build --parallel
ci-test --coverage
ci-deploy --environment staging
```

## 最佳实践提示

### ✅ 推荐

1. **使用包名前缀**
   ```json
   {
     "bin": {
       "mytool": "bin/main",
       "mytool-init": "scripts/init.sh"
     }
   }
   ```

2. **提供帮助命令**
   ```bash
   mytool --help
   mytool version
   ```

3. **使用 shebang**
   ```bash
   #!/usr/bin/env bash
   #!/usr/bin/env python3
   ```

### ⚠️ 注意

1. **避免常见命令名**
   - ❌ `ls`, `cd`, `cat`
   - ✅ `mytool-ls`, `mt-cd`

2. **检查文件存在**
   - 确保 bin 配置的文件在包中

3. **设置执行权限**
   - `chmod +x bin/*`

---

**提示**: 查看完整文档 [docs/BIN_FEATURE.md](docs/BIN_FEATURE.md)
