package lutris_dao

import (
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
	hidden INTEGER,
	service TEXT,
	service_id TEXT,
	discord_id TEXT

)
*/
type LutrisGame struct {
	Id                   int64      `json:"id"`
	Name                 string     `json:"name"`
	Sortname             string     `json:"sortname"`
	Slug                 string     `json:"slug"`
	InstallerSlug        string     `json:"installer_slug"`
	ParentSlug           string     `json:"parent_slug"`
	Platform             string     `json:"platform"`
	Runner               string     `json:"runner"`
	Executable           string     `json:"executable"`
	Directory            string     `json:"directory"`
	Updated              *time.Time `json:"updated"`
	Lastplayed           int64      `json:"lastplayed"`
	Installed            int        `json:"installed"`
	InstalledAt          int64      `json:"installed_at"`
	Year                 int        `json:"year"`
	Configpath           string     `json:"configpath"`
	HasCustomBanner      int        `json:"has_custom_banner"`
	HasCustomIcon        int        `json:"has_custom_icon"`
	HasCustomCoverartBig int        `json:"has_custom_coverart_big"`
	Playtime             float64    `json:"playtime"`
	Hidden               int        `json:"hidden"`
	Service              string     `json:"service"`
	ServiceId            string     `json:"service_id"`
	DiscordId            string     `json:"discord_id"`
}
