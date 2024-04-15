package mapper

import (
	"database/sql"
	"path/filepath"
	"strconv"
	"time"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_setting"
	"github.com/XDwanj/go-gsgm/lutris_dao"
)

func GsgmToLutrisLutrisGameDao(path string, info *gsgm_setting.GsgmInfo, setting *gsgm_setting.GsgmSetting, history *gsgm_setting.GsgmHistory) *lutris_dao.LutrisGame {

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

	// history
	var (
		lastplayed sql.NullInt64   = sql.NullInt64{Int64: 0, Valid: false}
		playtime   sql.NullFloat64 = sql.NullFloat64{Float64: 0, Valid: false}
	)
	if history != nil {
		lastplayed = sql.NullInt64{Int64: history.LastPlayedTime, Valid: true}
		playtime = sql.NullFloat64{Float64: time.Duration(history.PlayedDuration).Hours(), Valid: true}
	}

	lutrisGame := &lutris_dao.LutrisGame{
		Name:                 sql.NullString{String: name, Valid: true},
		Slug:                 sql.NullString{String: slug, Valid: true},
		Platform:             sql.NullString{String: string(setting.Platform), Valid: true},
		Runner:               sql.NullString{String: string(runner), Valid: true},
		Directory:            sql.NullString{String: path, Valid: true},
		Lastplayed:           lastplayed,
		Playtime:             playtime,
		Installed:            sql.NullInt32{Int32: 1, Valid: true},
		InstalledAt:          sql.NullInt64{Int64: installDate, Valid: true},
		Updated:              nil,
		Configpath:           sql.NullString{String: slug, Valid: true},
		HasCustomBanner:      sql.NullInt32{Int32: 1, Valid: true},
		HasCustomIcon:        sql.NullInt32{Int32: 1, Valid: true},
		HasCustomCoverartBig: sql.NullInt32{Int32: 1, Valid: true},
		Hidden:               sql.NullInt32{Int32: 0, Valid: true},
	}

	return lutrisGame
}

func GsgmToLutrisLutrisCategoryDao(packName string) *lutris_dao.LutrisCategory {
	return &lutris_dao.LutrisCategory{
		Name: packName,
	}
}
