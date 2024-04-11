package gsgm_setting

import "time"

// deprecated: 暂时不支持游戏游玩记录
type GsgmHistory struct {
	LastGameMoment time.Time `json:"lastGameMoment,omitempty"`
	GameTime       float64   `json:"gameTime,omitempty"`
}
