package cmd

import (
	"github.com/XDwanj/go-gsgm/contro"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:    "sync",
	Short:  "同步游戏时长",
	Long:   "将 Lutris 的游戏记录流转到 Gsgm 中，如需反向同步，使用 install 命令即可",
	Args:   cobra.MinimumNArgs(1),
	PreRun: InitLog,
	Run: func(cmd *cobra.Command, args []string) {
		ToAbsolutePath(args)
		if syncIsLibrary {
			contro.SyncByLibraries(args, syncForce)
			return
		}
		contro.SyncBySingles(args, syncForce)
	},
}

var (
	syncIsLibrary bool
	syncForce     bool
)

func init() {
	defer rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().BoolVarP(&syncIsLibrary, "lib", "l", false, "是否是 Gsgm 游戏库")
	syncCmd.Flags().BoolVarP(&syncForce, "force", "f", false, "是否覆盖同步")
}
