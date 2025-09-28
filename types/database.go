package types

import (
	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type DBService = IDBService

type JWT struct {
	Reference             *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath              string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	RefreshSecret         string       `gorm:"omitempty" json:"refresh_secret" yaml:"refresh_secret" xml:"refresh_secret" toml:"refresh_secret" mapstructure:"refresh_secret"`
	PrivateKey            string       `gorm:"omitempty" json:"private_key" yaml:"private_key" xml:"private_key" toml:"private_key" mapstructure:"private_key"`
	PublicKey             string       `gorm:"omitempty" json:"public_key" yaml:"public_key" xml:"public_key" toml:"public_key" mapstructure:"public_key"`
	ExpiresIn             int          `gorm:"omitempty" json:"expires_in" yaml:"expires_in" xml:"expires_in" toml:"expires_in" mapstructure:"expires_in"`
	IDExpirationSecs      int          `gorm:"omitempty" json:"id_expiration_secs" yaml:"id_expiration_secs" xml:"id_expiration_secs" toml:"id_expiration_secs" mapstructure:"id_expiration_secs"`
	RefreshExpirationSecs int          `gorm:"omitempty" json:"refresh_expiration_secs" yaml:"refresh_expiration_secs" xml:"refresh_expiration_secs" toml:"refresh_expiration_secs" mapstructure:"refresh_expiration_secs"`
	Mapper                ci.IMapper[JWT]
}
type Database struct {
	Reference        *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	Enabled          bool         `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	FilePath         string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Type             string       `gorm:"not null" json:"type" yaml:"type" xml:"type" toml:"type" mapstructure:"type"`
	Driver           string       `gorm:"not null" json:"driver" yaml:"driver" xml:"driver" toml:"driver" mapstructure:"driver"`
	ConnectionString string       `gorm:"omitempty" json:"connection_string" yaml:"connection_string" xml:"connection_string" toml:"connection_string" mapstructure:"connection_string"`
	Dsn              string       `gorm:"omitempty" json:"dsn" yaml:"dsn" xml:"dsn" toml:"dsn" mapstructure:"dsn"`
	Path             string       `gorm:"omitempty" json:"path" yaml:"path" xml:"path" toml:"path" mapstructure:"path"`
	Host             string       `gorm:"omitempty" json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Port             interface{}  `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username         string       `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password         string       `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Name             string       `gorm:"omitempty" json:"name" yaml:"name" xml:"name" toml:"name" mapstructure:"name"`
	Volume           string       `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	Mapper           ci.IMapper[Database]
}
type Redis struct {
	Reference *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath  string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled   bool         `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Addr      string       `gorm:"omitempty" json:"addr" yaml:"addr" xml:"addr" toml:"addr" mapstructure:"addr"`
	Port      interface{}  `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username  string       `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password  string       `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	DB        interface{}  `gorm:"omitempty" json:"db" yaml:"db" xml:"db" toml:"db" mapstructure:"db"`
	Volume    string       `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	Mapper    ci.IMapper[Redis]
}
type RabbitMQ struct {
	Reference      *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath       string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled        bool         `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Username       string       `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password       string       `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Vhost          string       `gorm:"omitempty" json:"vhost" yaml:"vhost" xml:"vhost" toml:"vhost" mapstructure:"vhost"`
	Port           interface{}  `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Host           string       `gorm:"omitempty" json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Volume         string       `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	ErlangCookie   string       `gorm:"omitempty" json:"erlang_cookie" yaml:"erlang_cookie" xml:"erlang_cookie" toml:"erlang_cookie" mapstructure:"erlang_cookie"`
	ManagementUser string       `gorm:"omitempty" json:"management_user" yaml:"management_user" xml:"management_user" toml:"management_user" mapstructure:"management_user"`
	ManagementPass string       `gorm:"omitempty" json:"management_pass" yaml:"management_pass" xml:"management_pass" toml:"management_pass" mapstructure:"management_pass"`
	ManagementHost string       `gorm:"omitempty" json:"management_host" yaml:"management_host" xml:"management_host" toml:"management_host" mapstructure:"management_host"`
	ManagementPort string       `gorm:"omitempty" json:"management_port" yaml:"management_port" xml:"management_port" toml:"management_port" mapstructure:"management_port"`
	Mapper         ci.IMapper[RabbitMQ]
}
type Messagery struct {
	RabbitMQ *RabbitMQ `json:"rabbitmq" yaml:"rabbitmq" xml:"rabbitmq" toml:"rabbitmq" mapstructure:"rabbitmq"`
	Redis    *Redis    `json:"redis" yaml:"redis" xml:"redis" toml:"redis" mapstructure:"redis"`
	//*Kafka    `json:"kafka" yaml:"kafka" xml:"kafka" toml:"kafka" mapstructure:"kafka"`
	//*Nats     `json:"nats" yaml:"nats" xml:"nats" toml:"nats" mapstructure:"nats"`
	//*ActiveMQ `json:"activemq" yaml:"activemq" xml:"activemq" toml:"activemq" mapstructure:"activemq"`
	//*AMQP     `json:"amqp" yaml:"amqp" xml:"amqp" toml:"amqp" mapstructure:"amqp"`
	Mapper ci.IMapper[Messagery]
}
