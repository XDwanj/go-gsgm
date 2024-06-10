package contro

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/XDwanj/go-gsgm/gsgm_dto"
	"github.com/XDwanj/go-gsgm/gsgm_service"
	"github.com/jedib0t/go-pretty/v6/text"
)

func ScanByLibraries(libPaths []string) {
	packs := make([]gsgm_dto.GamePack, 0)
	for _, libPath := range libPaths {
		packs = slices.Concat(packs, gsgm_service.DeepGamePack(libPath))
	}

	for _, pack := range packs {
		packName := pack.PackName
		for _, path := range pack.Paths {
			info, _ := gsgm_service.GetGsgmInfoByPath(path)
			setting, _ := gsgm_service.GetGsgmSettingByPath(path)

			var builder strings.Builder
			builder.WriteString(text.FgGreen.Sprint(packName))
			builder.WriteString("/")
			builder.WriteString(filepath.Base(path))
			builder.WriteString(" ")
			builder.WriteString(text.FgCyan.Sprint(string(setting.Platform)))
			builder.WriteString(" ")
			builder.WriteString(text.FgBlue.Sprint(strconv.FormatInt(info.Id, 10)))
			builder.WriteString(" ")
			builder.WriteString("alone=" + strconv.FormatBool(setting.PrefixAlone))
			builder.WriteString("\n")
			builder.WriteString("\t")
			builder.WriteString(path)

			fmt.Println(builder.String())
		}
	}
}
