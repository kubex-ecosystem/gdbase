package types

import ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"

type Messagery struct {
	RabbitMQ *RabbitMQ `json:"rabbitmq" yaml:"rabbitmq" xml:"rabbitmq" toml:"rabbitmq" mapstructure:"rabbitmq"`
	Redis    *Redis    `json:"redis" yaml:"redis" xml:"redis" toml:"redis" mapstructure:"redis"`
	//*Kafka    `json:"kafka" yaml:"kafka" xml:"kafka" toml:"kafka" mapstructure:"kafka"`
	//*Nats     `json:"nats" yaml:"nats" xml:"nats" toml:"nats" mapstructure:"nats"`
	//*ActiveMQ `json:"activemq" yaml:"activemq" xml:"activemq" toml:"activemq" mapstructure:"activemq"`
	//*AMQP     `json:"amqp" yaml:"amqp" xml:"amqp" toml:"amqp" mapstructure:"amqp"`
	Mapper ci.IMapper[Messagery]
}
