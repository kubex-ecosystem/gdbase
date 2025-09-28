// Package provider defines interfaces and types for managing service providers like Postgres, MongoDB, Redis, and RabbitMQ.
package provider

import "context"

type Engine string

const (
	EnginePostgres Engine = "postgres"
	EngineMongo    Engine = "mongo"
	EngineRedis    Engine = "redis"
	EngineRabbit   Engine = "rabbitmq"
)

type ServiceRef struct {
	Name   string // "pg", "mongo", "redis", "rabbit"
	Engine Engine
}

type Endpoint struct {
	DSN      string // ex: postgres://user:pass@host:port/db?sslmode=disable
	Redacted string
	Host     string
	Port     int
}

type Capabilities struct {
	Managed  bool
	Notes    []string
	Features map[string]bool // ex: "extensions.pgcrypto": true
}

type StartSpec struct {
	Services      []ServiceRef // quais serviços subir/anexar
	PreferredPort map[string]int
	Secrets       map[string]string // senhas já geradas pelo GoBE
	Labels        map[string]string // rastreabilidade
}

type Provider interface {
	Name() string
	Capabilities(ctx context.Context) (Capabilities, error)

	// Start levanta ou anexa os serviços solicitados e devolve endpoints prontos.
	Start(ctx context.Context, spec StartSpec) (map[string]Endpoint, error)

	// Health: verifying connectivity de cada serviço (ex: SELECT 1, PING, AMQP open/close)
	Health(ctx context.Context, eps map[string]Endpoint) error

	// Stop (opcional nos managed)
	Stop(ctx context.Context, refs []ServiceRef) error
}
