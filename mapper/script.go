package mapper

import (
	"path/filepath"
	"strconv"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_setting"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_script"
	"github.com/duke-git/lancet/v2/formatter"
)

func GsgmToLutrisRunScript(path string, info *gsgm_setting.GsgmInfo, setting *gsgm_setting.GsgmSetting) *lutris_script.LutrisRunScript {
	slug := config.SlugPrefix + strconv.FormatInt(info.Id, 10)
	exe := filepath.Join(path, setting.Execute)

	var locale string
	if string(setting.LocaleCharSet) == "" {
		locale = string(gsgm_setting.ChinaUTF8)
	} else {
		locale = string(setting.LocaleCharSet)
	}

	var prefix string
	if setting.PrefixAlone {
		prefix = filepath.Join(config.GsgmPrefixPath, strconv.FormatInt(info.Id, 10))
	} else {
		prefix = config.GsgmDefaultPrefixPath
	}

	runScript := &lutris_script.LutrisRunScript{
		Slug:     slug,
		GameSlug: slug,
		Game: &lutris_script.GameDetail{
			Exe:        exe,
			Prefix:     prefix,
			WorkingDir: path,
		},
		System: &lutris_script.SystemDetail{
			Locale: locale,
			Env: map[string]string{
				"HOST_LC_ALL": locale,
			},
		},
		Wine: nil,
	}
	runScriptJson, err := formatter.Pretty(runScript)
	if err != nil {
		panic(err)
	}
	logger.Info("run script: ", runScriptJson)

	return runScript
}
