// Package generator 提供规则和 MCP 配置的生成器
// 生成器只负责生成内容，不负责写入文件
package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/shichao402/Dec/pkg/ide"
	"github.com/shichao402/Dec/pkg/rules"
	"github.com/shichao402/Dec/pkg/types"
)

// RulesGenerator 规则生成器（只生成内容，不写文件）
type RulesGenerator struct{}

// NewRulesGenerator 创建规则生成器
func NewRulesGenerator() *RulesGenerator {
	return &RulesGenerator{}
}

// RulePackInfo 规则包信息
type RulePackInfo struct {
	Name        string                 // 包名
	InstallPath string                 // 安装路径
	LocalPath   string                 // 本地开发路径（如果是链接的包）
	Pack        *types.Pack            // 包元数据
	UserConfig  map[string]interface{} // 用户配置
}

// GenerateAll 生成所有规则文件内容
// 返回规则文件列表，不写入磁盘
func (g *RulesGenerator) GenerateAll(packs []RulePackInfo, enabledBuiltinPacks map[string]bool) ([]ide.RuleFile, error) {
	var allRules []ide.RuleFile

	// 生成核心规则
	coreRules, err := g.generateCoreRules()
	if err != nil {
		return nil, fmt.Errorf("生成核心规则失败: %w", err)
	}
	allRules = append(allRules, coreRules...)

	// 生成内置包规则
	builtinRules, err := g.generateBuiltinPackRules(enabledBuiltinPacks)
	if err != nil {
		return nil, fmt.Errorf("生成内置包规则失败: %w", err)
	}
	allRules = append(allRules, builtinRules...)

	// 生成外部包规则
	for _, pack := range packs {
		packRules, err := g.generatePackRules(pack)
		if err != nil {
			return nil, fmt.Errorf("生成规则失败 (%s): %w", pack.Name, err)
		}
		allRules = append(allRules, packRules...)
	}

	return allRules, nil
}

// generateCoreRules 生成核心规则
func (g *RulesGenerator) generateCoreRules() ([]ide.RuleFile, error) {
	entries, err := fs.ReadDir(rules.EmbeddedRules, "resources/core")
	if err != nil {
		return nil, fmt.Errorf("读取核心规则目录失败: %w", err)
	}

	var result []ide.RuleFile
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".mdc") {
			continue
		}

		content, err := fs.ReadFile(rules.EmbeddedRules, "resources/core/"+entry.Name())
		if err != nil {
			return nil, fmt.Errorf("读取核心规则文件失败 (%s): %w", entry.Name(), err)
		}

		result = append(result, ide.RuleFile{
			Name:    fmt.Sprintf("dec-core-%s", entry.Name()),
			Content: string(content),
		})
	}

	return result, nil
}

// generateBuiltinPackRules 生成内置包规则
func (g *RulesGenerator) generateBuiltinPackRules(enabledPacks map[string]bool) ([]ide.RuleFile, error) {
	if enabledPacks == nil {
		return nil, nil
	}

	entries, err := fs.ReadDir(rules.EmbeddedRules, "resources/packs")
	if err != nil {
		return nil, fmt.Errorf("读取包规则目录失败: %w", err)
	}

	var result []ide.RuleFile
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".mdc") {
			continue
		}

		// 从文件名提取包名
		packName := strings.TrimSuffix(entry.Name(), ".mdc")
		if !enabledPacks[packName] {
			continue
		}

		content, err := fs.ReadFile(rules.EmbeddedRules, "resources/packs/"+entry.Name())
		if err != nil {
			return nil, fmt.Errorf("读取包规则文件失败 (%s): %w", entry.Name(), err)
		}

		result = append(result, ide.RuleFile{
			Name:    fmt.Sprintf("dec-pack-%s", entry.Name()),
			Content: string(content),
		})
	}

	return result, nil
}

// generatePackRules 生成单个包的规则
func (g *RulesGenerator) generatePackRules(pack RulePackInfo) ([]ide.RuleFile, error) {
	if pack.Pack == nil {
		return nil, nil
	}

	var result []ide.RuleFile

	// 确定包的实际路径
	packPath := pack.InstallPath
	if pack.LocalPath != "" {
		packPath = pack.LocalPath
	}

	// 处理规则包的规则文件
	for _, ruleFile := range pack.Pack.Rules {
		srcPath := filepath.Join(packPath, ruleFile)
		content, err := os.ReadFile(srcPath)
		if err != nil {
			return nil, fmt.Errorf("读取规则文件失败 (%s): %w", ruleFile, err)
		}

		// 渲染模板
		rendered := renderTemplate(string(content), pack.UserConfig)

		baseName := filepath.Base(ruleFile)
		result = append(result, ide.RuleFile{
			Name:    fmt.Sprintf("dec-%s-%s", pack.Name, baseName),
			Content: rendered,
		})
	}

	// 处理 MCP 包附带的规则
	for _, attachedRule := range pack.Pack.AttachedRules {
		srcPath := filepath.Join(packPath, attachedRule.File)
		content, err := os.ReadFile(srcPath)
		if err != nil {
			return nil, fmt.Errorf("读取附带规则失败 (%s): %w", attachedRule.File, err)
		}

		rendered := renderTemplate(string(content), pack.UserConfig)

		baseName := filepath.Base(attachedRule.File)
		result = append(result, ide.RuleFile{
			Name:    fmt.Sprintf("dec-%s-%s", pack.Name, baseName),
			Content: rendered,
		})
	}

	return result, nil
}

// renderTemplate 渲染模板，替换配置变量
func renderTemplate(content string, config map[string]interface{}) string {
	if config == nil {
		return content
	}

	result := content
	for key, value := range config {
		placeholder := fmt.Sprintf("{{config.%s}}", key)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", value))
	}

	return result
}

// GetAvailableBuiltinPacks 获取可用的内置规则包列表
func GetAvailableBuiltinPacks() ([]string, error) {
	entries, err := fs.ReadDir(rules.EmbeddedRules, "resources/packs")
	if err != nil {
		return nil, fmt.Errorf("读取包规则目录失败: %w", err)
	}

	var packs []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".mdc") {
			continue
		}
		packName := strings.TrimSuffix(entry.Name(), ".mdc")
		packs = append(packs, packName)
	}

	return packs, nil
}

// GetManagedRulePrefix 返回托管规则的前缀
func GetManagedRulePrefix() string {
	return "dec-"
}
