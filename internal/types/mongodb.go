package types

type MongoDB struct {
	Reference *Reference        `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath  string            `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled   bool              `json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Host      string            `json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Port      interface{}       `json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username  string            `json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password  string            `json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Mapper    *Mapper[*MongoDB] `json:"-" yaml:"-" xml:"-" toml:"-" mapstructure:"-"`
}
