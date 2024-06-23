package cmd

import (
	"os"
	"path/filepath"

	"github.com/XDwanj/go-gsgm/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gsgm",
	Short:   "Gsgm Linux 游戏管理工具",
	PreRun:  InitLog,
	Version: "v0.1.2",
}

func InitLog(cmd *cobra.Command, args []string) {
	if Verbose {
		logger.Level = logger.Debug
	} else {
		logger.Level = logger.Error
	}
}

func ToAbsolutePath(paths []string) {
	for i := range paths {
		path, err := filepath.Abs(paths[i])
		if err != nil {
			panic("无法转换路径为绝对路径!!!")
		}
		logger.Info("args[", i, "] => ", path)
		paths[i] = path
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	Verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
