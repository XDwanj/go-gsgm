package contro

import (
	"fmt"
	"path/filepath"
	"slices"
	"sync"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_dto"
	"github.com/XDwanj/go-gsgm/gsgm_service"
	"github.com/XDwanj/go-gsgm/gsgm_setting"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_service"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/formatter"
	"github.com/duke-git/lancet/v2/strutil"
)

func SyncByLibraries(libPaths []string) {
	packs := make([]gsgm_dto.GamePack, 0)
	for _, libPath := range libPaths {
		packs = slices.Concat(packs, gsgm_service.DeepGamePack(libPath))
	}

	for _, pack := range packs {
		var wg sync.WaitGroup
		wg.Add(len(pack.Paths))

		for _, path := range pack.Paths {
			path := path
			go func() {
				defer wg.Done()
				SyncByOne(path)
			}()
		}
		wg.Wait()
	}

}

func SyncBySingles(paths []string) {

	var wg sync.WaitGroup
	wg.Add(len(paths))
	for _, path := range paths {
		path := path
		go func() {
			defer wg.Done()
			SyncByOne(path)
		}()
	}
	wg.Wait()
}

func SyncByOne(path string) {
	historyDic := GetGameMap()
	name := filepath.Base(path)
	historyByDb := historyDic[name]
	if historyByDb == nil {
		return
	}

	json, err := formatter.Pretty(historyByDb)
	if err != nil {
		logger.Erro(fmt.Sprintf("sync [%s] error:", name), err)
	}

	historyPath := filepath.Join(path, config.GsgmDirName, config.GsgmHistoryName)
	if err := fileutil.WriteStringToFile(historyPath, json, false); err != nil {
		logger.Erro(fmt.Sprintf("sync [%s] error:", name), err)
	}
	fmt.Printf("sync [%v] success content: %v\n", name, strutil.RemoveWhiteSpace(json, false))
}

var _action sync.Once
var _historyMap map[string]*gsgm_setting.GsgmHistory

func GetGameMap() map[string]*gsgm_setting.GsgmHistory {
	_action.Do(func() {
		games, err := lutris_service.ListNameAndLastplayedAndPlaytime()
		if err != nil {
			panic(err)
		}
		_historyMap = make(map[string]*gsgm_setting.GsgmHistory, len(games))

		var wg sync.WaitGroup
		wg.Add(len(games))
		for _, game := range games {

			var (
				lastPlayedTime int64 = 0
				playedDuration int64 = 0
			)
			if game.Lastplayed.Valid {
				lastPlayedTime = game.Lastplayed.Int64
				playedDuration = int64(game.Playtime.Float64 * 60)
			}

			history := &gsgm_setting.GsgmHistory{
				LastPlayedTime: lastPlayedTime,
				PlayedDuration: playedDuration,
			}

			_historyMap[game.Name.String] = history
		}
	})
	return _historyMap
}
