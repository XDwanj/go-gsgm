package gsgm_service

import (
	"encoding/json"
	"path/filepath"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_setting"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/duke-git/lancet/v2/fileutil"
)

// 接受游戏路径，返回 GsgmInfo
func GetGsgmInfoByPath(path string) (*gsgm_setting.GsgmInfo, error) {
	infoPath := filepath.Join(path, ".gsgm", config.GsgmInfoName)
	logger.Info("info path ", infoPath)
	jsonStr, err := fileutil.ReadFileToString(infoPath)
	if err != nil {
		return nil, err
	}

	info := &gsgm_setting.GsgmInfo{}
	if err = json.Unmarshal([]byte(jsonStr), info); err != nil {
		return nil, err
	}

	return info, nil
}

// 接受游戏路径，返回 GsgmSetting
func GetGsgmSettingByPath(path string) (*gsgm_setting.GsgmSetting, error) {
	settingPath := filepath.Join(path, ".gsgm", config.GsgmSettingName)
	jsonStr, err := fileutil.ReadFileToString(settingPath)
	if err != nil {
		return nil, err
	}

	setting := &gsgm_setting.GsgmSetting{}
	if err = json.Unmarshal([]byte(jsonStr), setting); err != nil {
		return nil, err
	}

	return setting, nil
}

// 接受游戏路径，返回 GsgmHistory
func GetGsgmHistoryByPath(path string) (*gsgm_setting.GsgmHistory, error) {
	historyPath := filepath.Join(path, ".gsgm", config.GsgmHistoryName)
	jsonStr, err := fileutil.ReadFileToString(historyPath)
	if err != nil {
		return nil, err
	}

	history := &gsgm_setting.GsgmHistory{}
	if err = json.Unmarshal([]byte(jsonStr), history); err != nil {
		return nil, err
	}

	return history, nil
}
