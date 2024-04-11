// Package cmd /*
package cmd

import (
	"github.com/XDwanj/go-gsgm/contro"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:    "init",
	Short:  "初始化单个游戏或者一整个 Gsgm 库",
	PreRun: InitLog,
	Run: func(cmd *cobra.Command, args []string) {
		if initIsLibrary {
			contro.InitByLibraries(args)
			return
		}
		contro.InitBySingles(args)
	},
}

var (
	initIsLibrary bool
)

func init() {
	defer rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&initIsLibrary, "is-library", "l", false, "是否是 Gsgm 游戏库")
}
