package lutris_script

type LutrisRunScript struct {
	Slug     string        `yaml:"slug,omitempty"`
	GameSlug string        `yaml:"game_slug,omitempty"`
	Game     *GameDetail   `yaml:"game,omitempty"`
	System   *SystemDetail `yaml:"system,omitempty"`
	Wine     *WineDetail   `yaml:"wine,omitempty"`
}

type GameDetail struct {
	Exe        string `yaml:"exe,omitempty"`
	Prefix     string `yaml:"prefix,omitempty"`
	WorkingDir string `yaml:"working_dir,omitempty"`
}

type SystemDetail struct {
	Locale string `yaml:"locale,omitempty"`
}

type WineDetail struct {
	Version string `yaml:"version,omitempty"`
}
