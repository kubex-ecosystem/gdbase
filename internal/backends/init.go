// Package backends registers all available backend providers
package backends

import (
	"github.com/kubex-ecosystem/gdbase/internal/backends/dockerstack"
	"github.com/kubex-ecosystem/gdbase/internal/provider"
)

func init() {
	// Register dockerstack provider (default)
	provider.Register(dockerstack.New())

	// TODO: Register other providers as they're implemented
	// provider.Register(localhost.New())
	// provider.Register(supabase.New())
	// provider.Register(aws.New())
}
