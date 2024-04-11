package gsgm_setting

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/formatter"
)

func TestPrintSetting(t *testing.T) {
	setting := &GsgmSetting{
		ExecuteLocation: "run.exe",
		WinePrefixAlone: false,
		LocaleCharSet:   ChinaUTF8,
		Platform:        Windows,
		Runner:          Wine,
	}

	json, err := formatter.Pretty(setting)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println(json)
}
