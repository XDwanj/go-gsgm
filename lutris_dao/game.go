package lutris_dao

import (
	"time"
)

type LutrisGame struct {
	Id                   int64      `gorm:"column:id;type:INTEGER;primaryKey;" json:"id"`
	Name                 string     `gorm:"column:name;type:TEXT;" json:"name"`
	Sortname             string     `gorm:"column:sortname;type:TEXT;" json:"sortname"`
	Slug                 string     `gorm:"column:slug;type:TEXT;" json:"slug"`
	InstallerSlug        string     `gorm:"column:installer_slug;type:TEXT;" json:"installer_slug"`
	ParentSlug           string     `gorm:"column:parent_slug;type:TEXT;" json:"parent_slug"`
	Platform             string     `gorm:"column:platform;type:TEXT;" json:"platform"`
	Runner               string     `gorm:"column:runner;type:TEXT;" json:"runner"`
	Executable           string     `gorm:"column:executable;type:TEXT;" json:"executable"`
	Directory            string     `gorm:"column:directory;type:TEXT;" json:"directory"`
	Updated              *time.Time `gorm:"column:updated;type:DATETIME;" json:"updated"`
	Lastplayed           int64      `gorm:"column:lastplayed;type:INTEGER;" json:"lastplayed"`
	Installed            int        `gorm:"column:installed;type:INTEGER;" json:"installed"`
	InstalledAt          int64      `gorm:"column:installed_at;type:INTEGER;" json:"installed_at"`
	Year                 int        `gorm:"column:year;type:INTEGER;" json:"year"`
	Configpath           string     `gorm:"column:configpath;type:TEXT;" json:"configpath"`
	HasCustomBanner      int        `gorm:"column:has_custom_banner;type:INTEGER;" json:"has_custom_banner"`
	HasCustomIcon        int        `gorm:"column:has_custom_icon;type:INTEGER;" json:"has_custom_icon"`
	HasCustomCoverartBig int        `gorm:"column:has_custom_coverart_big;type:INTEGER;" json:"has_custom_coverart_big"`
	Playtime             float64    `gorm:"column:playtime;type:REAL;" json:"playtime"`
	Hidden               int        `gorm:"column:hidden;type:INTEGER;" json:"hidden"`
	Service              string     `gorm:"column:service;type:TEXT;" json:"service"`
	ServiceId            string     `gorm:"column:service_id;type:TEXT;" json:"service_id"`
	DiscordId            string     `gorm:"column:discord_id;type:TEXT;" json:"discord_id"`
}

func (l *LutrisGame) TableName() string {
	return "games"
}
