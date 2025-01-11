package lutris_service

import (
	"fmt"
	"testing"
	"time"

	"github.com/XDwanj/go-gsgm/dao"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_dao"
	"gorm.io/gorm"
)

func TestListGame(t *testing.T) {
	logger.Level = logger.Infomation
	games, err := ListNameAndLastplayedAndPlaytime()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, game := range games {
		fmt.Println(game.Name, game.Lastplayed, game.Playtime)
	}
}

func TestTime(t *testing.T) {
	ti := (time.Duration(32) * time.Minute).Hours()
	fmt.Printf("ti: %v\n", ti)
}

func Test(t *testing.T) {
	db := dao.LutrisDb
	db = db.Session(&gorm.Session{DryRun: true})
	var gameId string
	_ = db.Model(&lutris_dao.LutrisGame{}).
		Select("id").
		Where("slug = ?", "hhhh").
		First(&gameId).Statement
	// fmt.Printf("stmt.SQL.String(): %v\n", stmt.SQL.String())

}
