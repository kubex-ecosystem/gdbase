package models

import (
	m "github.com/kubex-ecosystem/gdbase/internal/models/job_queue"
	"gorm.io/gorm"
)

type JobQueueModel = m.IJobQueue
type JobQueueService = m.IJobQueueService
type JobQueueRepo = m.IJobQueueRepo

func NewJobQueueService(jobQueueRepo JobQueueRepo) JobQueueService {
	return m.NewJobQueueService(jobQueueRepo)
}

func NewJobQueueRepo(db *gorm.DB) JobQueueRepo {
	return m.NewJobQueueRepository(db)
}
