package contro

import (
	"fmt"
	"slices"
	"sync"

	"github.com/XDwanj/go-gsgm/gsgm_dto"
	"github.com/XDwanj/go-gsgm/gsgm_service"
	"github.com/jedib0t/go-pretty/v6/text"
)

func CheckByLibraries(libraryPaths []string) {
	packs := make([]gsgm_dto.GamePack, 0, 100)
	for _, libPath := range libraryPaths {
		packs = slices.Concat(packs, gsgm_service.DeepGamePack(libPath))
	}

	for _, pack := range packs {
		for _, path := range pack.Paths {
			if CheckOne(path) {
				fmt.Println(text.BgGreen.Sprint("check >>>"), path)
			} else {
				fmt.Println(text.BgYellow.Sprint("check >>>"), path)
			}
		}
	}
}

func CheckBySingles(paths []string) {
	for _, path := range paths {
		fmt.Println(text.BgGreen.Sprint("check >>>"), path)
		CheckOne(path)
	}
}

func CheckOne(path string) bool {
	checkFuncs := [](func(path string) error){
		gsgm_service.CheckInfo,
		gsgm_service.CheckSetting,
		// gsgm_service.CheckImg,
	}

	checkFuncLen := len(checkFuncs)
	errorArr := make([]error, checkFuncLen)

	var waitGroup sync.WaitGroup
	waitGroup.Add(checkFuncLen)
	for i := range checkFuncs {
		i := i
		go func() {
			defer waitGroup.Done()
			errorArr[i] = checkFuncs[i](path)
		}()
	}
	waitGroup.Wait()

	errorArr = slices.DeleteFunc(errorArr, func(err error) bool {
		return err == nil
	})

	for _, err := range errorArr {
		fmt.Println(text.BgRed.Sprint("err -> "), text.BgYellow.Sprint(err))
	}

	return len(errorArr) == 0
}
