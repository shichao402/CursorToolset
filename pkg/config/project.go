package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shichao402/Dec/pkg/paths"
	"github.com/shichao402/Dec/pkg/types"
	"gopkg.in/yaml.v3"
)

// ========================================
// 项目配置管理（.dec/config/）
// ========================================

// ProjectConfigManager 项目配置管理器
type ProjectConfigManager struct {
	projectRoot string
}

// NewProjectConfigManager 创建项目配置管理器
func NewProjectConfigManager(projectRoot string) *ProjectConfigManager {
	return &ProjectConfigManager{projectRoot: projectRoot}
}

// GetConfigDir 获取项目配置目录
func (m *ProjectConfigManager) GetConfigDir() string {
	return paths.GetProjectConfigDir(m.projectRoot)
}

// Exists 检查项目配置是否存在
func (m *ProjectConfigManager) Exists() bool {
	configPath := paths.GetProjectConfigPath(m.projectRoot)
	_, err := os.Stat(configPath)
	return err == nil
}

// ========================================
// 项目配置（project.yaml）
// ========================================

// LoadProjectConfig 加载项目配置
func (m *ProjectConfigManager) LoadProjectConfig() (*types.ProjectConfig, error) {
	configPath := paths.GetProjectConfigPath(m.projectRoot)
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("项目未初始化，请先运行 dec init")
		}
		return nil, fmt.Errorf("读取项目配置失败: %w", err)
	}

	var config types.ProjectConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析项目配置失败: %w", err)
	}

	return &config, nil
}

// SaveProjectConfig 保存项目配置
func (m *ProjectConfigManager) SaveProjectConfig(config *types.ProjectConfig) error {
	configPath := paths.GetProjectConfigPath(m.projectRoot)

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %w", err)
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化项目配置失败: %w", err)
	}

	return os.WriteFile(configPath, data, 0644)
}

// ========================================
// 技术栈配置（technology.yaml）
// ========================================

// LoadTechnologyConfig 加载技术栈配置
func (m *ProjectConfigManager) LoadTechnologyConfig() (*types.TechnologyConfig, error) {
	configPath := paths.GetTechnologyConfigPath(m.projectRoot)
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 返回空配置
			return &types.TechnologyConfig{}, nil
		}
		return nil, fmt.Errorf("读取技术栈配置失败: %w", err)
	}

	var config types.TechnologyConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析技术栈配置失败: %w", err)
	}

	return &config, nil
}

// SaveTechnologyConfig 保存技术栈配置
func (m *ProjectConfigManager) SaveTechnologyConfig(config *types.TechnologyConfig) error {
	configPath := paths.GetTechnologyConfigPath(m.projectRoot)

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %w", err)
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化技术栈配置失败: %w", err)
	}

	return os.WriteFile(configPath, data, 0644)
}

// ========================================
// 包配置（packs.yaml）
// ========================================

// LoadPacksConfig 加载包配置
func (m *ProjectConfigManager) LoadPacksConfig() (types.PacksConfig, error) {
	configPath := paths.GetPacksConfigPath(m.projectRoot)
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 返回空配置
			return make(types.PacksConfig), nil
		}
		return nil, fmt.Errorf("读取包配置失败: %w", err)
	}

	var config types.PacksConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析包配置失败: %w", err)
	}

	return config, nil
}

// SavePacksConfig 保存包配置
func (m *ProjectConfigManager) SavePacksConfig(config types.PacksConfig) error {
	configPath := paths.GetPacksConfigPath(m.projectRoot)

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %w", err)
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化包配置失败: %w", err)
	}

	return os.WriteFile(configPath, data, 0644)
}

// ========================================
// 便捷方法
// ========================================

// GetEnabledPacks 获取所有启用的包
func (m *ProjectConfigManager) GetEnabledPacks() ([]string, error) {
	config, err := m.LoadPacksConfig()
	if err != nil {
		return nil, err
	}

	var enabled []string
	for name, entry := range config {
		// 跳过注释字段
		if len(name) > 0 && name[0] == '_' {
			continue
		}
		if entry.Enabled {
			enabled = append(enabled, name)
		}
	}

	return enabled, nil
}

// EnablePack 启用包
// 注意：包类型由包自身的 package.json 定义，用户只需指定是否启用
func (m *ProjectConfigManager) EnablePack(name string, config map[string]interface{}) error {
	packs, err := m.LoadPacksConfig()
	if err != nil {
		return err
	}

	packs[name] = types.PackEntry{
		Enabled: true,
		Config:  config,
	}

	return m.SavePacksConfig(packs)
}

// DisablePack 禁用包
func (m *ProjectConfigManager) DisablePack(name string) error {
	packs, err := m.LoadPacksConfig()
	if err != nil {
		return err
	}

	if entry, exists := packs[name]; exists {
		entry.Enabled = false
		packs[name] = entry
		return m.SavePacksConfig(packs)
	}

	return nil
}

// createTechnologyTemplate 创建带注释的技术栈配置模板
func (m *ProjectConfigManager) createTechnologyTemplate() error {
	configPath := paths.GetTechnologyConfigPath(m.projectRoot)

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %w", err)
	}

	// 带注释的 YAML 模板
	template := `# 技术栈配置
# 根据项目实际情况修改，注释/取消注释即可开关

# 编程语言
# 可选: go, dart, typescript, javascript, python, java, kotlin, swift, rust, c, cpp
languages: []

# 框架
# 可选: flutter, react, vue, angular, django, fastapi, gin, echo, spring, nextjs
frameworks: []

# 目标平台
# 可选: cli, web, android, ios, macos, windows, linux, embedded
platforms: []

# 设计模式（项目的主要设计模式）
# 可选值及适用场景:
#   command   - CLI 工具、命令分发、操作系统
#   pipeline  - 数据处理流、编译器、ETL
#   mvc       - Web 应用、传统后端
#   mvvm      - 前端框架、数据绑定、GUI
#   repository - 数据访问层、领域驱动设计
#   factory   - 对象创建、依赖注入
#   strategy  - 算法可切换、策略模式
#   observer  - 事件驱动、发布订阅
#   plugin    - 插件系统、可扩展架构
#   decorator - 功能增强、中间件
#   adapter   - 接口适配、兼容层
patterns: []
`

	return os.WriteFile(configPath, []byte(template), 0644)
}

// ========================================
// 初始化
// ========================================

// InitProject 初始化项目配置
func (m *ProjectConfigManager) InitProject(name string, ides []string) error {
	// 创建项目配置
	projectConfig := &types.ProjectConfig{
		Name: name,
		IDEs: ides,
	}
	if err := m.SaveProjectConfig(projectConfig); err != nil {
		return err
	}

	// 创建带注释的技术栈配置模板
	if err := m.createTechnologyTemplate(); err != nil {
		return err
	}

	// 创建默认的包配置（dec 是 MCP 包，类型由其 package.json 定义）
	packsConfig := types.PacksConfig{
		"dec": {
			Enabled: true,
		},
	}
	if err := m.SavePacksConfig(packsConfig); err != nil {
		return err
	}

	return nil
}
