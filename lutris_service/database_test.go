package lutris_service

import (
	"fmt"
	"testing"

	"github.com/XDwanj/go-gsgm/logger"
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
