package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/XDwanj/go-gsgm/cmd"
	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/logger"
)

func main() {
	envPrintln()
	cmd.Execute()
}

// 判断环境，打印软件系统变量
func envPrintln() {

	debugMode, ok := os.LookupEnv("GSGM_DEBUG")
	if ok && debugMode == "1" {
		logger.Level = logger.Debug

		slices.Sort(config.LutrisEnvs)
		slices.Sort(config.GsgmEnvs)
		fmt.Println("lutris env")
		for _, v := range config.LutrisEnvs {
			fmt.Println("\t" + v)
		}
		fmt.Println("gsgm env")
		for _, v := range config.GsgmEnvs {
			fmt.Println("\t" + v)
		}
		fmt.Println()
	}
}
