package config

import "github.com/XDwanj/go-gsgm/logger"

var (
	lutrisConfigPath = home + "/.config/lutris"
	lutrisCachePath  = home + "/.cache/lutris"
	lutrisLocalPath  = home + "/.local/share/lutris"
)

var (
	PgaDbPath = lutrisLocalPath + "/pga.db"
	RunScriptPath = lutrisConfigPath + "/games"
	CoverartPath  = lutrisCachePath + "/coverart"
	BannerPath    = lutrisCachePath + "/banners"
	IconPath      = home + "/.local/share/icons/hicolor/128x128/apps"
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

func init() {
	logger.Info("pgaDb path ", PgaDbPath)
}
