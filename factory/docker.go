package factory

import (
	l "github.com/faelmori/logz"
	dkrs "github.com/rafa-mori/gdbase/internal/services"
	t "github.com/rafa-mori/gdbase/types"
)

type DockerSrv interface {
	dkrs.IDockerService
}

func NewDockerService(config *t.DBConfig, logger l.Logger) (DockerSrv, error) {
	return dkrs.NewDockerService(config, logger)
}
