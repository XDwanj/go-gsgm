package lutris_service

import (
	"encoding/json"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/dao"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_dao"
	"gorm.io/gorm"
)

func UpsertLutrisDb(
	game *lutris_dao.LutrisGame,
	category *lutris_dao.LutrisCategory,
) error {

	dao.LutrisDb.Session(&gorm.Session{SkipDefaultTransaction: true})
	return dao.LutrisDb.Transaction(func(tx *gorm.DB) error {
		err := tx.FirstOrCreate(game, lutris_dao.LutrisGame{Slug: game.Slug}).Error
		if err != nil {
			return err
		}
		err = tx.FirstOrCreate(category, lutris_dao.LutrisCategory{Name: category.Name}).Error
		if err != nil {
			return err
		}

		lutrisGameId := game.Id
		lutrisCategoryId := category.Id

		rel := &lutris_dao.LutrisRelGameToCategory{
			GameId:     lutrisGameId,
			CategoryId: lutrisCategoryId,
		}

		err = tx.Create(rel).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func CleanLutrisDb() error {
	var games []lutris_dao.LutrisGame
	if err := dao.LutrisDb.Model(&lutris_dao.LutrisGame{}).Select("id").Where("slug like ?", config.SlugPrefix+"%").Find(&games).Error; err != nil {
		return err
	}
	ids := make([]int64, len(games))
	for i := range games {
		ids[i] = games[i].Id
	}

	err := dao.LutrisDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("slug like ?", config.SlugPrefix+"%").Delete(&lutris_dao.LutrisGame{}).Error; err != nil {
			return err
		}
		if err := tx.Where("name like ?", "@%").Or("name like ?", "$%").Delete(&lutris_dao.LutrisCategory{}).Error; err != nil {
			return err
		}
		if err := tx.Where("game_id in ?", ids).Delete(&lutris_dao.LutrisRelGameToCategory{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	for i := range games {
		bytes, _ := json.Marshal(games[i])
		json := string(bytes)
		if len(json) > 50 {
			json = json[:50] + "..."
		}
		logger.Info("rm db ", json)
	}

	return nil
}
