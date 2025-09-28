
go
package dockerstack

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/yourorg/gdbase/internal/gdbase/provider"
)

type Provider struct{}

func (p *Provider) Name() string { return "dockerstack" }

func (p *Provider) Capabilities(ctx context.Context) (provider.Capabilities, error) {
	return provider.Capabilities{
		Managed:  false,
		Notes:    []string{"Zero-config local stack (PG, Mongo, Redis, Rabbit)"},
		Features: map[string]bool{"network.internal": true, "publish.ports": true},
	}, nil
}

// Start aqui só delega pro SEU fluxo atual já funcional,
// retornando os DSNs/hosts/ports exatamente como hoje.
// Você pode chamar suas funções existentes e mapear para provider.Endpoint.
func (p *Provider) Start(ctx context.Context, spec provider.StartSpec) (map[string]provider.Endpoint, error) {
	// exemplo: chamar sua orquestração já existente
	// startDockerStack(spec) -> retorna mapa com host/port/user/pass
	// Abaixo, placeholder de forma explícita:
	endpoints := map[string]provider.Endpoint{}

	// Postgres
	if has(spec.Services, "pg") {
		user := "postgres"
		pass := spec.Secrets["pg_admin"]
		host := os.Getenv("GDBASE_PG_HOST")
		if host == "" { host = "127.0.0.1" }
		port := pick(spec.PreferredPort, "pg", 5432)
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/gdbase?sslmode=disable", user, pass, host, port)
		endpoints["pg"] = provider.Endpoint{DSN: dsn, Redacted: red(dsn, pass), Host: host, Port: port}
	}

	// Mongo
	if has(spec.Services, "mongo") {
		user := "root"
		pass := spec.Secrets["mongo_root"]
		host := "127.0.0.1"
		port := pick(spec.PreferredPort, "mongo", 27017)
		dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pass, host, port)
		endpoints["mongo"] = provider.Endpoint{DSN: dsn, Redacted: red(dsn, pass), Host: host, Port: port}
	}

	// Redis
	if has(spec.Services, "redis") {
		pass := spec.Secrets["redis_pass"]
		host := "127.0.0.1"
		port := pick(spec.PreferredPort, "redis", 6379)
		dsn := fmt.Sprintf("redis://:%s@%s:%d", pass, host, port)
		endpoints["redis"] = provider.Endpoint{DSN: dsn, Redacted: red(dsn, pass), Host: host, Port: port}
	}

	// Rabbit
	if has(spec.Services, "rabbit") {
		user := "admin"
		pass := spec.Secrets["rabbit_pass"]
		host := "127.0.0.1"
		port := pick(spec.PreferredPort, "rabbit", 5672)
		dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pass, host, port)
		endpoints["rabbit"] = provider.Endpoint{DSN: dsn, Redacted: red(dsn, pass), Host: host, Port: port}
	}

	// TODO: ligar isso às suas funções reais de start; aqui é mapeamento de retorno.
	_ = time.Now()
	return endpoints, nil
}

func (p *Provider) Health(ctx context.Context, eps map[string]provider.Endpoint) error {
	// Chamar seus healths já existentes (pg_isready; ping Redis; handshake AMQP; Mongo client ping)
	return nil
}
func (p *Provider) Stop(ctx context.Context, refs []provider.ServiceRef) error { return nil }

func has(services []provider.ServiceRef, name string) bool {
	for _, s := range services { if s.Name == name { return true } }
	return false
}
func pick(m map[string]int, key string, def int) int {
	if m == nil { return def }
	if v, ok := m[key]; ok && v > 0 { return v }
	return def
}
func red(dsn, secret string) string {
	if secret == "" { return dsn }
	return "***redacted***"
}