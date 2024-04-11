package config

import (
	"os"

	"github.com/XDwanj/go-gsgm/logger"
	"github.com/mitchellh/go-homedir"
)

var home, _ = homedir.Dir()

var (
	gsgmLocalPath  = home + "/.local/share/gsgm"
	gsgmConfigPath = home + "/.gsgmConfig/gsgm"
	gsgmCachePath  = home + "/.cache/gsgm"
)

var (
	GsgmTmpPath           = gsgmCachePath + "/tmp"
	GsgmDbPath            = gsgmLocalPath + "/gsgm.db"
	GsgmPrefixPath        = gsgmLocalPath + "/prefix"
	GsgmDefaultPrefixPath = gsgmLocalPath + "/prefix/0"
)

const (
	GsgmInfoName     = "info.json"
	GsgmSettingName  = "setting.json"
	GsgmHistoryName  = "history.json"
	GsgmDirName      = ".gsgm"
	GsgmCoverName    = "cover"
	GsgmIsPackName   = ".is-group"
	DefaultGroupName = "$default"
)

func init() {
	mkDirPaths := []string{
		gsgmLocalPath,
		gsgmConfigPath,
		gsgmCachePath,
		GsgmPrefixPath,
	}

	for _, path := range mkDirPaths {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Erro(err)
			panic(err)
		}
	}
}
