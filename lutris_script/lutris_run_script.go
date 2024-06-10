package lutris_script

type LutrisRunScript struct {
	Slug     string        `yaml:"slug,omitempty"`
	GameSlug string        `yaml:"game_slug,omitempty"`
	Game     *GameDetail   `yaml:"game,omitempty"`
	System   *SystemDetail `yaml:"system,omitempty"`
	Wine     *WineDetail   `yaml:"wine,omitempty"`
}

type GameDetail struct {
	Exe        string `yaml:"exe,omitempty"`         // 可执行文件位置
	Prefix     string `yaml:"prefix,omitempty"`      // 容器前缀位置
	WorkingDir string `yaml:"working_dir,omitempty"` // 工作路径，默认游戏所在路径
}

type SystemDetail struct {
	Locale   string            `yaml:"locale,omitempty"`           // 游戏字符集
	PostExit string            `yaml:"postexit_command,omitempty"` // 游戏结束时，执行的脚本位置，切记，需可执行权限
	Env      map[string]string `yaml:"env,omitempty"`              // lutris启动游戏时，环境变量
}

type WineDetail struct {
	Version string `yaml:"version,omitempty"` // wine版本
}
