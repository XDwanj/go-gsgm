package cmd

import (
	"os"

	gsgmLogger "github.com/XDwanj/go-gsgm/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gsgm",
	Short:   "Gsgm Linux 游戏管理工具",
	Version: "v0.1.0",
}

func InitLog(cmd *cobra.Command, args []string) {
	if Verbose {
		gsgmLogger.Level = gsgmLogger.Debug
	} else {
		gsgmLogger.Level = gsgmLogger.NoLog
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
