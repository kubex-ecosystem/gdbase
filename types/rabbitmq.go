package types

import (
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type RabbitMQ struct {
	Reference      *t.Reference         `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath       string               `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled        bool                 `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Username       string               `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password       string               `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Vhost          string               `gorm:"omitempty" json:"vhost" yaml:"vhost" xml:"vhost" toml:"vhost" mapstructure:"vhost"`
	Port           interface{}          `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Host           string               `gorm:"omitempty" json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Volume         string               `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	ErlangCookie   string               `gorm:"omitempty" json:"erlang_cookie" yaml:"erlang_cookie" xml:"erlang_cookie" toml:"erlang_cookie" mapstructure:"erlang_cookie"`
	ManagementUser string               `gorm:"omitempty" json:"management_user" yaml:"management_user" xml:"management_user" toml:"management_user" mapstructure:"management_user"`
	ManagementPass string               `gorm:"omitempty" json:"management_pass" yaml:"management_pass" xml:"management_pass" toml:"management_pass" mapstructure:"management_pass"`
	ManagementHost string               `gorm:"omitempty" json:"management_host" yaml:"management_host" xml:"management_host" toml:"management_host" mapstructure:"management_host"`
	ManagementPort string               `gorm:"omitempty" json:"management_port" yaml:"management_port" xml:"management_port" toml:"management_port" mapstructure:"management_port"`
	Mapper         *t.Mapper[*RabbitMQ] `json:"-" yaml:"-" xml:"-" toml:"-" mapstructure:"-"`
}
