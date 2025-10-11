package factory

import (
	prv "github.com/kubex-ecosystem/gdbase/internal/provider"
)

type ServiceRef = prv.ServiceRef
type Endpoint = prv.Endpoint
type Capabilities = prv.Capabilities
type StartSpec = prv.StartSpec
type Provider = prv.Provider

const (
	EnginePostgres = prv.EnginePostgres
	EngineMongo    = prv.EngineMongo
	EngineRedis    = prv.EngineRedis
	EngineRabbit   = prv.EngineRabbit
)

type Engine = prv.Engine

func Register(p Provider)              { prv.Register(p) }
func Get(name string) (Provider, bool) { return prv.Get(name) }
func All() []Provider                  { return prv.All() }
