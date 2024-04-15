package cmd

import (
	"github.com/XDwanj/go-gsgm/contro"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:    "scan",
	Short:  "扫描游戏当前库",
	Args:   cobra.MinimumNArgs(1),
	PreRun: InitLog,
	Run: func(cmd *cobra.Command, args []string) {
		contro.ScanByLibraries(args)
	},
}

func init() {
	defer rootCmd.AddCommand(scanCmd)
}
