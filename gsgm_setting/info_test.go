package gsgm_setting

import (
	"fmt"
	"testing"
	"time"

	"github.com/XDwanj/go-gsgm/util"
	"github.com/duke-git/lancet/v2/formatter"
)

func TestPrintlnGsgmInfoJson(t *testing.T) {
	info := &GsgmInfo{
		Id:          util.NextSnowId(),
		InitTime:    time.Now().Unix(),
		Description: "",
	}
	json, _ := formatter.Pretty(info)
	fmt.Println(string(json))
}
