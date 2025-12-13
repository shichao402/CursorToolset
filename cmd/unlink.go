package cmd

import (
	"fmt"

	"github.com/shichao402/Dec/pkg/registry"
	"github.com/spf13/cobra"
)

var unlinkAll bool

var unlinkCmd = &cobra.Command{
	Use:   "unlink [package-name]",
	Short: "移除本地链接",
	Long: `移除本地开发包的链接。

示例:
  dec unlink github-issue     # 移除指定包的链接
  dec unlink --all            # 移除所有链接`,
	RunE: runUnlink,
}

func init() {
	RootCmd.AddCommand(unlinkCmd)
	unlinkCmd.Flags().BoolVar(&unlinkAll, "all", false, "移除所有链接")
}

func runUnlink(cmd *cobra.Command, args []string) error {
	mgr := registry.NewMultiRegistryManager()
	_ = mgr.Load() // 忽略加载错误

	// 移除所有链接
	if unlinkAll {
		packs := mgr.ListLinkedPacks()
		if len(packs) == 0 {
			fmt.Println("没有已链接的包")
			return nil
		}

		if err := mgr.UnlinkAll(); err != nil {
			return fmt.Errorf("移除链接失败: %w", err)
		}

		fmt.Printf("✅ 已移除 %d 个链接\n", len(packs))
		return nil
	}

	// 移除指定包的链接
	if len(args) == 0 {
		return fmt.Errorf("请指定要移除链接的包名，或使用 --all 移除所有链接")
	}

	packName := args[0]
	if err := mgr.UnlinkPack(packName); err != nil {
		return err
	}

	fmt.Printf("✅ 已移除链接: %s\n", packName)
	fmt.Println()
	fmt.Println("💡 运行 dec sync 使更改生效")

	return nil
}
