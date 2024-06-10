/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

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
		fmt.Println(
			strings.Trim(`
#!/bin/sh
/usr/bin/gsgm sync -v "$GAME_DIRECTORY"
				`,
				"\n",
			))
	},
}

func init() {
	genCmd.AddCommand(genPostExitScript)
	rootCmd.AddCommand(genCmd)
}
