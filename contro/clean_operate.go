package contro

import (
	"fmt"

	"github.com/XDwanj/go-gsgm/lutris_service"
)

func CleanAction() {
	// img
	imgFlag := true
	if lutris_service.CleanLutrisCover() != nil {
		imgFlag = false
	}
	if lutris_service.CleanLutrisBanner() != nil {
		imgFlag = false
	}
	if lutris_service.CleanLutrisIcon() != nil {
		imgFlag = false
	}
	if imgFlag {
		fmt.Println("clean image successed")
	} else {
		fmt.Println("clean image failed")
	}

	// script
	if err := lutris_service.CleanLutrisRunScript(); err != nil {
		fmt.Println("clean script failed:", err)
	} else {
		fmt.Println("clean script successed")
	}

	// db
	if err := lutris_service.CleanLutrisDb(); err != nil {
		fmt.Println("clean db failed:", err)
	} else {
		fmt.Println("clean db successed")
	}
}
