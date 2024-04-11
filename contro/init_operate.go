package contro

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_dto"
	"github.com/XDwanj/go-gsgm/gsgm_service"
	"github.com/XDwanj/go-gsgm/gsgm_setting"
	"github.com/XDwanj/go-gsgm/util"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/formatter"
	"github.com/jedib0t/go-pretty/v6/text"
)

func InitByLibraries(libPaths []string) {
	packs := make([]gsgm_dto.GamePack, 0)
	for _, libPath := range libPaths {
		packs = slices.Concat(packs, gsgm_service.DeepGamePack(libPath))
	}

	for _, pack := range packs {
		packName := pack.PackName
		fmt.Println("===================================")
		fmt.Println("当前在初始化游戏库", text.FgGreen.Sprint(packName))
		for _, path := range pack.Paths {
			InitByOne(path)
		}
	}
}

func InitBySingles(paths []string) {
	for _, path := range paths {
		InitByOne(path)
	}
}

func InitByOne(path string) {
	fmt.Println("---------------------------")
	fmt.Println("当前游戏所在目录:", path)
	fmt.Println("游戏名:", filepath.Base(path))
	// info
	info := &gsgm_setting.GsgmInfo{
		Id:       util.NextSnowId(),
		InitTime: time.Now().Unix(),
	}

	files := deepTwo(path)
	platformId := 1
	fmt.Println("游戏平台[Windows=1, Linux=2]，默认为 1")
	fmt.Print(">> ")
	fmt.Scanln(&platformId)

	platform := gsgm_setting.Windows
	if platformId != 1 {
		platform = gsgm_setting.Linux
	}

	files = slices.DeleteFunc(files, func(path string) bool {
		return DeleteRule(path, platform)
	})
	for i, file := range files {
		fmt.Print(text.FgMagenta.Sprint(i) + " " + file[len(path)+1:] + "\n")
	}
	idx := 0
	fmt.Println("选择可执行文件，回车默认为 0")
	fmt.Print(">>")
	fmt.Scanln(&idx)

	exe := files[idx][len(path)+1:]

	prefixFlag := 0
	if platform == gsgm_setting.Windows {
		fmt.Println("WINE_PREFIX 隔离[是=1, 否=0]，回车默认为 0")
		fmt.Print(">>")
		fmt.Scanln(&prefixFlag)
	}
	prefixAlone := (prefixFlag == 1)

	// TODO: Runner 参数未做分类
	runner, _ := platform.DefaultRunner()

	// setting
	setting := &gsgm_setting.GsgmSetting{
		ExecuteLocation: exe,
		WinePrefixAlone: prefixAlone,
		LocaleCharSet:   gsgm_setting.ChinaUTF8,
		Platform:        platform,
		Runner:          runner,
	}

	infoJson, _ := formatter.Pretty(info)
	settingJson, _ := formatter.Pretty(setting)
	fmt.Println(infoJson)
	fmt.Println(settingJson)

	infoPath := filepath.Join(path, config.GsgmDirName, config.GsgmInfoName)
	settingPath := filepath.Join(path, config.GsgmDirName, config.GsgmSettingName)

	// write
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		fileutil.WriteStringToFile(infoPath, infoJson, false)
	}()
	go func() {
		defer waitGroup.Done()
		fileutil.WriteStringToFile(settingPath, settingJson, false)
	}()
	waitGroup.Wait()

	fmt.Println("该游戏的初始化已结束，请别忘了在 {game}/.gsgm/ 路径下放一个 cover.[jpg,jpeg,png] 文件")
}

func DeleteRule(path string, platform gsgm_setting.Platform) bool {
	windowsSuffixs := []string{".exe", ".bat"}
	linuxSuffix := []string{".appimage", ".sh"}

	fileName := filepath.Base(path)
	fileName = strings.ToLower(fileName)

	switch platform {
	case gsgm_setting.Windows:
		for _, suffix := range windowsSuffixs {
			if strings.HasSuffix(fileName, suffix) {
				return false
			}
		}
	case gsgm_setting.Linux:
		for _, suffix := range linuxSuffix {
			if strings.HasSuffix(fileName, suffix) {
				return false
			}
			if !strings.Contains(fileName, ".") {
				return false
			}
		}
	}
	return true
}

func deepTwo(path string) []string {
	if !fileutil.IsDir(path) {
		return make([]string, 0)
	}

	paths, err := util.Ls(path)
	if err != nil {
		return make([]string, 0)
	}

	files := make([]string, 0)
	dirs := make([]string, 0)
	for i := range paths {
		if fileutil.IsDir(paths[i]) {
			dirs = append(dirs, paths[i])
			continue
		}
		files = append(files, paths[i])
	}

	for _, dir := range dirs {
		mix, _ := util.Ls(dir)
		for i := range mix {
			if fileutil.IsDir(mix[i]) {
				continue
			}
			files = append(files, mix[i])
		}
	}

	return files
}
