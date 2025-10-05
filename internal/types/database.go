package types

type Database struct {
	Reference        *Reference         `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	IsDefault        bool               `gorm:"default:false" json:"is_default" yaml:"is_default" xml:"is_default" toml:"is_default" mapstructure:"is_default"`
	Enabled          bool               `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	FilePath         string             `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Type             string             `gorm:"not null" json:"type" yaml:"type" xml:"type" toml:"type" mapstructure:"type"`
	Driver           string             `gorm:"not null" json:"driver" yaml:"driver" xml:"driver" toml:"driver" mapstructure:"driver"`
	ConnectionString string             `gorm:"omitempty" json:"connection_string" yaml:"connection_string" xml:"connection_string" toml:"connection_string" mapstructure:"connection_string"`
	Dsn              string             `gorm:"omitempty" json:"dsn" yaml:"dsn" xml:"dsn" toml:"dsn" mapstructure:"dsn"`
	Path             string             `gorm:"omitempty" json:"path" yaml:"path" xml:"path" toml:"path" mapstructure:"path"`
	Host             string             `gorm:"omitempty" json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Port             any                `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username         string             `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password         string             `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Name             string             `gorm:"omitempty" json:"name" yaml:"name" xml:"name" toml:"name" mapstructure:"name"`
	Volume           string             `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	Mapper           *Mapper[*Database] `json:"-" yaml:"-" xml:"-" toml:"-" mapstructure:"-"`
}
