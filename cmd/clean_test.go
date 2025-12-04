package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCleanDirectory(t *testing.T) {
	// 创建临时测试目录
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "test-clean")

	// 创建测试目录和文件
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	testFile := filepath.Join(testDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 测试清理
	if err := cleanDirectory(testDir); err != nil {
		t.Errorf("cleanDirectory failed: %v", err)
	}

	// 验证目录已被删除
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Error("Directory should be deleted but still exists")
	}
}

func TestCleanDirectoryNonExistent(t *testing.T) {
	// 测试清理不存在的目录（应该不报错）
	tmpDir := t.TempDir()
	nonExistentDir := filepath.Join(tmpDir, "nonexistent")

	if err := cleanDirectory(nonExistentDir); err != nil {
		t.Errorf("cleanDirectory should not fail for non-existent directory: %v", err)
	}
}

func TestCleanDirectoryWithSubdirectories(t *testing.T) {
	// 创建临时测试目录
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "test-clean")

	// 创建嵌套目录结构
	subDir := filepath.Join(testDir, "subdir", "nested")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}

	// 创建多个文件
	files := []string{
		filepath.Join(testDir, "file1.txt"),
		filepath.Join(subDir, "file2.txt"),
	}

	for _, f := range files {
		if err := os.WriteFile(f, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", f, err)
		}
	}

	// 测试清理
	if err := cleanDirectory(testDir); err != nil {
		t.Errorf("cleanDirectory failed: %v", err)
	}

	// 验证整个目录树已被删除
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Error("Directory tree should be completely deleted but still exists")
	}
}

