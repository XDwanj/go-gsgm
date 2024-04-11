package mapper

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_setting"
	"github.com/XDwanj/go-gsgm/lutris_dao"
)

func GsgmToLutrisLutrisGameDao(path string, info *gsgm_setting.GsgmInfo, setting *gsgm_setting.GsgmSetting) *lutris_dao.LutrisGame {

	name := filepath.Base(path)
	slug := config.SlugPrefix + strconv.FormatInt(info.Id, 10)
	// exe := filepath.Join(path, setting.ExecuteLocation)
	installDate := info.InitTime
	if info.InitTime == 0 {
		installDate = time.Now().Unix()
	}

	runner := setting.Runner
	if len(string(runner)) == 0 {
		var err error
		runner, err = setting.Platform.DefaultRunner()
		if err != nil {
			runner = gsgm_setting.Wine
		}
	}

	// TODO: 没有处理游戏游玩记录
	lutrisGame := &lutris_dao.LutrisGame{
		Name:                 name,
		Slug:                 slug,
		Platform:             string(setting.Platform),
		Runner:               string(runner),
		Directory:            path,
		Installed:            1,
		InstalledAt:          installDate,
		Updated:              nil,
		Configpath:           slug,
		HasCustomBanner:      1,
		HasCustomIcon:        1,
		HasCustomCoverartBig: 1,
		Hidden:               0,
	}

	return lutrisGame
}

func GsgmToLutrisLutrisCategoryDao(packName string) *lutris_dao.LutrisCategory {
	return &lutris_dao.LutrisCategory{
		Name: packName,
	}
}
