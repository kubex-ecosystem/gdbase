package types

import (
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type Redis struct {
	Reference *t.Reference      `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath  string            `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled   bool              `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Addr      string            `gorm:"omitempty" json:"addr" yaml:"addr" xml:"addr" toml:"addr" mapstructure:"addr"`
	Port      interface{}       `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username  string            `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password  string            `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	DB        interface{}       `gorm:"omitempty" json:"db" yaml:"db" xml:"db" toml:"db" mapstructure:"db"`
	Volume    string            `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	Mapper    *t.Mapper[*Redis] `json:"-" yaml:"-" xml:"-" toml:"-" mapstructure:"-"`
}
