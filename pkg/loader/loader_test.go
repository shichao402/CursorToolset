package loader

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/firoyang/CursorToolset/pkg/types"
)

func TestLoadToolsets(t *testing.T) {
	// 创建临时测试文件
	tmpDir := t.TempDir()
	toolsetsPath := filepath.Join(tmpDir, "available-toolsets.json")

	testData := []*types.ToolsetInfo{
		{
			Name:        "test-toolset",
			DisplayName: "Test Toolset",
			GitHubURL:   "https://github.com/test/test.git",
			Description: "A test toolset",
			Version:     "1.0.0",
		},
	}

	data, err := json.MarshalIndent(testData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	if err := os.WriteFile(toolsetsPath, data, 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// 测试加载
	toolsets, err := LoadToolsets(toolsetsPath)
	if err != nil {
		t.Fatalf("LoadToolsets failed: %v", err)
	}

	if len(toolsets) != 1 {
		t.Errorf("Expected 1 toolset, got %d", len(toolsets))
	}

	if toolsets[0].Name != "test-toolset" {
		t.Errorf("Expected name 'test-toolset', got '%s'", toolsets[0].Name)
	}

	if toolsets[0].GitHubURL != "https://github.com/test/test.git" {
		t.Errorf("Expected GitHub URL 'https://github.com/test/test.git', got '%s'", toolsets[0].GitHubURL)
	}
}

func TestLoadToolsetsFileNotFound(t *testing.T) {
	_, err := LoadToolsets("/nonexistent/available-toolsets.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestLoadToolsetsInvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	toolsetsPath := filepath.Join(tmpDir, "available-toolsets.json")

	// 写入无效的 JSON
	if err := os.WriteFile(toolsetsPath, []byte("invalid json"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	_, err := LoadToolsets(toolsetsPath)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestFindToolset(t *testing.T) {
	toolsets := []*types.ToolsetInfo{
		{Name: "toolset-1", DisplayName: "Toolset 1"},
		{Name: "toolset-2", DisplayName: "Toolset 2"},
		{Name: "toolset-3", DisplayName: "Toolset 3"},
	}

	// 测试找到工具集
	found := FindToolset(toolsets, "toolset-2")
	if found == nil {
		t.Error("Expected to find toolset-2, got nil")
	} else if found.Name != "toolset-2" {
		t.Errorf("Expected name 'toolset-2', got '%s'", found.Name)
	}

	// 测试找不到工具集
	notFound := FindToolset(toolsets, "nonexistent")
	if notFound != nil {
		t.Error("Expected nil for non-existent toolset, got result")
	}
}

func TestGetToolsetsPath(t *testing.T) {
	tmpDir := t.TempDir()
	
	// 创建 available-toolsets.json
	toolsetsPath := filepath.Join(tmpDir, "available-toolsets.json")
	if err := os.WriteFile(toolsetsPath, []byte("[]"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 测试获取路径
	path := GetToolsetsPath(tmpDir)
	if path != toolsetsPath {
		t.Errorf("Expected path '%s', got '%s'", toolsetsPath, path)
	}

	// 验证文件存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Returned path does not exist")
	}
}

