package paths

import (
	"os"
	"path/filepath"
)

const (
	// EnvRootDir 环境变量名，用于指定 CursorToolset 安装根目录
	EnvRootDir = "CURSOR_TOOLSET_ROOT"
)

// GetRootDir 获取 CursorToolset 安装根目录
// 优先级：
// 1. 环境变量 CURSOR_TOOLSET_ROOT（如果设置）
// 2. 默认路径：~/.cursor/toolsets/CursorToolset
func GetRootDir() (string, error) {
	// 优先使用环境变量
	if rootDir := os.Getenv(EnvRootDir); rootDir != "" {
		return filepath.Abs(rootDir)
	}

	// 使用默认路径
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".cursor", "toolsets", "CursorToolset"), nil
}

// GetToolsetsDir 获取工具集安装目录
// 优先级：
// 1. 环境变量 CURSOR_TOOLSET_ROOT（如果设置，则使用 CURSOR_TOOLSET_ROOT/.cursor/toolsets）
// 2. 工作目录下的 .cursor/toolsets（如果 workDir 不为空）
// 3. 默认路径：~/.cursor/toolsets
func GetToolsetsDir(workDir string) (string, error) {
	// 如果设置了环境变量，使用环境变量下的 .cursor/toolsets
	if rootDir := os.Getenv(EnvRootDir); rootDir != "" {
		absRootDir, err := filepath.Abs(rootDir)
		if err != nil {
			return "", err
		}
		// 如果环境变量指向的是项目根目录，使用 .cursor/toolsets
		return filepath.Join(absRootDir, ".cursor", "toolsets"), nil
	}

	// 如果提供了工作目录，使用工作目录下的 .cursor/toolsets
	if workDir != "" {
		return filepath.Join(workDir, ".cursor", "toolsets"), nil
	}

	// 使用默认路径
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".cursor", "toolsets"), nil
}

// GetBinDir 获取可执行文件目录
// 返回 CursorToolset 的 bin 目录路径
func GetBinDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "bin"), nil
}

