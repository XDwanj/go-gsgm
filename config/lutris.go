package config

import (
	"os"
	"path/filepath"

	"github.com/XDwanj/go-gsgm/logger"
)

var (
	lutrisConfigPath = filepath.Join(home, ".config", "lutris")
	lutrisCachePath  = filepath.Join(home, ".cache", "lutris")
	lutrisLocalPath  = filepath.Join(home, ".local", "share", "lutris")
)

var (
	PgaDbPath     = filepath.Join(lutrisLocalPath, "pga.db")
	RunScriptPath = filepath.Join(lutrisConfigPath, "games")
	CoverartPath  = filepath.Join(lutrisCachePath, "coverart")
	BannerPath    = filepath.Join(lutrisCachePath, "banners")
	IconPath      = filepath.Join(home, ".local", "share", "icons", "hicolor", "128x128", "apps")
)

const (
	SlugPrefix     = "gsgm-"
	ScriptPrefix   = "gsgm-"
	BannerPrefix   = "gsgm-"
	CoverartPrefix = "gsgm-"
	IconPrefix     = "lutris_gsgm-"
)

const (
	CoverartSuffix = "jpg"
	BannerSuffix   = "jpg"
	IconSuffix     = "png"
	ScriptSuffix   = "yml"
)

var (
	CoverartStd = imgStand{
		Width:  264,
		Height: 352,
	}
	BannerStd = imgStand{
		Width:  184,
		Height: 69,
	}
	IconStd = imgStand{
		Width:  128,
		Height: 128,
	}
)

type imgStand struct {
	Width, Height int
}

// Lutris 默认配置了一些环境变量，用于执行脚本
var LutrisEnvName = lutrisEnv{
	GameName:      "GAME_NAME",
	GameDirectory: "GAME_DIRECTORY",
}

type lutrisEnv struct {
	GameName      string
	GameDirectory string
}

func init() {
	logger.Info("pgaDb path ", PgaDbPath)
	mkDirsPaths := []string{
		filepath.Base(PgaDbPath),
		RunScriptPath,
		CoverartPath,
		BannerPath,
		IconPath,
	}

	for _, path := range mkDirsPaths {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Erro(err)
			panic(err)
		}
	}
}
