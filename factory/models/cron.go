package models

import (
	"context"
	"errors"

	m "github.com/rafa-mori/gdbase/internal/models/cron"
	"gorm.io/gorm"
)

type CronJobType = m.CronJob
type CronJobModel = m.CronJob
type CronJobService = m.ICronJobService
type CronJobRepo = m.ICronJobRepo

func NewCronJobService(cronJobRepo CronJobRepo) CronJobService {
	return m.NewCronJobService(cronJobRepo)
}
func NewCronJobRepo(ctx context.Context, db *gorm.DB) CronJobRepo {
	return m.NewCronJobRepo(ctx, db)
}
func NewCronJob(ctx context.Context, cron *CronJobModel, restrict bool) (*CronJobModel, error) {
	if cn, ok := m.NewCronJob(ctx, cron, restrict).(*CronJobModel); ok {
		return cn, nil
	}
	return nil, errors.New("failed to create cron job")
}
