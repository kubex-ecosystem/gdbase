package jobqueue

import (
	"context"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type IJobQueue interface {
	TableName() string
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	GetCronJobID() uuid.UUID
	SetCronJobID(cronJobID uuid.UUID)
	GetStatus() string
	SetStatus(status string)
	GetScheduledAt() time.Time
	SetScheduledAt(scheduledAt time.Time)
	GetExecutionTime() time.Time
	SetExecutionTime(executionTime time.Time)
	GetErrorMessage() string
	SetErrorMessage(errorMessage string)
	GetRetryCount() int
	SetRetryCount(retryCount int)
	GetNextRunTime() time.Time
	SetNextRunTime(nextRunTime time.Time)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetMetadata() string
	SetMetadata(metadata string)
	GetUserID() uuid.UUID
	SetUserID(userID uuid.UUID)
	GetCreatedBy() uuid.UUID
	SetCreatedBy(createdBy uuid.UUID)
	GetUpdatedBy() uuid.UUID
	SetUpdatedBy(updatedBy uuid.UUID)
	GetLastExecutedBy() uuid.UUID
	SetLastExecutedBy(lastExecutedBy uuid.UUID)
	GetJobType() string
	SetJobType(jobType string)
	GetJobExpression() string
	SetJobExpression(jobExpression string)
	GetJobCommand() string
	SetJobCommand(jobCommand string)
	GetJobMethod() string
	SetJobMethod(jobMethod string)
	GetJobAPIEndpoint() string
	SetJobAPIEndpoint(jobAPIEndpoint string)
	GetJobPayload() t.JSONB
	SetJobPayload(jobPayload t.JSONB)
	GetJobHeaders() t.JSONB
	SetJobHeaders(jobHeaders t.JSONB)
	GetJobRetries() int
	SetJobRetries(jobRetries int)
	GetJobTimeout() int
	SetJobTimeout(jobTimeout int)
}

type JobQueue struct {
	ID             uuid.UUID `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Code           int       `json:"code" xml:"code" yaml:"code" gorm:"column:code;primaryKey;autoIncrement"`
	CronJobID      uuid.UUID `json:"cronjob_id" xml:"cronjob_id" yaml:"cronjob_id" gorm:"column:cronjob_id"`
	Status         string    `json:"status" xml:"status" yaml:"status" gorm:"column:status;default:'PENDING'"`
	ScheduledAt    time.Time `json:"scheduled_time" xml:"scheduled_time" yaml:"scheduled_time" gorm:"column:scheduled_time;default:now()"`
	ExecutionTime  time.Time `json:"execution_time" xml:"execution_time" yaml:"execution_time" gorm:"column:execution_time"`
	ErrorMessage   string    `json:"error_message" xml:"error_message" yaml:"error_message" gorm:"column:error_message;default:null"`
	RetryCount     int       `json:"retry_count" xml:"retry_count" yaml:"retry_count" gorm:"column:retry_count;default:0"`
	NextRunTime    time.Time `json:"next_run_time" xml:"next_run_time" yaml:"next_run_time" gorm:"column:next_run_time"`
	CreatedAt      time.Time `json:"created_at" xml:"created_at" yaml:"created_at" gorm:"column:created_at;default:now()"`
	UpdatedAt      time.Time `json:"updated_at" xml:"updated_at" yaml:"updated_at" gorm:"column:updated_at;default:now()"`
	Metadata       string    `json:"metadata" xml:"metadata" yaml:"metadata" gorm:"column:metadata;default:null"`
	UserID         uuid.UUID `json:"user_id" xml:"user_id" yaml:"user_id" gorm:"column:user_id;references:users(id)"`
	CreatedBy      uuid.UUID `json:"created_by" xml:"created_by" yaml:"created_by" gorm:"column:created_by;references:users(id)"`
	UpdatedBy      uuid.UUID `json:"updated_by" xml:"updated_by" yaml:"updated_by" gorm:"column:updated_by;references:users(id)"`
	LastExecutedBy uuid.UUID `json:"last_executed_by" xml:"last_executed_by" yaml:"last_executed_by" gorm:"column:last_executed_by;references:users(id)"`
	JobType        string    `json:"job_type" xml:"job_type" yaml:"job_type" gorm:"column:job_type;default:'cron'"`
	JobExpression  string    `json:"job_expression" xml:"job_expression" yaml:"job_expression" gorm:"column:job_expression;default:'2 * * * *'"`
	JobCommand     string    `json:"job_command" xml:"job_command" yaml:"job_command" gorm:"column:job_command"`
	JobMethod      string    `json:"job_method" xml:"job_method" yaml:"job_method" gorm:"column:job_method"`
	JobAPIEndpoint string    `json:"job_api_endpoint" xml:"job_api_endpoint" yaml:"job_api_endpoint" gorm:"column:job_api_endpoint"`
	JobPayload     t.JSONB   `json:"job_payload" xml:"job_payload" yaml:"job_payload" gorm:"column:job_payload"`
	JobHeaders     t.JSONB   `json:"job_headers" xml:"job_headers" yaml:"job_headers" gorm:"column:job_headers"`
	JobRetries     int       `json:"job_retries" xml:"job_retries" yaml:"job_retries" gorm:"column:job_retries;default:0"`
	JobTimeout     int       `json:"job_timeout" xml:"job_timeout" yaml:"job_timeout" gorm:"column:job_timeout;default:0"`
}

func NewJobQueueModel() IJobQueue {
	return &JobQueue{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (j *JobQueue) TableName() string                          { return "job_queue" }
func (j *JobQueue) GetID() uuid.UUID                           { return j.ID }
func (j *JobQueue) GetCode() int                               { return j.Code }
func (j *JobQueue) SetCode(code int)                           { j.Code = code }
func (j *JobQueue) SetID(id uuid.UUID)                         { j.ID = id }
func (j *JobQueue) GetCronJobID() uuid.UUID                    { return j.CronJobID }
func (j *JobQueue) SetCronJobID(cronJobID uuid.UUID)           { j.CronJobID = cronJobID }
func (j *JobQueue) GetStatus() string                          { return j.Status }
func (j *JobQueue) SetStatus(status string)                    { j.Status = status }
func (j *JobQueue) GetScheduledAt() time.Time                  { return j.ScheduledAt }
func (j *JobQueue) SetScheduledAt(scheduledAt time.Time)       { j.ScheduledAt = scheduledAt }
func (j *JobQueue) GetExecutionTime() time.Time                { return j.ExecutionTime }
func (j *JobQueue) SetExecutionTime(executionTime time.Time)   { j.ExecutionTime = executionTime }
func (j *JobQueue) GetErrorMessage() string                    { return j.ErrorMessage }
func (j *JobQueue) SetErrorMessage(errorMessage string)        { j.ErrorMessage = errorMessage }
func (j *JobQueue) GetRetryCount() int                         { return j.RetryCount }
func (j *JobQueue) SetRetryCount(retryCount int)               { j.RetryCount = retryCount }
func (j *JobQueue) GetNextRunTime() time.Time                  { return j.NextRunTime }
func (j *JobQueue) SetNextRunTime(nextRunTime time.Time)       { j.NextRunTime = nextRunTime }
func (j *JobQueue) GetCreatedAt() time.Time                    { return j.CreatedAt }
func (j *JobQueue) SetCreatedAt(createdAt time.Time)           { j.CreatedAt = createdAt }
func (j *JobQueue) GetUpdatedAt() time.Time                    { return j.UpdatedAt }
func (j *JobQueue) SetUpdatedAt(updatedAt time.Time)           { j.UpdatedAt = updatedAt }
func (j *JobQueue) GetMetadata() string                        { return j.Metadata }
func (j *JobQueue) SetMetadata(metadata string)                { j.Metadata = metadata }
func (j *JobQueue) GetUserID() uuid.UUID                       { return j.UserID }
func (j *JobQueue) SetUserID(userID uuid.UUID)                 { j.UserID = userID }
func (j *JobQueue) GetCreatedBy() uuid.UUID                    { return j.CreatedBy }
func (j *JobQueue) SetCreatedBy(createdBy uuid.UUID)           { j.CreatedBy = createdBy }
func (j *JobQueue) GetUpdatedBy() uuid.UUID                    { return j.UpdatedBy }
func (j *JobQueue) SetUpdatedBy(updatedBy uuid.UUID)           { j.UpdatedBy = updatedBy }
func (j *JobQueue) GetLastExecutedBy() uuid.UUID               { return j.LastExecutedBy }
func (j *JobQueue) SetLastExecutedBy(lastExecutedBy uuid.UUID) { j.LastExecutedBy = lastExecutedBy }
func (j *JobQueue) GetJobType() string                         { return j.JobType }
func (j *JobQueue) SetJobType(jobType string)                  { j.JobType = jobType }
func (j *JobQueue) GetJobExpression() string                   { return j.JobExpression }
func (j *JobQueue) SetJobExpression(jobExpression string)      { j.JobExpression = jobExpression }
func (j *JobQueue) GetJobCommand() string                      { return j.JobCommand }
func (j *JobQueue) SetJobCommand(jobCommand string)            { j.JobCommand = jobCommand }
func (j *JobQueue) GetJobMethod() string                       { return j.JobMethod }
func (j *JobQueue) SetJobMethod(jobMethod string)              { j.JobMethod = jobMethod }
func (j *JobQueue) GetJobAPIEndpoint() string                  { return j.JobAPIEndpoint }
func (j *JobQueue) SetJobAPIEndpoint(jobAPIEndpoint string)    { j.JobAPIEndpoint = jobAPIEndpoint }
func (j *JobQueue) GetJobPayload() t.JSONB                     { return j.JobPayload }
func (j *JobQueue) SetJobPayload(jobPayload t.JSONB)           { j.JobPayload = jobPayload }
func (j *JobQueue) GetJobHeaders() t.JSONB                     { return j.JobHeaders }
func (j *JobQueue) SetJobHeaders(jobHeaders t.JSONB)           { j.JobHeaders = jobHeaders }
func (j *JobQueue) GetJobRetries() int                         { return j.JobRetries }
func (j *JobQueue) SetJobRetries(jobRetries int)               { j.JobRetries = jobRetries }
func (j *JobQueue) GetJobTimeout() int                         { return j.JobTimeout }
func (j *JobQueue) SetJobTimeout(jobTimeout int)               { j.JobTimeout = jobTimeout }

type IExecutionLog interface {
	TableName() string
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	GetCronJobID() uuid.UUID
	SetCronJobID(cronJobID uuid.UUID)
	GetExecutionTime() time.Time
	SetExecutionTime(executionTime time.Time)
	GetStatus() uuid.UUID
	SetStatus(status uuid.UUID)
	GetOutput() uuid.UUID
	SetOutput(output uuid.UUID)
	GetErrorMessage() uuid.UUID
	SetErrorMessage(errorMessage uuid.UUID)
	GetRetryCount() int
	SetRetryCount(retryCount int)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetUserID() uuid.UUID
	SetUserID(userID uuid.UUID)
	GetCreatedBy() uuid.UUID
	SetCreatedBy(createdBy uuid.UUID)
	GetUpdatedBy() uuid.UUID
	SetUpdatedBy(updatedBy uuid.UUID)
	GetMetadata() string
	SetMetadata(metadata string)
}

type ExecutionLog struct {
	ID            uuid.UUID `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	CronJobID     uuid.UUID `json:"cronjob_id" xml:"cronjob_id" yaml:"cronjob_id" gorm:"column:cronjob_id;type:uuid;references:CronJob(id)"`
	ExecutionTime time.Time `json:"execution_time" xml:"execution_time" yaml:"execution_time" gorm:"column:execution_time;default:now()"`
	Status        string    `json:"status" xml:"status" yaml:"status" gorm:"column:status;default:'PENDING'"`
	Output        string    `json:"output" xml:"output" yaml:"output" gorm:"column:output;default:null"`
	ErrorMessage  string    `json:"error_message" xml:"error_message" yaml:"error_message" gorm:"column:error_message;default:null"`
	RetryCount    int       `json:"retry_count" xml:"retry_count" yaml:"retry_count" gorm:"column:retry_count;default:0"`
	CreatedAt     time.Time `json:"created_at" xml:"created_at" yaml:"created_at" gorm:"column:created_at;default:now()"`
	UpdatedAt     time.Time `json:"updated_at" xml:"updated_at" yaml:"updated_at" gorm:"column:updated_at;default:now()"`
	UserID        uuid.UUID `json:"user_id" xml:"user_id" yaml:"user_id" gorm:"column:user_id;type:uuid;references:users(id)"`
	CreatedBy     uuid.UUID `json:"created_by" xml:"created_by" yaml:"created_by" gorm:"column:created_by;type:uuid;references:users(id)"`
	UpdatedBy     uuid.UUID `json:"updated_by" xml:"updated_by" yaml:"updated_by" gorm:"column:updated_by;type:uuid;references:users(id)"`
	Metadata      string    `json:"metadata" xml:"metadata" yaml:"metadata" gorm:"column:metadata;type:jsonb"`
}

func (e *ExecutionLog) TableName() string                        { return "execution_logs" }
func (e *ExecutionLog) GetID() uuid.UUID                         { return e.ID }
func (e *ExecutionLog) SetID(id uuid.UUID)                       { e.ID = id }
func (e *ExecutionLog) GetCronJobID() uuid.UUID                  { return e.CronJobID }
func (e *ExecutionLog) SetCronJobID(cronJobID uuid.UUID)         { e.CronJobID = cronJobID }
func (e *ExecutionLog) GetExecutionTime() time.Time              { return e.ExecutionTime }
func (e *ExecutionLog) SetExecutionTime(executionTime time.Time) { e.ExecutionTime = executionTime }
func (e *ExecutionLog) GetStatus() string                        { return e.Status }
func (e *ExecutionLog) SetStatus(status string)                  { e.Status = status }
func (e *ExecutionLog) GetOutput() string                        { return e.Output }
func (e *ExecutionLog) SetOutput(output string)                  { e.Output = output }
func (e *ExecutionLog) GetErrorMessage() string                  { return e.ErrorMessage }
func (e *ExecutionLog) SetErrorMessage(errorMessage string)      { e.ErrorMessage = errorMessage }
func (e *ExecutionLog) GetRetryCount() int                       { return e.RetryCount }
func (e *ExecutionLog) SetRetryCount(retryCount int)             { e.RetryCount = retryCount }
func (e *ExecutionLog) GetCreatedAt() time.Time                  { return e.CreatedAt }
func (e *ExecutionLog) SetCreatedAt(createdAt time.Time)         { e.CreatedAt = createdAt }
func (e *ExecutionLog) GetUpdatedAt() time.Time                  { return e.UpdatedAt }
func (e *ExecutionLog) SetUpdatedAt(updatedAt time.Time)         { e.UpdatedAt = updatedAt }
func (e *ExecutionLog) GetUserID() uuid.UUID                     { return e.UserID }
func (e *ExecutionLog) SetUserID(userID uuid.UUID)               { e.UserID = userID }
func (e *ExecutionLog) GetCreatedBy() uuid.UUID                  { return e.CreatedBy }
func (e *ExecutionLog) SetCreatedBy(createdBy uuid.UUID)         { e.CreatedBy = createdBy }
func (e *ExecutionLog) GetUpdatedBy() uuid.UUID                  { return e.UpdatedBy }
func (e *ExecutionLog) SetUpdatedBy(updatedBy uuid.UUID)         { e.UpdatedBy = updatedBy }
func (e *ExecutionLog) GetMetadata() string                      { return e.Metadata }
func (e *ExecutionLog) SetMetadata(metadata string)              { e.Metadata = metadata }

// IExecutionLogService defines the interface for execution log services.
type IExecutionLogService interface {
	CreateLog(ctx context.Context, log ExecutionLog) error
}
