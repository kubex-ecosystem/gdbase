// Package analysisjobs provides a repository interface and implementation for managing analysis jobs.
package analysisjobs

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type IAnalysisJobRepo interface {
	Create(ctx context.Context, job *AnalysisJob) (*AnalysisJob, error)
	FindByID(ctx context.Context, id uuid.UUID) (*AnalysisJob, error)
	FindAll(ctx context.Context) ([]*AnalysisJob, error)
	Update(ctx context.Context, job *AnalysisJob) (*AnalysisJob, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByStatus(ctx context.Context, status string) ([]*AnalysisJob, error)
	FindByJobType(ctx context.Context, jobType string) ([]*AnalysisJob, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*AnalysisJob, error)
	FindByProjectID(ctx context.Context, projectID uuid.UUID) ([]*AnalysisJob, error)
	FindByStatusAndJobType(ctx context.Context, status, jobType string) ([]*AnalysisJob, error)
	FindByUserIDAndStatus(ctx context.Context, userID uuid.UUID, status string) ([]*AnalysisJob, error)
	FindPendingJobs(ctx context.Context) ([]*AnalysisJob, error)
	FindRunningJobs(ctx context.Context) ([]*AnalysisJob, error)
	FindFailedJobs(ctx context.Context) ([]*AnalysisJob, error)
	UpdateProgress(ctx context.Context, id uuid.UUID, progress float64) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	MarkAsStarted(ctx context.Context, id uuid.UUID) error
	MarkAsCompleted(ctx context.Context, id uuid.UUID, outputData interface{}) error
	MarkAsFailed(ctx context.Context, id uuid.UUID, errorMessage string) error
	IncrementRetryCount(ctx context.Context, id uuid.UUID) error
}

type AnalysisJobRepository struct {
	db *gorm.DB
}

func NewAnalysisJobRepository(ctx context.Context, dbService *svc.DBServiceImpl) IAnalysisJobRepo {
	db, err := svc.GetDB(ctx, dbService)
	if err != nil {
		return nil
	}
	return &AnalysisJobRepository{db: db}
}

func (r *AnalysisJobRepository) Create(ctx context.Context, job *AnalysisJob) (*AnalysisJob, error) {
	if job == nil {
		return nil, errors.New("repository: job cannot be nil")
	}
	if err := r.db.WithContext(ctx).Create(job).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to create analysis job: %w", err)
	}
	return job, nil
}

func (r *AnalysisJobRepository) FindByID(ctx context.Context, id uuid.UUID) (*AnalysisJob, error) {
	if id == uuid.Nil {
		return nil, errors.New("repository: id cannot be empty")
	}
	var job AnalysisJob
	if err := r.db.WithContext(ctx).First(&job, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("repository: analysis job not found: %w", err)
		}
		return nil, fmt.Errorf("repository: failed to find analysis job: %w", err)
	}
	return &job, nil
}

func (r *AnalysisJobRepository) FindAll(ctx context.Context) ([]*AnalysisJob, error) {
	var jobs []*AnalysisJob
	if err := r.db.WithContext(ctx).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to find all analysis jobs: %w", err)
	}
	return jobs, nil
}

func (r *AnalysisJobRepository) Update(ctx context.Context, job *AnalysisJob) (*AnalysisJob, error) {
	if job == nil {
		return nil, errors.New("repository: job cannot be nil")
	}
	if job.ID == uuid.Nil {
		return nil, errors.New("repository: job ID cannot be empty")
	}
	if err := r.db.WithContext(ctx).Save(job).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to update analysis job: %w", err)
	}
	return job, nil
}

func (r *AnalysisJobRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("repository: id cannot be empty")
	}
	if err := r.db.WithContext(ctx).Delete(&AnalysisJob{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("repository: failed to delete analysis job: %w", err)
	}
	return nil
}

func (r *AnalysisJobRepository) FindByStatus(ctx context.Context, status string) ([]*AnalysisJob, error) {
	if status == "" {
		return nil, errors.New("repository: status cannot be empty")
	}
	var jobs []*AnalysisJob
	if err := r.db.WithContext(ctx).Where("status = ?", status).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to find jobs by status: %w", err)
	}
	return jobs, nil
}

func (r *AnalysisJobRepository) FindByJobType(ctx context.Context, jobType string) ([]*AnalysisJob, error) {
	if jobType == "" {
		return nil, errors.New("repository: job type cannot be empty")
	}
	var jobs []*AnalysisJob
	if err := r.db.WithContext(ctx).Where("job_type = ?", jobType).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to find jobs by type: %w", err)
	}
	return jobs, nil
}

func (r *AnalysisJobRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*AnalysisJob, error) {
	if userID == uuid.Nil {
		return nil, errors.New("repository: user ID cannot be empty")
	}
	var jobs []*AnalysisJob
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to find jobs by user ID: %w", err)
	}
	return jobs, nil
}

func (r *AnalysisJobRepository) FindByProjectID(ctx context.Context, projectID uuid.UUID) ([]*AnalysisJob, error) {
	if projectID == uuid.Nil {
		return nil, errors.New("repository: project ID cannot be empty")
	}
	var jobs []*AnalysisJob
	if err := r.db.WithContext(ctx).Where("project_id = ?", projectID).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to find jobs by project ID: %w", err)
	}
	return jobs, nil
}

func (r *AnalysisJobRepository) FindByStatusAndJobType(ctx context.Context, status, jobType string) ([]*AnalysisJob, error) {
	if status == "" {
		return nil, errors.New("repository: status cannot be empty")
	}
	if jobType == "" {
		return nil, errors.New("repository: job type cannot be empty")
	}
	var jobs []*AnalysisJob
	if err := r.db.WithContext(ctx).Where("status = ? AND job_type = ?", status, jobType).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to find jobs by status and type: %w", err)
	}
	return jobs, nil
}

func (r *AnalysisJobRepository) FindByUserIDAndStatus(ctx context.Context, userID uuid.UUID, status string) ([]*AnalysisJob, error) {
	if userID == uuid.Nil {
		return nil, errors.New("repository: user ID cannot be empty")
	}
	if status == "" {
		return nil, errors.New("repository: status cannot be empty")
	}
	var jobs []*AnalysisJob
	if err := r.db.WithContext(ctx).Where("user_id = ? AND status = ?", userID, status).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("repository: failed to find jobs by user ID and status: %w", err)
	}
	return jobs, nil
}

func (r *AnalysisJobRepository) FindPendingJobs(ctx context.Context) ([]*AnalysisJob, error) {
	return r.FindByStatus(ctx, "PENDING")
}

func (r *AnalysisJobRepository) FindRunningJobs(ctx context.Context) ([]*AnalysisJob, error) {
	return r.FindByStatus(ctx, "RUNNING")
}

func (r *AnalysisJobRepository) FindFailedJobs(ctx context.Context) ([]*AnalysisJob, error) {
	return r.FindByStatus(ctx, "FAILED")
}

func (r *AnalysisJobRepository) UpdateProgress(ctx context.Context, id uuid.UUID, progress float64) error {
	if id == uuid.Nil {
		return errors.New("repository: id cannot be empty")
	}
	if progress < 0 || progress > 100 {
		return errors.New("repository: progress must be between 0 and 100")
	}
	if err := r.db.WithContext(ctx).Model(&AnalysisJob{}).Where("id = ?", id).Update("progress", progress).Error; err != nil {
		return fmt.Errorf("repository: failed to update progress: %w", err)
	}
	return nil
}

func (r *AnalysisJobRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	if id == uuid.Nil {
		return errors.New("repository: id cannot be empty")
	}
	if status == "" {
		return errors.New("repository: status cannot be empty")
	}
	if err := r.db.WithContext(ctx).Model(&AnalysisJob{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return fmt.Errorf("repository: failed to update status: %w", err)
	}
	return nil
}

func (r *AnalysisJobRepository) MarkAsStarted(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("repository: id cannot be empty")
	}
	updates := map[string]interface{}{
		"status":     "RUNNING",
		"started_at": time.Now(),
	}
	if err := r.db.WithContext(ctx).Model(&AnalysisJob{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return fmt.Errorf("repository: failed to mark job as started: %w", err)
	}
	return nil
}

func (r *AnalysisJobRepository) MarkAsCompleted(ctx context.Context, id uuid.UUID, outputData interface{}) error {
	if id == uuid.Nil {
		return errors.New("repository: id cannot be empty")
	}
	updates := map[string]interface{}{
		"status":       "COMPLETED",
		"completed_at": time.Now(),
		"progress":     100.0,
	}
	if outputData != nil {
		updates["output_data"] = outputData
	}
	if err := r.db.WithContext(ctx).Model(&AnalysisJob{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return fmt.Errorf("repository: failed to mark job as completed: %w", err)
	}
	return nil
}

func (r *AnalysisJobRepository) MarkAsFailed(ctx context.Context, id uuid.UUID, errorMessage string) error {
	if id == uuid.Nil {
		return errors.New("repository: id cannot be empty")
	}
	updates := map[string]interface{}{
		"status":        "FAILED",
		"error_message": errorMessage,
	}
	if err := r.db.WithContext(ctx).Model(&AnalysisJob{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return fmt.Errorf("repository: failed to mark job as failed: %w", err)
	}
	return nil
}

func (r *AnalysisJobRepository) IncrementRetryCount(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("repository: id cannot be empty")
	}
	if err := r.db.WithContext(ctx).Model(&AnalysisJob{}).Where("id = ?", id).Update("retry_count", gorm.Expr("retry_count + 1")).Error; err != nil {
		return fmt.Errorf("repository: failed to increment retry count: %w", err)
	}
	return nil
}
