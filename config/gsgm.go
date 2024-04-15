package config

import (
	"os"
	"path/filepath"

	"github.com/XDwanj/go-gsgm/logger"
	"github.com/mitchellh/go-homedir"
)

var home, _ = homedir.Dir()

var (
	gsgmLocalPath  = filepath.Join(home, ".local", "share", "gsgm")
	gsgmConfigPath = filepath.Join(home, ".config", "gsgm")
	gsgmCachePath  = filepath.Join(home, ".cache", "gsgm")
)

var (
	GsgmTmpPath           = filepath.Join(gsgmCachePath, "tmp")
	GsgmDbPath            = filepath.Join(gsgmLocalPath, "gsgm.db")
	GsgmPrefixPath        = filepath.Join(gsgmLocalPath, "prefix")
	GsgmDefaultPrefixPath = filepath.Join(gsgmLocalPath, "prefix", "0")
)

const (
	GsgmInfoName     = "info.json"
	GsgmSettingName  = "setting.json"
	GsgmHistoryName  = "history.json"
	GsgmDirName      = ".gsgm"
	GsgmCoverName    = "cover"
	GsgmIsPackPrefix = "@_"
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
