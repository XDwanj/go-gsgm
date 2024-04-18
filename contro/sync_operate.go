package contro

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
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

func SyncByLibraries(libPaths []string, force bool) {
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
				SyncByOne(path, force)
			}()
		}
		wg.Wait()
	}

}

func SyncBySingles(paths []string, force bool) {

	var wg sync.WaitGroup
	wg.Add(len(paths))
	for _, path := range paths {
		path := path
		go func() {
			defer wg.Done()
			SyncByOne(path, force)
		}()
	}
	wg.Wait()
}

func SyncByOne(path string, force bool) {
	historyMap := getHistoryMap()
	name := filepath.Base(path)
	info, err := gsgm_service.GetGsgmInfoByPath(path)
	if err != nil {
		logger.Erro(err)
		return
	}
	// db
	historyByDb := historyMap[info.Id]
	if historyByDb == nil {
		return
	}

	// forceFunc
	forceWrite := func() {
		json, err := formatter.Pretty(historyByDb)
		if err != nil {
			logger.Erro(err)
		}
		historyPath := filepath.Join(path, config.GsgmDirName, config.GsgmHistoryName)
		if err := fileutil.WriteStringToFile(historyPath, json, false); err != nil {
			logger.Erro(err)
		}
		fmt.Printf("sync [%v] success content: %v\n", name, strutil.RemoveWhiteSpace(json, false))
	}

	// force
	if force {
		forceWrite()
		return
	}

	historyByDisk, err := gsgm_service.GetGsgmHistoryByPath(path)
	if err != nil {
		forceWrite()
		return
	}

	// equals yes
	if historyByDb.LastPlayedTime == historyByDisk.LastPlayedTime &&
		historyByDb.PlayedDuration == historyByDisk.LastPlayedTime {
		return
	}

	// final
	forceWrite()
}

var (
	_dbAction   sync.Once
	_historyMap map[int64]*gsgm_setting.GsgmHistory
)

func getHistoryMap() map[int64]*gsgm_setting.GsgmHistory {
	_dbAction.Do(func() {
		games, err := lutris_service.ListNameAndLastplayedAndPlaytime()
		if err != nil {
			panic(err)
		}
		_historyMap = make(map[int64]*gsgm_setting.GsgmHistory, len(games))

		var wg sync.WaitGroup
		wg.Add(len(games))
		for _, game := range games {

			var (
				gsgmId         int64 = 0
				lastPlayedTime int64 = 0
				playedDuration int64 = 0
			)
			if !game.Slug.Valid {
				panic("slug 不可能为空")
			}
			gsgmId, err := strconv.ParseInt(game.Slug.String[5:], 10, 64)
			if err != nil {
				panic(err)
			}
			if game.Lastplayed.Valid {
				lastPlayedTime = game.Lastplayed.Int64
				playedDuration = int64(game.Playtime.Float64 * 60)
			}

			history := &gsgm_setting.GsgmHistory{
				LastPlayedTime: lastPlayedTime,
				PlayedDuration: playedDuration,
			}

			_historyMap[gsgmId] = history
		}
	})
	return _historyMap
}
