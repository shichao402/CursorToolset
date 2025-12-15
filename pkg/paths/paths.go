package paths

import (
	"os"
	"path/filepath"
)

const (
	// EnvRootDir 环境变量名，用于指定 Dec 根目录
	// 类似于 pip 的 PYTHONUSERBASE，brew 的 HOMEBREW_PREFIX
	EnvRootDir = "DEC_HOME"
)

// 目录结构设计（重构后）：
// ~/.dec/                        <- DEC_HOME（默认）
// ├── mcp/                       <- MCP 工具包安装目录
// │   └── github-issue/
// │       ├── bin/
// │       ├── rules/
// │       └── package.json
// ├── rules/                     <- 规则包安装目录
// │   └── documentation/
// │       ├── rules/
// │       └── package.json
// ├── registry/                  <- 多注册表目录
// │   ├── local.json             <- 本地开发注册表
// │   ├── test.json              <- 测试注册表
// │   └── registry.json          <- 正式注册表
// ├── cache/                     <- 缓存目录
// │   └── packages/              <- 下载的 tarball 缓存
// ├── config/                    <- 全局配置
// │   ├── system.json            <- 系统配置
// │   └── settings.json          <- 用户设置
// └── bin/                       <- 可执行文件
//     └── dec
//
// 项目配置目录：
// <project>/.dec/config/
// ├── project.yaml              <- 项目信息
// ├── technology.yaml           <- 技术栈
// └── packs.yaml                <- 启用的包

// GetRootDir 获取 Dec 根目录
// 优先级：
// 1. 环境变量 DEC_HOME（如果设置）
// 2. 默认路径：~/.dec
func GetRootDir() (string, error) {
	if rootDir := os.Getenv(EnvRootDir); rootDir != "" {
		return filepath.Abs(rootDir)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".dec"), nil
}

// ========================================
// 全局目录
// ========================================

// GetMCPDir 获取 MCP 工具包安装目录
func GetMCPDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "mcp"), nil
}

// GetRulesDir 获取规则包安装目录
func GetRulesDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "rules"), nil
}

// GetRegistryDir 获取注册表目录
func GetRegistryDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "registry"), nil
}

// GetCacheDir 获取缓存根目录
func GetCacheDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "cache"), nil
}

// GetPackageCacheDir 获取下载包的缓存目录
func GetPackageCacheDir() (string, error) {
	cacheDir, err := GetCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cacheDir, "packages"), nil
}

// GetConfigDir 获取全局配置文件目录
func GetConfigDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "config"), nil
}

// GetBinDir 获取可执行文件目录
func GetBinDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "bin"), nil
}

// ========================================
// 注册表路径
// ========================================

// GetLocalRegistryPath 获取本地开发注册表路径
func GetLocalRegistryPath() (string, error) {
	registryDir, err := GetRegistryDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(registryDir, "local.json"), nil
}

// GetTestRegistryPath 获取测试注册表路径
func GetTestRegistryPath() (string, error) {
	registryDir, err := GetRegistryDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(registryDir, "test.json"), nil
}

// GetOfficialRegistryPath 获取正式注册表路径
func GetOfficialRegistryPath() (string, error) {
	registryDir, err := GetRegistryDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(registryDir, "registry.json"), nil
}

// ========================================
// 包安装路径
// ========================================

// GetMCPPackPath 获取 MCP 工具包的安装路径
func GetMCPPackPath(packName string) (string, error) {
	mcpDir, err := GetMCPDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(mcpDir, packName), nil
}

// GetRulePackPath 获取规则包的安装路径
func GetRulePackPath(packName string) (string, error) {
	rulesDir, err := GetRulesDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rulesDir, packName), nil
}

// GetPackageCachePath 获取下载包的缓存路径
func GetPackageCachePath(packageName, version string) (string, error) {
	cacheDir, err := GetPackageCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cacheDir, packageName+"-"+version+".tar.gz"), nil
}

// ========================================
// 项目配置路径
// ========================================

// GetProjectConfigDir 获取项目配置目录
func GetProjectConfigDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".dec", "config")
}

// GetProjectConfigPath 获取项目配置文件路径
func GetProjectConfigPath(projectRoot string) string {
	return filepath.Join(GetProjectConfigDir(projectRoot), "project.yaml")
}

// GetTechnologyConfigPath 获取技术栈配置文件路径
func GetTechnologyConfigPath(projectRoot string) string {
	return filepath.Join(GetProjectConfigDir(projectRoot), "technology.yaml")
}

// GetPacksConfigPath 获取包配置文件路径
func GetPacksConfigPath(projectRoot string) string {
	return filepath.Join(GetProjectConfigDir(projectRoot), "packs.yaml")
}

// ========================================
// IDE 输出路径
// ========================================

// IDE 目录名映射
var ideDirectories = map[string]string{
	"cursor":    ".cursor",
	"codebuddy": ".codebuddy",
	"windsurf":  ".windsurf",
	"trae":      ".trae",
}

// GetIDERulesDir 获取指定 IDE 的规则输出目录
func GetIDERulesDir(projectRoot, ide string) string {
	dir, ok := ideDirectories[ide]
	if !ok {
		dir = "." + ide // 默认使用 .{ide} 格式
	}
	return filepath.Join(projectRoot, dir, "rules")
}

// GetIDEMCPConfigPath 获取指定 IDE 的 MCP 配置文件路径
func GetIDEMCPConfigPath(projectRoot, ide string) string {
	dir, ok := ideDirectories[ide]
	if !ok {
		dir = "." + ide
	}
	return filepath.Join(projectRoot, dir, "mcp.json")
}

// GetCursorRulesDir 获取 Cursor 规则输出目录
func GetCursorRulesDir(projectRoot string) string {
	return GetIDERulesDir(projectRoot, "cursor")
}

// GetCursorMCPConfigPath 获取 Cursor MCP 配置文件路径
func GetCursorMCPConfigPath(projectRoot string) string {
	return GetIDEMCPConfigPath(projectRoot, "cursor")
}

// ========================================
// 工具函数
// ========================================

// EnsureDir 确保目录存在
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// EnsureAllDirs 确保所有必要的目录存在
func EnsureAllDirs() error {
	dirs := []func() (string, error){
		GetMCPDir,
		GetRulesDir,
		GetRegistryDir,
		GetPackageCacheDir,
		GetConfigDir,
		GetBinDir,
	}

	for _, getDirFunc := range dirs {
		dir, err := getDirFunc()
		if err != nil {
			return err
		}
		if err := EnsureDir(dir); err != nil {
			return err
		}
	}

	return nil
}

// ========================================
// 兼容旧版本（待移除）
// ========================================

// GetReposDir 获取已安装包的目录（兼容旧版本）
// Deprecated: 使用 GetMCPDir 或 GetRulesDir
func GetReposDir() (string, error) {
	rootDir, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, "repos"), nil
}

// GetManifestCacheDir 获取 manifest 缓存目录（兼容旧版本）
// Deprecated: 不再需要单独缓存 manifest
func GetManifestCacheDir() (string, error) {
	cacheDir, err := GetCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cacheDir, "manifests"), nil
}

// GetRegistryPath 获取本地 registry 缓存文件路径（兼容旧版本）
// Deprecated: 使用 GetOfficialRegistryPath
func GetRegistryPath() (string, error) {
	return GetOfficialRegistryPath()
}

// GetManifestPath 获取指定包的 manifest 缓存路径（兼容旧版本）
// Deprecated: 不再需要单独缓存 manifest
func GetManifestPath(packageName string) (string, error) {
	manifestDir, err := GetManifestCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(manifestDir, packageName+".json"), nil
}

// GetPackagePath 获取已安装包的路径（兼容旧版本）
// Deprecated: 使用 GetMCPPackPath 或 GetRulePackPath
func GetPackagePath(packageName string) (string, error) {
	reposDir, err := GetReposDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(reposDir, packageName), nil
}
