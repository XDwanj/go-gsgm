package cmd

import (
	"fmt"
	"strings"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:    "gen",
	Short:  "生成辅助性的脚本命令",
	PreRun: InitLog,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

var genPostExitScript = &cobra.Command{
	Use:   "lupes",
	Short: "generate lutris post exit script",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var sh strings.Builder
		sh.WriteString("#!/bin/sh")
		sh.WriteString("\n")
		sh.WriteString("/usr/bin/gsgm " + syncCmd.Name() + " -v \"$" + config.LutrisEnvName.GameDirectory + "\"")
		sh.WriteString("\n")
		fmt.Println(sh.String())
	},
}

func init() {
	genCmd.AddCommand(genPostExitScript)
	rootCmd.AddCommand(genCmd)
}
