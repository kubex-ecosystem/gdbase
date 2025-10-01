// Package dockerstack provides a local Docker-based stack implementation
package dockerstack

// Provider is an alias to the real implementation in adapter.go
// This file exists for backward compatibility
type Provider = DockerStackProvider

// New creates a new dockerstack provider instance
func New() *Provider {
	return NewDockerStackProvider()
}
