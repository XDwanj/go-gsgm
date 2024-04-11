package contro

import (
	"fmt"
	"slices"
	"sync"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_dto"
	"github.com/XDwanj/go-gsgm/gsgm_service"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_service"
	"github.com/XDwanj/go-gsgm/mapper"
	"github.com/jedib0t/go-pretty/v6/text"
)

func InstallByLibraries(libPaths []string, force bool, safeMode bool) {
	packs := make([]gsgm_dto.GamePack, 0)
	for _, libPath := range libPaths {
		packs = slices.Concat(packs, gsgm_service.DeepGamePack(libPath))
	}

	for _, pack := range packs {
		var waitGroup sync.WaitGroup
		waitGroup.Add(len(pack.Paths))

		packName := pack.PackName

		for _, path := range pack.Paths {
			path, force, packName := path, force, packName
			go func() {
				defer waitGroup.Done()
				InstallByOne(path, force, packName, safeMode)
			}()
		}
		waitGroup.Wait()
	}
}

func InstallBySingles(paths []string, force bool, safeMode bool) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(paths))
	for _, path := range paths {
		go func(path string, force bool) {
			defer waitGroup.Done()
			InstallByOne(path, force, config.DefaultGroupName, safeMode)
		}(path, force)
	}
	waitGroup.Wait()
}

func InstallByOne(path string, force bool, packName string, safeMode bool) {
	logger.Info("try install one: ", path)

	if !CheckOne(path) {
		fmt.Println(text.BgYellow.Sprint("add error:"), path)
		return
	}

	// before
	info, err := gsgm_service.GetGsgmInfoByPath(path)
	if err != nil {
		logger.Erro(err)
	}
	setting, err := gsgm_service.GetGsgmSettingByPath(path)
	if err != nil {
		logger.Erro(err)
	}
	runScript := mapper.GsgmToLutrisRunScript(path, info, setting)
	lutrisGame := mapper.GsgmToLutrisLutrisGameDao(path, info, setting)
	lutrisCategory := mapper.GsgmToLutrisLutrisCategoryDao(packName)

	// script
	if force {
		err = lutris_service.InstallRunScript(info.Id, runScript)
	} else {
		err = lutris_service.UpsertRunScript(info.Id, runScript)
	}
	if err != nil {
		logger.Erro(err)
	}

	// db
	if force {
		err = lutris_service.InstallLutrisDb(lutrisGame, lutrisCategory)
	} else {
		err = lutris_service.UpsertLutrisDb(lutrisGame, lutrisCategory)
	}
	if err != nil {
		logger.Erro(err)
	}

	installFunc := []func(int64, string) error{
		lutris_service.InstallGameCoverart,
		lutris_service.InstallGameBanner,
		lutris_service.InstallGameIcon,
	}
	upsertFunc := []func(int64, string) error{
		lutris_service.UpsertGameCoverart,
		lutris_service.UpsertGameBanner,
		lutris_service.UpsertGameIcon,
	}

	// img
	err = gsgm_service.CheckImg(path)
	funcLen := len(installFunc)
	if safeMode || err != nil {
		funcLen = 0
	}
	var wg sync.WaitGroup
	wg.Add(funcLen)
	errArr := make([]error, funcLen)
	for i := 0; i < funcLen; i++ {
		i := i
		go func() {
			defer wg.Done()
			if force {
				errArr[i] = installFunc[i](info.Id, path)
			} else {
				errArr[i] = upsertFunc[i](info.Id, path)
			}
		}()
	}
	wg.Wait()

	for _, err := range errArr {
		if err == nil {
			continue
		}
		logger.Erro(err)
	}

	fmt.Println("completed => " + path)
}
