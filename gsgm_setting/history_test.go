package gsgm_setting

import (
	"fmt"
	"testing"
	"time"

	"github.com/XDwanj/go-gsgm/logger"
	"github.com/duke-git/lancet/v2/formatter"
)

func TestPrint(t *testing.T) {
	history := &GsgmHistory{
		LastPlayedTime: time.Now().Unix(),
		PlayedDuration: int64((6 * time.Hour).Minutes()),
	}

	json, err := formatter.Pretty(history)
	if err != nil {
		logger.Erro(err)
	}
	fmt.Println(json)
}
