package installer

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/firoyang/CursorToolset/pkg/types"
)

func TestNewInstaller(t *testing.T) {
	toolsetsDir := "/tmp/toolsets"
	workDir := "/tmp/work"

	installer := NewInstaller(toolsetsDir, workDir)

	if installer.ToolsetsDir != toolsetsDir {
		t.Errorf("Expected ToolsetsDir '%s', got '%s'", toolsetsDir, installer.ToolsetsDir)
	}

	if installer.WorkDir != workDir {
		t.Errorf("Expected WorkDir '%s', got '%s'", workDir, installer.WorkDir)
	}
}

func TestCopyFile(t *testing.T) {
	tmpDir := t.TempDir()
	installer := NewInstaller(tmpDir, tmpDir)

	// 创建源文件
	sourceFile := filepath.Join(tmpDir, "source.txt")
	content := []byte("test content")
	if err := os.WriteFile(sourceFile, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// 测试拷贝普通文件
	targetFile := filepath.Join(tmpDir, "target.txt")
	if err := installer.copyFile(sourceFile, targetFile, false); err != nil {
		t.Errorf("copyFile failed: %v", err)
	}

	// 验证文件内容
	targetContent, err := os.ReadFile(targetFile)
	if err != nil {
		t.Fatalf("Failed to read target file: %v", err)
	}

	if string(targetContent) != string(content) {
		t.Errorf("Expected content '%s', got '%s'", content, targetContent)
	}

	// 验证文件权限（普通文件）
	info, err := os.Stat(targetFile)
	if err != nil {
		t.Fatalf("Failed to stat target file: %v", err)
	}
	// Windows 上的文件权限显示不同，跳过权限检查
	if runtime.GOOS != "windows" {
		if info.Mode().Perm() != 0644 {
			t.Errorf("Expected file mode 0644, got %o", info.Mode().Perm())
		}
	}
}

func TestCopyFileExecutable(t *testing.T) {
	tmpDir := t.TempDir()
	installer := NewInstaller(tmpDir, tmpDir)

	// 创建源文件
	sourceFile := filepath.Join(tmpDir, "source.sh")
	content := []byte("#!/bin/bash\necho 'test'")
	if err := os.WriteFile(sourceFile, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// 测试拷贝可执行文件
	targetFile := filepath.Join(tmpDir, "target.sh")
	if err := installer.copyFile(sourceFile, targetFile, true); err != nil {
		t.Errorf("copyFile failed: %v", err)
	}

	// 验证文件权限（可执行文件）
	info, err := os.Stat(targetFile)
	if err != nil {
		t.Fatalf("Failed to stat target file: %v", err)
	}
	// Windows 上的文件权限显示不同，跳过权限检查
	if runtime.GOOS != "windows" {
		if info.Mode().Perm() != 0755 {
			t.Errorf("Expected file mode 0755, got %o", info.Mode().Perm())
		}
	}
}

func TestGetPlatformSuffix(t *testing.T) {
	installer := NewInstaller("", "")
	suffix := installer.getPlatformSuffix()

	// 验证格式：应该是 "os-arch" 的形式
	if suffix == "" {
		t.Error("Platform suffix should not be empty")
	}

	// 应该包含连字符
	if len(suffix) < 3 || suffix[0] == '-' || suffix[len(suffix)-1] == '-' {
		t.Errorf("Platform suffix has invalid format: %s", suffix)
	}
}

func TestIsPlatformSpecificFile(t *testing.T) {
	installer := NewInstaller("", "")

	tests := []struct {
		filename string
		expected bool
	}{
		{"tool-darwin-amd64", true},
		{"tool-darwin-arm64", true},
		{"tool-linux-amd64", true},
		{"tool-linux-arm64", true},
		{"tool-windows-amd64", true},
		{"tool", false},
		{"tool.exe", false},
		{"tool-v1.0.0", false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := installer.isPlatformSpecificFile(tt.filename)
			if result != tt.expected {
				t.Errorf("isPlatformSpecificFile(%s) = %v, expected %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestGetBaseExecutableName(t *testing.T) {
	installer := NewInstaller("", "")

	tests := []struct {
		filename string
		expected string
	}{
		{"tool-darwin-amd64", "tool"},
		{"tool-linux-arm64", "tool"},
		{"tool-windows-amd64.exe", "tool-windows-amd64"},
		{"gh-action-debug-darwin-arm64", "gh-action-debug"},
		{"simple-tool", "simple-tool"},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := installer.getBaseExecutableName(tt.filename)
			if result != tt.expected {
				t.Errorf("getBaseExecutableName(%s) = %s, expected %s", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestCopyTargetWithNonExistentSource(t *testing.T) {
	tmpDir := t.TempDir()
	installer := NewInstaller(tmpDir, tmpDir)

	target := types.InstallTarget{
		Source:    "nonexistent/",
		Files:     []string{"*.txt"},
		Overwrite: false,
	}

	// 应该不返回错误，只是警告
	err := installer.copyTarget("target/", target, tmpDir)
	if err != nil {
		t.Errorf("Expected no error for non-existent source, got: %v", err)
	}
}

func TestCopyFilesByPattern(t *testing.T) {
	tmpDir := t.TempDir()
	installer := NewInstaller(tmpDir, tmpDir)

	sourceDir := filepath.Join(tmpDir, "source")
	targetDir := filepath.Join(tmpDir, "target")

	// 创建源目录和文件
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		t.Fatalf("Failed to create target dir: %v", err)
	}

	// 创建测试文件
	files := []string{"test1.txt", "test2.txt", "readme.md"}
	for _, f := range files {
		path := filepath.Join(sourceDir, f)
		if err := os.WriteFile(path, []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// 测试通配符匹配
	config := types.InstallTarget{
		Overwrite: false,
	}

	matched, err := installer.copyFilesByPattern(sourceDir, targetDir, "*.txt", config)
	if err != nil {
		t.Errorf("copyFilesByPattern failed: %v", err)
	}

	if !matched {
		t.Error("Expected to match files, but got no match")
	}

	// 验证文件已拷贝
	for _, f := range []string{"test1.txt", "test2.txt"} {
		targetPath := filepath.Join(targetDir, f)
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist", targetPath)
		}
	}

	// readme.md 不应该被拷贝
	if _, err := os.Stat(filepath.Join(targetDir, "readme.md")); err == nil {
		t.Error("Expected readme.md NOT to be copied")
	}
}

func TestCopyFilesByPatternSingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	installer := NewInstaller(tmpDir, tmpDir)

	sourceDir := filepath.Join(tmpDir, "source")
	targetDir := filepath.Join(tmpDir, "target")

	// 创建源目录和文件
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		t.Fatalf("Failed to create target dir: %v", err)
	}

	sourceFile := filepath.Join(sourceDir, "config.yaml")
	if err := os.WriteFile(sourceFile, []byte("config content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 测试单个文件匹配
	config := types.InstallTarget{
		Overwrite: false,
	}

	matched, err := installer.copyFilesByPattern(sourceDir, targetDir, "config.yaml", config)
	if err != nil {
		t.Errorf("copyFilesByPattern failed: %v", err)
	}

	if !matched {
		t.Error("Expected to match file, but got no match")
	}

	// 验证文件已拷贝
	targetPath := filepath.Join(targetDir, "config.yaml")
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		t.Error("Expected config.yaml to exist")
	}
}

