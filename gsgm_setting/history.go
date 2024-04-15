package gsgm_setting

type GsgmHistory struct {
	LastPlayedTime int64 `json:"lastPlayedTime,omitempty"` // 上次游玩时刻，单位：秒
	PlayedDuration int64 `json:"playedDuration,omitempty"` // 游玩时间，单位：分钟
}
