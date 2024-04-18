package lutris_service

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_script"
	"github.com/XDwanj/go-gsgm/util"
	"github.com/duke-git/lancet/v2/fileutil"
	"gopkg.in/yaml.v3"
)

func UpsertRunScript(gsgmId int64, script *lutris_script.LutrisRunScript) error {
	scriptPath := filepath.Join(config.RunScriptPath, config.SlugPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.ScriptSuffix)
	logger.Info(scriptPath)
	if fileutil.IsExist(scriptPath) {
		return nil
	}
	return InstallRunScript(gsgmId, script)
}

func InstallRunScript(gsgmId int64, script *lutris_script.LutrisRunScript) error {
	scriptPath := filepath.Join(config.RunScriptPath, config.SlugPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.ScriptSuffix)
	bytes, err := yaml.Marshal(script)
	if err != nil {
		return err
	}
	return fileutil.WriteStringToFile(scriptPath, string(bytes), false)
}

func CleanLutrisRunScript() error {
	paths, err := util.Ls(config.RunScriptPath)
	if err != nil {
		return err
	}
	for _, path := range paths {
		if !strings.HasPrefix(filepath.Base(path), config.SlugPrefix) {
			continue
		}
		logger.Info("rm runScript ", path)
		if err := fileutil.RemoveFile(path); err != nil {
			logger.Erro(err)
		}
	}
	return nil
}

func RemoveLutrisRunScriptBySlug(slug string) error {
	curScriptPath := filepath.Join(config.RunScriptPath, slug+"."+config.ScriptSuffix)
	return fileutil.RemoveFile(curScriptPath)
}
