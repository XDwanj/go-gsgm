package lutris_script

type LutrisInstallScript struct {
	Name     string     `yaml:"name,omitempty"`
	GameSlug string     `yaml:"game_slug,omitempty"`
	Version  string     `yaml:"version,omitempty"`
	Slug     string     `yaml:"slug,omitempty"`
	Runner   string     `yaml:"runner,omitempty"`
	Script   GameScript `yaml:"script,omitempty"`
}

type GameScript struct {
	Game GameDetail `yaml:"game,omitempty"`
}
