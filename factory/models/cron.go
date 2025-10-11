package models

import (
	"context"
	"errors"

	m "github.com/kubex-ecosystem/gdbase/internal/models/cron"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type CronJobType = m.CronJob
type CronJobModel = m.CronJob
type CronJobService = m.ICronJobService
type CronJobServiceImpl = m.CronJobService

type CronJobRepoImpl = m.CronJobRepo
type CronJobRepo = m.ICronJobRepo

func NewCronJobServiceImpl(cronJobRepo *CronJobRepoImpl) *CronJobServiceImpl {
	return m.NewCronJobServiceImpl(cronJobRepo)
}

func NewCronJobService(cronJobRepo *CronJobRepoImpl) CronJobService {
	return m.NewCronJobService(cronJobRepo)
}

func NewCronJobRepoImpl(ctx context.Context, dbService *svc.DBServiceImpl) *CronJobRepoImpl {
	return m.NewCronJobRepoImpl(ctx, dbService)
}

func NewCronJobRepo(ctx context.Context, dbService *svc.DBServiceImpl) CronJobRepo {
	return m.NewCronJobRepoImpl(ctx, dbService)
}

func NewCronJob(ctx context.Context, cron *CronJobModel, restrict bool) (*CronJobModel, error) {
	if cn, ok := m.NewCronJob(ctx, cron, restrict).(*CronJobModel); ok {
		return cn, nil
	}
	return nil, errors.New("failed to create cron job")
}
