// Package cmd /*
package cmd

import (
	"github.com/XDwanj/go-gsgm/contro"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:    "install",
	Short:  "安装游戏",
	PreRun: InitLog,
	Run: func(cmd *cobra.Command, args []string) {
		if installIsLibrary {
			contro.InstallByLibraries(args, installForce, installSafeMode)
			return
		}
		contro.InstallBySingles(args, installForce, installSafeMode)
	},
}

var (
	installForce     bool
	installIsLibrary bool
	installSafeMode  bool
)

func init() {
	defer rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&installForce, "force", "f", false, "覆盖安装")
	installCmd.Flags().BoolVarP(&installIsLibrary, "lib", "l", true, "是否是 Gsgm 游戏库")
	installCmd.Flags().BoolVarP(&installSafeMode, "safe", "s", false, "安装，但不添加图片")
}
