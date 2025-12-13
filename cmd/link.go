package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/shichao402/Dec/pkg/registry"
	"github.com/shichao402/Dec/pkg/types"
	"github.com/spf13/cobra"
)

var (
	linkPath string
	linkList bool
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "链接本地开发包",
	Long: `将本地开发的包链接到 Dec，用于开发调试。

链接后的包优先级最高，dec sync 时会使用本地版本。

示例:
  dec link                    # 链接当前目录的包
  dec link --path /path/to/pack  # 链接指定路径的包
  dec link --list             # 列出所有已链接的包`,
	RunE: runLink,
}

func init() {
	RootCmd.AddCommand(linkCmd)
	linkCmd.Flags().StringVar(&linkPath, "path", "", "包路径（默认当前目录）")
	linkCmd.Flags().BoolVar(&linkList, "list", false, "列出所有已链接的包")
}

func runLink(cmd *cobra.Command, args []string) error {
	// 如果是列出模式
	if linkList {
		return listLinkedPacks()
	}

	// 确定包路径
	packPath := linkPath
	if packPath == "" {
		var err error
		packPath, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("获取当前目录失败: %w", err)
		}
	}

	// 转换为绝对路径
	absPath, err := filepath.Abs(packPath)
	if err != nil {
		return fmt.Errorf("解析路径失败: %w", err)
	}

	// 读取 package.json
	pack, err := loadPackageJSON(absPath)
	if err != nil {
		return err
	}

	// 验证包信息
	if pack.Name == "" {
		return fmt.Errorf("package.json 缺少 name 字段")
	}
	if pack.Type == "" {
		return fmt.Errorf("package.json 缺少 type 字段（应为 rule 或 mcp）")
	}

	// 链接包
	mgr := registry.NewMultiRegistryManager()
	_ = mgr.Load() // 忽略加载错误，可能是首次使用

	if err := mgr.LinkPack(pack.Name, absPath, pack.Version, pack.Type); err != nil {
		return fmt.Errorf("链接失败: %w", err)
	}

	fmt.Printf("✅ 已链接包: %s (%s)\n", pack.Name, pack.Type)
	fmt.Printf("   路径: %s\n", absPath)
	fmt.Printf("   版本: %s\n", pack.Version)
	fmt.Println()
	fmt.Println("💡 运行 dec sync 使链接生效")

	return nil
}

// loadPackageJSON 加载 package.json
func loadPackageJSON(packPath string) (*types.Pack, error) {
	packageJSONPath := filepath.Join(packPath, "package.json")
	data, err := os.ReadFile(packageJSONPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("未找到 package.json，请确保在包目录中运行")
		}
		return nil, fmt.Errorf("读取 package.json 失败: %w", err)
	}

	var pack types.Pack
	if err := json.Unmarshal(data, &pack); err != nil {
		return nil, fmt.Errorf("解析 package.json 失败: %w", err)
	}

	return &pack, nil
}

// listLinkedPacks 列出所有已链接的包
func listLinkedPacks() error {
	mgr := registry.NewMultiRegistryManager()
	_ = mgr.Load() // 忽略加载错误

	packs := mgr.ListLinkedPacks()
	if len(packs) == 0 {
		fmt.Println("没有已链接的包")
		fmt.Println()
		fmt.Println("使用 dec link 链接本地开发包")
		return nil
	}

	fmt.Println("已链接的包:")
	fmt.Println()
	for _, pack := range packs {
		fmt.Printf("  📦 %s (%s)\n", pack.Name, pack.Type)
		fmt.Printf("     路径: %s\n", pack.LocalPath)
		fmt.Printf("     版本: %s\n", pack.Version)
		fmt.Printf("     链接时间: %s\n", pack.LinkedAt)
		fmt.Println()
	}

	return nil
}
