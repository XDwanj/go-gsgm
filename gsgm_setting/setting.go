package gsgm_setting

import "errors"

type GsgmSetting struct {
	Execute       string        `json:"execute,omitempty"`
	PrefixAlone   bool          `json:"prefixAlone"` // json 默认打印
	LocaleCharSet LocaleCharSet `json:"localeCharSet,omitempty"`
	Platform      Platform      `json:"platform,omitempty"`
	Runner        Runner        `json:"runner,omitempty"`
	PreCommand    string        `json:"preCommand,omitempty"`
	ArgsCommand   string        `json:"argsCommand,omitempty"`
	EndCommand    string        `json:"endCommand,omitempty"`

	// deprecated: 不再需要单独设置 prefix，统一放在 ~/.local/share/gsgm/prefix/
	WinePrefix string `json:"winePrefix,omitempty"`
}

type LocaleCharSet string

const (
	ChinaGBK     LocaleCharSet = "zh_CN.GBK"
	ChinaGB18030 LocaleCharSet = "zh_CN.GB18030"
	ChinaUTF8    LocaleCharSet = "zh_CN.UTF-8"
	JapanEucJp   LocaleCharSet = "ja_JP.EUC-JP"
	JapanUTF8    LocaleCharSet = "ja_JP.UTF-8"
)

type Platform string

const (
	Windows Platform = "Windows"
	Linux   Platform = "Linux"
	WebApp  Platform = "网页"
)

func (p Platform) DefaultRunner() (Runner, error) {
	switch p {
	case Windows:
		return Wine, nil
	case Linux:
		return LinuxNative, nil
	case WebApp:
		return Web, nil
	default:
		return "", errors.New("找不到类型")
	}
}

type Runner string

const (
	LinuxNative Runner = "linux"
	Wine        Runner = "wine"
	Flatpak     Runner = "flatpak"
	Web         Runner = "web"
)
