package generator

import (
	"fmt"
	"path/filepath"

	"github.com/shichao402/Dec/pkg/types"
)

// MCPGenerator MCP 配置生成器（只生成内容，不写文件）
type MCPGenerator struct{}

// NewMCPGenerator 创建 MCP 配置生成器
func NewMCPGenerator() *MCPGenerator {
	return &MCPGenerator{}
}

// MCPPackInfo MCP 包信息
type MCPPackInfo struct {
	Name        string                 // 包名
	InstallPath string                 // 安装路径
	LocalPath   string                 // 本地开发路径（如果是链接的包）
	Pack        *types.Pack            // 包元数据
	UserConfig  map[string]interface{} // 用户配置
}

// GenerateAll 生成 MCP 配置
// 返回配置对象，不写入磁盘
func (g *MCPGenerator) GenerateAll(packs []MCPPackInfo) (*types.MCPConfig, []string, error) {
	config := &types.MCPConfig{
		MCPServers: make(map[string]types.MCPServer),
	}

	// 托管的 MCP 名称列表
	var managedNames []string

	// 添加 Dec 自身的 MCP Server
	config.MCPServers["dec"] = g.GenerateDecServer()
	managedNames = append(managedNames, "dec")

	// 添加其他 MCP 包
	for _, pack := range packs {
		if pack.Pack == nil || pack.Pack.MCP == nil {
			continue
		}

		server, err := g.buildMCPServer(pack)
		if err != nil {
			return nil, nil, fmt.Errorf("构建 MCP Server 配置失败 (%s): %w", pack.Name, err)
		}

		config.MCPServers[pack.Name] = *server
		managedNames = append(managedNames, pack.Name)
	}

	return config, managedNames, nil
}

// buildMCPServer 构建单个 MCP Server 配置
func (g *MCPGenerator) buildMCPServer(pack MCPPackInfo) (*types.MCPServer, error) {
	if pack.Pack == nil || pack.Pack.MCP == nil {
		return nil, fmt.Errorf("包 %s 缺少 MCP 配置", pack.Name)
	}

	mcpConfig := pack.Pack.MCP

	// 确定包的实际路径
	packPath := pack.InstallPath
	if pack.LocalPath != "" {
		packPath = pack.LocalPath
	}

	// 构建命令路径
	command := mcpConfig.Command
	if !filepath.IsAbs(command) {
		command = filepath.Join(packPath, command)
	}

	// 处理环境变量模板
	env := make(map[string]string)
	for key, value := range mcpConfig.Env {
		env[key] = value
	}

	// 应用用户配置
	if pack.UserConfig != nil {
		if tokenEnv, ok := pack.UserConfig["token_env"].(string); ok {
			for key := range env {
				if key == "GITHUB_TOKEN" || key == "TOKEN" {
					env[key] = fmt.Sprintf("${%s}", tokenEnv)
				}
			}
		}
	}

	return &types.MCPServer{
		Command: command,
		Args:    mcpConfig.Args,
		Env:     env,
	}, nil
}

// GenerateDecServer 生成 Dec 自身的 MCP Server 配置
func (g *MCPGenerator) GenerateDecServer() types.MCPServer {
	return types.MCPServer{
		Command: "dec",
		Args:    []string{"serve"},
	}
}

// MergeConfig 合并配置（保留用户手动添加的配置）
func (g *MCPGenerator) MergeConfig(existing *types.MCPConfig, generated *types.MCPConfig, managedPacks []string) *types.MCPConfig {
	result := &types.MCPConfig{
		MCPServers: make(map[string]types.MCPServer),
	}

	// 创建 managed packs 集合
	managed := make(map[string]bool)
	for _, name := range managedPacks {
		managed[name] = true
	}

	// 保留非托管的配置
	if existing != nil {
		for name, server := range existing.MCPServers {
			if !managed[name] {
				result.MCPServers[name] = server
			}
		}
	}

	// 添加生成的配置
	for name, server := range generated.MCPServers {
		result.MCPServers[name] = server
	}

	return result
}
