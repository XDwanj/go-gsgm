package cmd

import (
	"github.com/XDwanj/go-gsgm/contro"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:    "check",
	Short:  "检查当前游戏或者游戏库目录是否合法",
	Args:   cobra.MinimumNArgs(1),
	PreRun: InitLog,
	Run: func(cmd *cobra.Command, args []string) {
		if checkIsLibrary {
			contro.CheckByLibraries(args)
			return
		}
		contro.CheckBySingles(args)
	},
}

var checkIsLibrary bool

func init() {
	defer rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolVarP(&checkIsLibrary, "lib", "l", false, "是否是游戏库")
}
