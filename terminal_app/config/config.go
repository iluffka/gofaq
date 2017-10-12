package config

type Settings struct {
	Project struct {
		Name string `yaml:"name"`
		Lang	string	`yaml:"lang"`
		UseDB	bool	`yaml:"use_db"`
		VersionControl	VC	`yaml:"version_control"`
	}
}

type VC struct {
	Type	string	`yaml:"type"`
	URL		string	`yaml:"url"`
}