// Package bootstrap initializes and manages service providers based on configuration and environment variables.
package bootstrap

import (
	"context"
	"errors"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/kubex-ecosystem/gdbase/internal/provider"
)

type Config struct {
	Backends      []string // ordem de preferência; e.g. ["dockerstack"] por enquanto
	Strict        bool     // se true, não faz fallback silencioso
	Services      []provider.ServiceRef
	PreferredPort map[string]int
	Secrets       map[string]string
}

func FromEnv() Config {
	raw := os.Getenv("GDBASE_BACKENDS")
	if raw == "" {
		raw = "dockerstack"
	}
	backends := strings.Split(raw, ",")
	for i := range backends {
		backends[i] = strings.TrimSpace(backends[i])
	}
	strict := strings.EqualFold(os.Getenv("GDBASE_STRICT"), "true")
	return Config{
		Backends: backends, Strict: strict,
	}
}

type Result struct {
	Backend   string
	Endpoints map[string]provider.Endpoint
}

func Start(ctx context.Context, cfg Config) (Result, error) {
	// ordenar por prioridade “fixa” caso queira; hoje respeita ordem vinda
	cands := make([]string, 0, len(cfg.Backends))
	for _, b := range cfg.Backends {
		if _, ok := provider.Get(b); ok {
			cands = append(cands, b)
		}
	}
	if len(cands) == 0 {
		return Result{}, errors.New("no providers registered")
	}
	// opcional: ordenar alfabeticamente para determinismo quando empatar
	sort.Strings(cands)

	var lastErr error
	for _, name := range cands {
		p, _ := provider.Get(name)
		spec := provider.StartSpec{
			Services:      cfg.Services,
			PreferredPort: cfg.PreferredPort,
			Secrets:       cfg.Secrets,
			Labels:        map[string]string{"owner": "gdbase"},
		}
		eps, err := p.Start(ctx, spec)
		if err != nil {
			lastErr = err
			if cfg.Strict {
				return Result{}, err
			}
			continue
		}
		// health com backoff curto
		deadline := time.Now().Add(15 * time.Second)
		for {
			if err = p.Health(ctx, eps); err == nil || time.Now().After(deadline) {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		if err != nil {
			lastErr = err
			if cfg.Strict {
				return Result{}, err
			}
			continue
		}
		return Result{Backend: name, Endpoints: eps}, nil
	}
	return Result{}, lastErr
}
