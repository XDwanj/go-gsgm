package cmd

import (
	"github.com/XDwanj/go-gsgm/contro"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:    "clean",
	Short:  "清理 Lutris 中 Gsgm 游戏",
	Args:   cobra.NoArgs,
	PreRun: InitLog,
	Run: func(cmd *cobra.Command, args []string) {
		ToAbsolutePath(args)
		contro.CleanAction()
	},
}

func init() {
	defer rootCmd.AddCommand(cleanCmd)
}
