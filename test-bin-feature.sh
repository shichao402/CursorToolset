#!/usr/bin/env bash
# 测试 bin 功能的简单脚本

set -e

echo "🧪 测试 bin 功能"
echo ""

# 创建临时测试目录
TEST_DIR=$(mktemp -d)
echo "📁 创建测试目录: $TEST_DIR"

# 设置测试环境变量
export CURSOR_TOOLSET_HOME="$TEST_DIR/.cursortoolsets"

# 创建测试包
TEST_PKG="$TEST_DIR/test-bin-pkg"
mkdir -p "$TEST_PKG/bin"
mkdir -p "$TEST_PKG/scripts"

echo "📝 创建测试包..."

# 创建可执行文件
cat > "$TEST_PKG/bin/testcmd" << 'EOF'
#!/usr/bin/env bash
echo "Hello from testcmd!"
EOF
chmod +x "$TEST_PKG/bin/testcmd"

cat > "$TEST_PKG/scripts/helper.sh" << 'EOF'
#!/usr/bin/env bash
echo "Helper script executed"
EOF
chmod +x "$TEST_PKG/scripts/helper.sh"

# 创建 toolset.json
cat > "$TEST_PKG/toolset.json" << 'EOF'
{
  "name": "test-bin-pkg",
  "version": "1.0.0",
  "description": "测试 bin 功能的包",
  "bin": {
    "testcmd": "bin/testcmd",
    "testhelper": "scripts/helper.sh"
  },
  "dist": {
    "tarball": "file:///dev/null",
    "sha256": "fake"
  }
}
EOF

echo "📦 打包测试包..."
cd "$TEST_PKG"
tar -czf "$TEST_DIR/test-bin-pkg-1.0.0.tar.gz" ./*

echo "🔐 计算 SHA256..."
SHA256=$(shasum -a 256 "$TEST_DIR/test-bin-pkg-1.0.0.tar.gz" | awk '{print $1}')
echo "   SHA256: $SHA256"

# 更新 toolset.json 中的 dist
cat > "$TEST_PKG/toolset.json" << EOF
{
  "name": "test-bin-pkg",
  "version": "1.0.0",
  "description": "测试 bin 功能的包",
  "bin": {
    "testcmd": "bin/testcmd",
    "testhelper": "scripts/helper.sh"
  },
  "dist": {
    "tarball": "file://$TEST_DIR/test-bin-pkg-1.0.0.tar.gz",
    "sha256": "$SHA256"
  }
}
EOF

echo ""
echo "✅ 测试包创建完成"
echo ""
echo "📍 测试包位置: $TEST_PKG"
echo "📍 CURSOR_TOOLSET_HOME: $CURSOR_TOOLSET_HOME"
echo ""
echo "🔍 检查目录结构:"
tree "$TEST_PKG" 2>/dev/null || find "$TEST_PKG" -type f

echo ""
echo "💡 手动测试步骤:"
echo ""
echo "1. 设置环境变量:"
echo "   export CURSOR_TOOLSET_HOME=\"$CURSOR_TOOLSET_HOME\""
echo ""
echo "2. 模拟安装（手动解压）:"
echo "   mkdir -p \"\$CURSOR_TOOLSET_HOME/repos/test-bin-pkg\""
echo "   tar -xzf \"$TEST_DIR/test-bin-pkg-1.0.0.tar.gz\" -C \"\$CURSOR_TOOLSET_HOME/repos/test-bin-pkg\""
echo ""
echo "3. 使用 cursortoolset 测试 linkBinaries 功能"
echo "   （需要添加相应的测试代码）"
echo ""
echo "4. 清理:"
echo "   rm -rf \"$TEST_DIR\""
echo ""

# 提示不自动清理，让用户可以手动测试
echo "⚠️  测试目录未清理，请手动清理:"
echo "   rm -rf \"$TEST_DIR\""
