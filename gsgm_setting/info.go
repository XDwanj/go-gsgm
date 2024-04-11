package gsgm_setting

type GsgmInfo struct {
	Id       int64 `json:"id,omitempty"`
	InitTime int64 `json:"initTime,omitempty"` // Unix时间戳 单位秒

	// deprecated
	Description string `json:"description,omitempty"`
}
