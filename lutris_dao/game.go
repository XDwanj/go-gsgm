package lutris_dao

import (
	"database/sql"
	"time"
)

/*
CREATE TABLE games (

	id INTEGER PRIMARY KEY,
	name TEXT,
	sortname TEXT,
	slug TEXT,
	installer_slug TEXT,
	parent_slug TEXT,
	platform TEXT,
	runner TEXT,
	executable TEXT,
	directory TEXT,
	updated DATETIME,
	lastplayed INTEGER,
	installed INTEGER,
	installed_at INTEGER,
	year INTEGER,
	configpath TEXT,
	has_custom_banner INTEGER,
	has_custom_icon INTEGER,
	has_custom_coverart_big INTEGER,
	playtime REAL,
	// hidden INTEGER, // deprecated
	service TEXT,
	service_id TEXT,
	discord_id TEXT

)
*/
type LutrisGame struct {
	Id                   int64           `json:"id"`
	Name                 sql.NullString  `json:"name"`                    // 游戏名
	Sortname             sql.NullString  `json:"sortname"`                // 排序名 default blank
	Slug                 sql.NullString  `json:"slug"`                    // gsgm-xxxxx
	InstallerSlug        sql.NullString  `json:"installer_slug"`          // default blank
	ParentSlug           sql.NullString  `json:"parent_slug"`             // default blank
	Platform             sql.NullString  `json:"platform"`                // 游戏平台，可选：Windows，Linux
	Runner               sql.NullString  `json:"runner"`                  // 运行器，可选：linux，wine，flatpak
	Executable           sql.NullString  `json:"executable"`              // 可执行文件位置，default blank
	Directory            sql.NullString  `json:"directory"`               // 游戏工作目录
	Updated              *time.Time      `json:"updated"`                 // 更新时间，这里我们默认当前安装时间
	Lastplayed           sql.NullInt64   `json:"lastplayed"`              // 上次游玩时间 Unix 秒时间戳
	Installed            sql.NullInt32   `json:"installed"`               // 是否安装，1 表示安装
	InstalledAt          sql.NullInt64   `json:"installed_at"`            // 安装时间，Unix 时间戳，单位秒
	Year                 sql.NullInt32   `json:"year"`                    // 游戏年份，default 0
	Configpath           sql.NullString  `json:"configpath"`              // 执行配置文件名字 gsgm-xxxx
	HasCustomBanner      sql.NullInt32   `json:"has_custom_banner"`       // 是否有 banner
	HasCustomIcon        sql.NullInt32   `json:"has_custom_icon"`         // 是否有 icon
	HasCustomCoverartBig sql.NullInt32   `json:"has_custom_coverart_big"` // 是否有 cover
	Playtime             sql.NullFloat64 `json:"playtime"`                // 游玩时间，小数，单位小时
	Service              sql.NullString  `json:"service"`                 // 游戏服务提供商，如：Steam
	ServiceId            sql.NullString  `json:"service_id"`              // 游戏服务提供商id
	DiscordId            sql.NullString  `json:"discord_id"`              // dicord频道id

	// deprecated
	// Hidden sql.NullInt32 `json:"hidden"` // 是否隐藏
}
