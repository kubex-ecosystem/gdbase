package factory

import (
	dkrs "github.com/rafa-mori/gdbase/internal/services"
	t "github.com/rafa-mori/gdbase/types"
	l "github.com/rafa-mori/logz"
)

type DockerSrv interface {
	dkrs.IDockerService
}

func NewDockerService(config *t.DBConfig, logger l.Logger) (DockerSrv, error) {
	return dkrs.NewDockerService(config, logger)
}
