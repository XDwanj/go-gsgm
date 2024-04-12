package lutris_service

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/gsgm_service"
	"github.com/XDwanj/go-gsgm/img_service"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/duke-git/lancet/v2/fileutil"
)

func UpsertGameBanner(gsgmId int64, gamePath string) error {
	destPath := filepath.Join(config.BannerPath, config.BannerPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.BannerSuffix)
	logger.Info("upsert banner try ", destPath)
	if fileutil.IsExist(destPath) {
		return nil
	}
	return InstallGameBanner(gsgmId, gamePath)
}

func InstallGameBanner(gsgmId int64, gamePath string) error {
	destPath := filepath.Join(config.BannerPath, config.BannerPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.BannerSuffix)
	srcPath, err := gsgm_service.GetImgPath(gamePath)
	logger.Info("try install banner ", destPath)
	if err != nil {
		return err
	}

	logger.Info("dest banner ", destPath)
	return img_service.ZoomLutrisPicture(
		srcPath,
		destPath,
		config.BannerStd.Width,
		config.BannerStd.Height,
	)
}

func UpsertGameCoverart(gsgmId int64, gamePath string) error {
	destPath := filepath.Join(config.CoverartPath, config.CoverartPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.CoverartSuffix)
	if fileutil.IsExist(destPath) {
		return nil
	}
	return InstallGameCoverart(gsgmId, gamePath)
}

func InstallGameCoverart(gsgmId int64, gamePath string) error {
	destPath := filepath.Join(config.CoverartPath, config.CoverartPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.CoverartSuffix)
	srcPath, err := gsgm_service.GetImgPath(gamePath)
	if err != nil {
		return err
	}

	return img_service.ZoomLutrisPicture(
		srcPath,
		destPath,
		config.CoverartStd.Width,
		config.CoverartStd.Height,
	)
}

func UpsertGameIcon(gsgmId int64, gamePath string) error {
	destPath := filepath.Join(config.IconPath, config.IconPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.IconSuffix)
	logger.Info("try upsert icon ", destPath)
	if fileutil.IsExist(destPath) {
		return nil
	}
	logger.Info("try upsert icon ", destPath)
	return InstallGameIcon(gsgmId, gamePath)
}

func InstallGameIcon(gsgmId int64, gamePath string) error {
	destPath := filepath.Join(config.IconPath, config.IconPrefix+strconv.FormatInt(gsgmId, 10)+"."+config.IconSuffix)
	srcPath, err := gsgm_service.GetImgPath(gamePath)
	logger.Info("try install icon ", destPath)
	if err != nil {
		return err
	}

	logger.Info("dest icon ", destPath)
	return img_service.ZoomLutrisPicture(
		srcPath,
		destPath,
		config.IconStd.Width,
		config.IconStd.Height,
	)
}

func CleanLutrisCover() error {
	names, err := fileutil.ListFileNames(config.CoverartPath)
	if err != nil {
		return err
	}
	for _, name := range names {
		if !strings.HasPrefix(name, config.CoverartPrefix) {
			continue
		}
		path := filepath.Join(config.CoverartPath, name)
		logger.Info("try rm cover ", path)
		if err := fileutil.RemoveFile(path); err != nil {
			return err
		}
	}
	return nil
}

func CleanLutrisBanner() error {
	names, err := fileutil.ListFileNames(config.BannerPath)
	if err != nil {
		return err
	}
	for _, name := range names {
		if !strings.HasPrefix(name, config.BannerPrefix) {
			continue
		}
		path := filepath.Join(config.BannerPath, name)
		logger.Info("try rm banner ", path)
		if err := fileutil.RemoveFile(path); err != nil {
			return err
		}
	}
	return nil
}

func CleanLutrisIcon() error {
	names, err := fileutil.ListFileNames(config.IconPath)
	if err != nil {
		return err
	}
	for _, name := range names {
		if !strings.HasPrefix(name, config.IconPrefix) {
			continue
		}
		path := filepath.Join(config.IconPath, name)
		logger.Info("try rm icon ", path)
		if err := fileutil.RemoveFile(path); err != nil {
			return err
		}
	}
	return nil
}

func RemoveLutrisBannerBySlug(slug string) error {
	curBannerPath := filepath.Join(config.BannerPath, slug+"."+config.BannerSuffix)
	return fileutil.RemoveFile(curBannerPath)
}

func RemoveLutrisCoverBySlug(slug string) error {
	curCoverPath := filepath.Join(config.CoverartPath, slug+"."+config.CoverartSuffix)
	return fileutil.RemoveFile(curCoverPath)
}

func RemoveLutrisIconBySlug(slug string) error {
	curIconPath := filepath.Join(config.IconPath, "lutris_"+slug+"."+config.IconSuffix)
	return fileutil.RemoveFile(curIconPath)
}
