package analysisjobs

import (
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type IAnalysisJob interface {
	TableName() string
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	GetProjectID() uuid.UUID
	SetProjectID(projectID uuid.UUID)
	GetJobType() string
	SetJobType(jobType string)
	GetStatus() string
	SetStatus(status string)
	GetSourceURL() string
	SetSourceURL(sourceURL string)
	GetSourceType() string
	SetSourceType(sourceType string)
	GetInputData() t.JSONBImpl
	SetInputData(inputData t.JSONBImpl)
	GetOutputData() t.JSONBImpl
	SetOutputData(outputData t.JSONBImpl)
	GetErrorMessage() string
	SetErrorMessage(errorMessage string)
	GetProgress() float64
	SetProgress(progress float64)
	GetStartedAt() time.Time
	SetStartedAt(startedAt time.Time)
	GetCompletedAt() time.Time
	SetCompletedAt(completedAt time.Time)
	GetRetryCount() int
	SetRetryCount(retryCount int)
	GetMaxRetries() int
	SetMaxRetries(maxRetries int)
	GetMetadata() t.JSONBImpl
	SetMetadata(metadata t.JSONBImpl)
	GetUserID() uuid.UUID
	SetUserID(userID uuid.UUID)
	GetCreatedBy() uuid.UUID
	SetCreatedBy(createdBy uuid.UUID)
	GetUpdatedBy() uuid.UUID
	SetUpdatedBy(updatedBy uuid.UUID)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
}

type AnalysisJob struct {
	ID           uuid.UUID   `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProjectID    uuid.UUID   `json:"project_id" xml:"project_id" yaml:"project_id" gorm:"column:project_id;type:uuid"`
	JobType      string      `json:"job_type" xml:"job_type" yaml:"job_type" gorm:"column:job_type;not null;type:VARCHAR(50)"`
	Status       string      `json:"status" xml:"status" yaml:"status" gorm:"column:status;not null;default:'PENDING';type:analysis_job_status"`
	SourceURL    string      `json:"source_url" xml:"source_url" yaml:"source_url" gorm:"column:source_url;type:TEXT"`
	SourceType   string      `json:"source_type" xml:"source_type" yaml:"source_type" gorm:"column:source_type;type:VARCHAR(50)"`
	InputData    t.JSONBImpl `json:"input_data" xml:"input_data" yaml:"input_data" gorm:"column:input_data;type:jsonb"`
	OutputData   t.JSONBImpl `json:"output_data" xml:"output_data" yaml:"output_data" gorm:"column:output_data;type:jsonb"`
	ErrorMessage *string     `json:"error_message" xml:"error_message" yaml:"error_message" gorm:"column:error_message;type:TEXT"`
	Progress     float64     `json:"progress" xml:"progress" yaml:"progress" gorm:"column:progress;default:0.0;type:DECIMAL(5,2)"`
	StartedAt    time.Time   `json:"started_at" xml:"started_at" yaml:"started_at" gorm:"column:started_at;type:timestamp"`
	CompletedAt  *time.Time  `json:"completed_at" xml:"completed_at" yaml:"completed_at" gorm:"column:completed_at;type:timestamp"`
	RetryCount   int         `json:"retry_count" xml:"retry_count" yaml:"retry_count" gorm:"column:retry_count;default:0"`
	MaxRetries   int         `json:"max_retries" xml:"max_retries" yaml:"max_retries" gorm:"column:max_retries;default:3"`
	Metadata     t.JSONBImpl `json:"metadata" xml:"metadata" yaml:"metadata" gorm:"column:metadata;type:jsonb"`
	UserID       uuid.UUID   `json:"user_id" xml:"user_id" yaml:"user_id" gorm:"column:user_id;type:uuid"`
	CreatedBy    uuid.UUID   `json:"created_by" xml:"created_by" yaml:"created_by" gorm:"column:created_by;type:uuid"`
	UpdatedBy    *uuid.UUID  `json:"updated_by" xml:"updated_by" yaml:"updated_by" gorm:"column:updated_by;type:uuid"`
	CreatedAt    time.Time   `json:"created_at" xml:"created_at" yaml:"created_at" gorm:"column:created_at;default:now()"`
	UpdatedAt    time.Time   `json:"updated_at" xml:"updated_at" yaml:"updated_at" gorm:"column:updated_at;default:now()"`
}

func NewAnalysisJobModel() IAnalysisJob {
	return &AnalysisJob{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *AnalysisJob) TableName() string                    { return "mcp_analysis_jobs" }
func (a *AnalysisJob) GetID() uuid.UUID                     { return a.ID }
func (a *AnalysisJob) SetID(id uuid.UUID)                   { a.ID = id }
func (a *AnalysisJob) GetProjectID() uuid.UUID              { return a.ProjectID }
func (a *AnalysisJob) SetProjectID(projectID uuid.UUID)     { a.ProjectID = projectID }
func (a *AnalysisJob) GetJobType() string                   { return a.JobType }
func (a *AnalysisJob) SetJobType(jobType string)            { a.JobType = jobType }
func (a *AnalysisJob) GetStatus() string                    { return a.Status }
func (a *AnalysisJob) SetStatus(status string)              { a.Status = status }
func (a *AnalysisJob) GetSourceURL() string                 { return a.SourceURL }
func (a *AnalysisJob) SetSourceURL(sourceURL string)        { a.SourceURL = sourceURL }
func (a *AnalysisJob) GetSourceType() string                { return a.SourceType }
func (a *AnalysisJob) SetSourceType(sourceType string)      { a.SourceType = sourceType }
func (a *AnalysisJob) GetInputData() t.JSONBImpl            { return a.InputData }
func (a *AnalysisJob) SetInputData(inputData t.JSONBImpl)   { a.InputData = inputData }
func (a *AnalysisJob) GetOutputData() t.JSONBImpl           { return a.OutputData }
func (a *AnalysisJob) SetOutputData(outputData t.JSONBImpl) { a.OutputData = outputData }
func (a *AnalysisJob) GetErrorMessage() string              { return *a.ErrorMessage }
func (a *AnalysisJob) SetErrorMessage(errorMessage string)  { a.ErrorMessage = &errorMessage }
func (a *AnalysisJob) GetProgress() float64                 { return a.Progress }
func (a *AnalysisJob) SetProgress(progress float64)         { a.Progress = progress }
func (a *AnalysisJob) GetStartedAt() time.Time              { return a.StartedAt }
func (a *AnalysisJob) SetStartedAt(startedAt time.Time)     { a.StartedAt = startedAt }
func (a *AnalysisJob) GetCompletedAt() time.Time            { return *a.CompletedAt }
func (a *AnalysisJob) SetCompletedAt(completedAt time.Time) { a.CompletedAt = &completedAt }
func (a *AnalysisJob) GetRetryCount() int                   { return a.RetryCount }
func (a *AnalysisJob) SetRetryCount(retryCount int)         { a.RetryCount = retryCount }
func (a *AnalysisJob) GetMaxRetries() int                   { return a.MaxRetries }
func (a *AnalysisJob) SetMaxRetries(maxRetries int)         { a.MaxRetries = maxRetries }
func (a *AnalysisJob) GetMetadata() t.JSONBImpl             { return a.Metadata }
func (a *AnalysisJob) SetMetadata(metadata t.JSONBImpl)     { a.Metadata = metadata }
func (a *AnalysisJob) GetUserID() uuid.UUID                 { return a.UserID }
func (a *AnalysisJob) SetUserID(userID uuid.UUID)           { a.UserID = userID }
func (a *AnalysisJob) GetCreatedBy() uuid.UUID              { return a.CreatedBy }
func (a *AnalysisJob) SetCreatedBy(createdBy uuid.UUID)     { a.CreatedBy = createdBy }
func (a *AnalysisJob) GetUpdatedBy() uuid.UUID              { return *a.UpdatedBy }
func (a *AnalysisJob) SetUpdatedBy(updatedBy uuid.UUID)     { a.UpdatedBy = &updatedBy }
func (a *AnalysisJob) GetCreatedAt() time.Time              { return a.CreatedAt }
func (a *AnalysisJob) SetCreatedAt(createdAt time.Time)     { a.CreatedAt = createdAt }
func (a *AnalysisJob) GetUpdatedAt() time.Time              { return a.UpdatedAt }
func (a *AnalysisJob) SetUpdatedAt(updatedAt time.Time)     { a.UpdatedAt = updatedAt }
