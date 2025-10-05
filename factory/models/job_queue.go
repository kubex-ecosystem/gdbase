package models

import (
	"context"

	m "github.com/kubex-ecosystem/gdbase/internal/models/job_queue"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type JobQueue = m.JobQueue
type JobQueueModel = m.IJobQueue
type JobQueueService = m.IJobQueueService
type JobQueueRepo = m.IJobQueueRepo

func NewJobQueueService(jobQueueRepo JobQueueRepo) JobQueueService {
	return m.NewJobQueueService(jobQueueRepo)
}

func NewJobQueueRepo(ctx context.Context, dbService *svc.DBServiceImpl) JobQueueRepo {
	return m.NewJobQueueRepository(ctx, dbService)
}

func NewJobQueueModel() JobQueueModel {
	return m.NewJobQueueModel()
}
