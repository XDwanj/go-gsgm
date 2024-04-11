package gsgm_service

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_dto"
	"github.com/XDwanj/go-gsgm/gsgm_setting"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/util"
	"github.com/duke-git/lancet/v2/fileutil"
)

// 简单递归遍历游戏库，返回所有游戏路径
func DeepGamePath(path string) []string {
	tmps := make([]string, 0)
	paths, err := util.Ls(path)
	if err != nil {
		return make([]string, 0)
	}
	paths = slices.DeleteFunc(paths, func(s string) bool {
		if !fileutil.IsDir(s) {
			return false
		}
		return isIgnoreBranch(s)
	})
	for _, v := range paths {
		walkGameFolder(v, func(s string) { tmps = append(tmps, s) })
	}

	return tmps
}

// 递归游戏库，并返回按游戏组分类的游戏集合
func DeepGamePack(path string) []gsgm_dto.GamePack {
	paths, err := util.Ls(path)
	if err != nil {
		return make([]gsgm_dto.GamePack, 0)
	}

	packPaths := make([]string, 0)
	paths = slices.DeleteFunc(paths, func(s string) bool {
		return !fileutil.IsDir(s) || isIgnoreBranch(s)
	})
	for _, path := range paths {
		walkGamePackFolder(path, func(s string) {
			packPaths = append(packPaths, s)
		})
	}

	var wg sync.WaitGroup
	wg.Add(len(packPaths)) // add tasks
	packs := make([]gsgm_dto.GamePack, len(packPaths))
	for i := range packPaths {
		packName := filepath.Base(packPaths[i])
		packs[i] = gsgm_dto.GamePack{PackName: packName}
		go func(idx int) {
			packs[idx].Paths = DeepGamePath(packPaths[idx])
			wg.Done() // done one
		}(i)
	}
	wg.Wait() // waiting everyone task

	packPathSet := make(map[string]bool)
	for _, packPath := range packPaths {
		packPathSet[packPath] = true
	}

	defaultGamePackPaths := make([]string, 0)
	paths = slices.DeleteFunc(paths, func(s string) bool {
		return packPathSet[s]
	})

	for _, path := range paths {
		walkGameFolder(path, func(gamePath string) {
			if !packPathSet[gamePath] {
				defaultGamePackPaths = append(defaultGamePackPaths, gamePath)
			}
		})
	}
	defaultGamePack := gsgm_dto.GamePack{
		PackName: config.DefaultGroupName,
		Paths:    defaultGamePackPaths,
	}

	packs = append(packs, defaultGamePack)

	return packs
}

// 检查该路径的游戏 info.json 合法
func CheckInfo(path string) error {
	infoPath := filepath.Join(path, config.GsgmDirName, config.GsgmInfoName)
	infoJson, err := fileutil.ReadFileToString(infoPath)
	if err != nil {
		logger.Warn(err)
		return errors.New("info.json is not found")
	}
	info := &gsgm_setting.GsgmInfo{}
	if err := json.Unmarshal([]byte(infoJson), info); err != nil {
		logger.Warn(err, path)
		return errors.New("info.json wrong format: " + err.Error())
	}

	return nil
}

// 检查该路径的游戏 setting.json 合法
func CheckSetting(path string) error {
	settingPath := filepath.Join(path, config.GsgmDirName, config.GsgmSettingName)
	settingJson, err := fileutil.ReadFileToString(settingPath)
	if err != nil {
		logger.Warn(err)
		return errors.New("setting.json is not found")
	}
	info := &gsgm_setting.GsgmSetting{}
	if err := json.Unmarshal([]byte(settingJson), info); err != nil {
		logger.Warn(err, path)
		return errors.New("setting.json wrong format: " + err.Error())
	}

	return nil
}

// 检查该路径的游戏 cover.[png,jpg,jpeg] 合法
func CheckImg(path string) error {
	if _, err := GetImgPath(path); err != nil {
		logger.Warn(err)
		return err
	}

	return nil
}

// deprecated: 检查目录和文件是否存在，重复判断了
func CheckDir(path string) error {
	gsgmPath := filepath.Join(path, config.GsgmDirName)
	checkPaths := []string{
		gsgmPath,
		filepath.Join(gsgmPath, config.GsgmSettingName),
		filepath.Join(gsgmPath, config.GsgmInfoName),
	}
	msgs := make([]string, 0)
	for _, checkPath := range checkPaths {
		if fileutil.IsExist(checkPath) {
			continue
		}
		msgs = append(msgs, "文件或目录不存在: "+checkPath)
	}
	if len(msgs) != 0 {
		return errors.New("[" + strings.Join(msgs, ", ") + "]")
	}
	return nil
}

// 尝试从 path 获取真实存在的 cover.[jpg,jpeg,png]
func GetImgPath(gamePath string) (string, error) {
	suffixes := []string{".jpg", ".jpeg", ".png"}
	for _, suffix := range suffixes {
		imgPath := filepath.Join(gamePath, config.GsgmDirName, "cover"+suffix)
		if fileutil.IsExist(imgPath) {
			logger.Info("img found ", imgPath)
			return imgPath, nil
		}
		logger.Tran("img is not found ", imgPath)
	}
	return "", errors.New("图片不存在: " + gamePath)
}

func walkGamePackFolder(folder string, block func(gamePackPath string)) {
	if !fileutil.IsDir(folder) {
		return
	}
	if isIgnoreBranch(folder) {
		return
	}
	if !isGameBranch(folder) {
		return
	}
	isPackPath := filepath.Join(folder, config.GsgmIsPackName)
	if fileutil.IsExist(isPackPath) {
		block(folder)
	} else {
		dirs, err := os.ReadDir(folder)
		if err != nil {
			return
		}
		for _, dir := range dirs {
			if !dir.IsDir() {
				continue
			}
			path := filepath.Join(folder, dir.Name())
			walkGamePackFolder(path, block)
		}
	}
}

// 是否是忽略分支
func isIgnoreBranch(path string) bool {
	name := filepath.Base(path)
	return strings.HasPrefix(name, "@@")
}

// 是否是分支
func isGameBranch(path string) bool {
	// 分支得是文件夹
	name := filepath.Base(path)
	if strings.HasPrefix(name, "@@") { // 分支不能是忽略文件夹
		return false
	}
	// 分支得是 "@" 开头
	return strings.HasPrefix(name, "@")
}

func walkGameFolder(folder string, block func(gamePath string)) {
	if !fileutil.IsDir(folder) {
		return
	}
	if isIgnoreBranch(folder) {
		return
	}
	if !isGameBranch(folder) {
		block(folder)
		return
	}

	paths, err := util.Ls(folder)
	if err != nil {
		return
	}
	paths = slices.DeleteFunc(paths, func(s string) bool {
		return !fileutil.IsDir(s)
	})
	for _, v := range paths {
		walkGameFolder(v, block)
	}
}
