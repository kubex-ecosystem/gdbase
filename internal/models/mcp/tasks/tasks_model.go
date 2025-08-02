// Package tasks provides the model for managing tasks in the MCP (Microservices Control Plane).
package tasks

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	tp "github.com/rafa-mori/gdbase/types"
)

// TaskStatus represents the possible states of a task
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusRunning   TaskStatus = "RUNNING"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	TaskStatusFailed    TaskStatus = "FAILED"
	TaskStatusCancelled TaskStatus = "CANCELLED"
)

// TaskType represents the type of synchronization task
type TaskType string

const (
	TaskTypePull TaskType = "pull"
	TaskTypePush TaskType = "push"
	TaskTypeSync TaskType = "sync"
)

// JobScheduleType represents the scheduling type for tasks
type JobScheduleType string

const (
	JobScheduleTypeCron   JobScheduleType = "cron"
	JobScheduleTypeManual JobScheduleType = "manual"
	JobScheduleTypeEvent  JobScheduleType = "event"
)

// HTTPMethod represents supported HTTP methods
type HTTPMethod string

const (
	HTTPMethodGET    HTTPMethod = "GET"
	HTTPMethodPOST   HTTPMethod = "POST"
	HTTPMethodPUT    HTTPMethod = "PUT"
	HTTPMethodDELETE HTTPMethod = "DELETE"
	HTTPMethodPATCH  HTTPMethod = "PATCH"
)

// ITasksModel defines the interface for MCP Tasks
type ITasksModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetMCPProvider() string
	SetMCPProvider(provider string)
	GetTargetTask() string
	SetTargetTask(target string)
	GetTaskType() TaskType
	SetTaskType(taskType TaskType)
	GetTaskExpression() string
	SetTaskExpression(expression string)
	GetTaskAPIEndpoint() string
	SetTaskAPIEndpoint(endpoint string)
	GetTaskMethod() HTTPMethod
	SetTaskMethod(method HTTPMethod)
	GetTaskPayload() tp.JsonB
	SetTaskPayload(payload tp.JsonB)
	GetTaskHeaders() tp.JsonB
	SetTaskHeaders(headers tp.JsonB)
	GetTaskConfig() tp.JsonB
	SetTaskConfig(config tp.JsonB)
	GetActive() bool
	SetActive(active bool)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetCreatedBy() string
	SetCreatedBy(createdBy string)
	GetUpdatedBy() string
	SetUpdatedBy(updatedBy string)
	Validate() error
	Sanitize()
	// Integration methods with existing CronJob system
	ToCronJob() (*CronJobIntegration, error)
	FromCronJob(cronJob *CronJobIntegration) error
}

// TasksModel represents the MCP sync tasks table
type TasksModel struct {
	ID                 string          `gorm:"type:uuid;primaryKey" json:"id"`
	MCPProvider        string          `gorm:"type:text;not null" json:"mcp_provider" example:"github"`
	TargetTask         string          `gorm:"type:text;not null" json:"target_task" example:"my-repo"`
	LastSynced         *time.Time      `gorm:"type:timestamp;default:now()" json:"last_synced,omitempty"`
	CreatedAt          string          `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy          string          `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty"`
	UpdatedAt          string          `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedBy          string          `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty"`
	TaskType           TaskType        `gorm:"type:text;not null;check:task_type IN ('pull','push','sync')" json:"task_type" example:"sync"`
	TaskSchedule       JobScheduleType `gorm:"type:text;default:'cron'" json:"task_schedule" example:"cron"`
	TaskExpression     string          `gorm:"type:text;default:'2 * * * *'" json:"task_expression" example:"0 */6 * * *"`
	TaskCommandType    string          `gorm:"type:text;default:'api'" json:"task_command_type" example:"api"`
	TaskMethod         HTTPMethod      `gorm:"type:text;default:'POST'" json:"task_method" example:"POST"`
	TaskAPIEndpoint    string          `gorm:"type:text" json:"task_api_endpoint,omitempty" example:"/api/v1/sync"`
	TaskPayload        tp.JsonB        `json:"task_payload" binding:"omitempty"`
	TaskHeaders        tp.JsonB        `json:"task_headers" binding:"omitempty"`
	TaskRetries        int             `gorm:"type:integer;default:0" json:"task_retries" example:"3"`
	TaskTimeout        int             `gorm:"type:integer;default:0" json:"task_timeout" example:"300"`
	TaskStatus         TaskStatus      `gorm:"type:text;default:'PENDING'" json:"task_status" example:"PENDING"`
	TaskNextRun        *time.Time      `gorm:"type:timestamp" json:"task_next_run,omitempty"`
	TaskLastRun        *time.Time      `gorm:"type:timestamp" json:"task_last_run,omitempty"`
	TaskLastRunStatus  string          `gorm:"type:text;default:'pending'" json:"task_last_run_status" example:"success"`
	TaskLastRunMessage string          `gorm:"type:text;default:'pending'" json:"task_last_run_message" example:"Task completed successfully"`
	TaskCommand        string          `gorm:"type:text" json:"task_command,omitempty" example:"curl -X POST ..."`
	TaskActivated      bool            `gorm:"type:boolean;default:true" json:"task_activated" example:"true"`
	TaskConfig         tp.JsonB        `json:"task_config" binding:"omitempty"`
	TaskTags           []string        `gorm:"type:text[]" json:"task_tags,omitempty" example:"[\"sync\",\"github\"]"`
	TaskPriority       int             `gorm:"type:integer;default:0" json:"task_priority" example:"1"`
	TaskNotes          string          `gorm:"type:text" json:"task_notes,omitempty" example:"Daily sync with GitHub"`
	TaskCreatedAt      string          `gorm:"type:timestamp;default:now()" json:"task_created_at,omitempty"`
	TaskUpdatedAt      string          `gorm:"type:timestamp;default:now()" json:"task_updated_at,omitempty"`
	TaskCreatedBy      string          `gorm:"type:uuid;references:users(id)" json:"task_created_by,omitempty"`
	TaskUpdatedBy      string          `gorm:"type:uuid;references:users(id)" json:"task_updated_by,omitempty"`
	TaskLastExecutedBy string          `gorm:"type:uuid;references:users(id)" json:"task_last_executed_by,omitempty"`
	TaskLastExecutedAt *time.Time      `gorm:"type:timestamp" json:"task_last_executed_at,omitempty"`
	Config             tp.JsonB        `json:"config" binding:"omitempty"`
	Active             bool            `gorm:"type:boolean;default:true" json:"active" example:"true"`
}

// CronJobIntegration represents the structure for integrating with existing CronJob system
type CronJobIntegration struct {
	ID             uuid.UUID `json:"id"`
	CronExpression string    `json:"cron_expression"`
	Command        string    `json:"command"`
	Method         string    `json:"method"`
	APIEndpoint    string    `json:"api_endpoint"`
	Payload        string    `json:"payload"`
	Headers        string    `json:"headers"`
	Retries        int       `json:"retries"`
	Timeout        int       `json:"timeout"`
	UserID         uuid.UUID `json:"user_id"`
	IsActive       bool      `json:"is_active"`
}

func NewTasksModel() *TasksModel {
	return &TasksModel{
		ID:                 "",
		MCPProvider:        "",
		TargetTask:         "",
		TaskType:           TaskTypeSync,
		TaskSchedule:       JobScheduleTypeCron,
		TaskExpression:     "2 * * * *",
		TaskCommandType:    "api",
		TaskMethod:         HTTPMethodPOST,
		TaskRetries:        0,
		TaskTimeout:        0,
		TaskStatus:         TaskStatusPending,
		TaskActivated:      true,
		TaskPriority:       0,
		TaskPayload:        tp.JsonB{},
		TaskHeaders:        tp.JsonB{},
		TaskConfig:         tp.JsonB{},
		Config:             tp.JsonB{},
		Active:             true,
		CreatedAt:          time.Now().Format(time.RFC3339),
		UpdatedAt:          time.Now().Format(time.RFC3339),
		TaskCreatedAt:      time.Now().Format(time.RFC3339),
		TaskUpdatedAt:      time.Now().Format(time.RFC3339),
		TaskLastRunStatus:  "pending",
		TaskLastRunMessage: "pending",
	}
}

func (t *TasksModel) TableName() string                   { return "mcp_sync_tasks" }
func (t *TasksModel) GetID() string                       { return t.ID }
func (t *TasksModel) SetID(id string)                     { t.ID = id }
func (t *TasksModel) GetMCPProvider() string              { return t.MCPProvider }
func (t *TasksModel) SetMCPProvider(provider string)      { t.MCPProvider = provider }
func (t *TasksModel) GetTargetTask() string               { return t.TargetTask }
func (t *TasksModel) SetTargetTask(target string)         { t.TargetTask = target }
func (t *TasksModel) GetTaskType() TaskType               { return t.TaskType }
func (t *TasksModel) SetTaskType(taskType TaskType)       { t.TaskType = taskType }
func (t *TasksModel) GetTaskExpression() string           { return t.TaskExpression }
func (t *TasksModel) SetTaskExpression(expression string) { t.TaskExpression = expression }
func (t *TasksModel) GetTaskAPIEndpoint() string          { return t.TaskAPIEndpoint }
func (t *TasksModel) SetTaskAPIEndpoint(endpoint string)  { t.TaskAPIEndpoint = endpoint }
func (t *TasksModel) GetTaskMethod() HTTPMethod           { return t.TaskMethod }
func (t *TasksModel) SetTaskMethod(method HTTPMethod)     { t.TaskMethod = method }
func (t *TasksModel) GetTaskPayload() tp.JsonB            { return t.TaskPayload }
func (t *TasksModel) SetTaskPayload(payload tp.JsonB)     { t.TaskPayload = payload }
func (t *TasksModel) GetTaskHeaders() tp.JsonB            { return t.TaskHeaders }
func (t *TasksModel) SetTaskHeaders(headers tp.JsonB)     { t.TaskHeaders = headers }
func (t *TasksModel) GetTaskConfig() tp.JsonB             { return t.TaskConfig }
func (t *TasksModel) SetTaskConfig(config tp.JsonB)       { t.TaskConfig = config }
func (t *TasksModel) GetActive() bool                     { return t.Active }
func (t *TasksModel) SetActive(active bool)               { t.Active = active }
func (t *TasksModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, t.CreatedAt)
	return createdAt
}
func (t *TasksModel) SetCreatedAt(createdAt time.Time) { t.CreatedAt = createdAt.Format(time.RFC3339) }
func (t *TasksModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, t.UpdatedAt)
	return updatedAt
}
func (t *TasksModel) SetUpdatedAt(updatedAt time.Time) { t.UpdatedAt = updatedAt.Format(time.RFC3339) }
func (t *TasksModel) GetCreatedBy() string             { return t.CreatedBy }
func (t *TasksModel) SetCreatedBy(createdBy string)    { t.CreatedBy = createdBy }
func (t *TasksModel) GetUpdatedBy() string             { return t.UpdatedBy }
func (t *TasksModel) SetUpdatedBy(updatedBy string)    { t.UpdatedBy = updatedBy }

func (t *TasksModel) Validate() error {
	if t.MCPProvider == "" {
		return fmt.Errorf("MCP provider cannot be empty")
	}
	if t.TargetTask == "" {
		return fmt.Errorf("target task cannot be empty")
	}
	if t.TaskType == "" {
		return fmt.Errorf("task type cannot be empty")
	}
	if t.TaskExpression == "" {
		return fmt.Errorf("task expression cannot be empty")
	}
	return nil
}

func (t *TasksModel) Sanitize() {
	t.UpdatedAt = time.Now().Format(time.RFC3339)
	t.TaskUpdatedAt = time.Now().Format(time.RFC3339)
}

// ToCronJob converts a TasksModel to CronJobIntegration for integration with existing CronJob system
func (t *TasksModel) ToCronJob() (*CronJobIntegration, error) {
	if err := t.Validate(); err != nil {
		return nil, fmt.Errorf("task validation failed: %w", err)
	}

	// Parse user ID if present
	var userID uuid.UUID
	if t.CreatedBy != "" {
		parsedID, err := uuid.Parse(t.CreatedBy)
		if err != nil {
			return nil, fmt.Errorf("invalid user ID: %w", err)
		}
		userID = parsedID
	}

	// Convert JSONB to string for CronJob compatibility
	payloadStr := "{}"
	if t.TaskPayload != nil {
		if payloadBytes, err := json.Marshal(t.TaskPayload); err == nil {
			payloadStr = string(payloadBytes)
		}
	}

	headersStr := "{}"
	if t.TaskHeaders != nil {
		if headersBytes, err := json.Marshal(t.TaskHeaders); err == nil {
			headersStr = string(headersBytes)
		}
	}

	// Build command based on task type and configuration
	command := t.TaskCommand
	if command == "" && t.TaskAPIEndpoint != "" {
		command = fmt.Sprintf("MCP_TASK:%s:%s:%s", t.MCPProvider, t.TargetTask, string(t.TaskType))
	}

	return &CronJobIntegration{
		ID:             uuid.New(),
		CronExpression: t.TaskExpression,
		Command:        command,
		Method:         string(t.TaskMethod),
		APIEndpoint:    t.TaskAPIEndpoint,
		Payload:        payloadStr,
		Headers:        headersStr,
		Retries:        t.TaskRetries,
		Timeout:        t.TaskTimeout,
		UserID:         userID,
		IsActive:       t.Active && t.TaskActivated,
	}, nil
}

// FromCronJob updates a TasksModel from CronJobIntegration
func (t *TasksModel) FromCronJob(cronJob *CronJobIntegration) error {
	if cronJob == nil {
		return fmt.Errorf("cronJob cannot be nil")
	}

	t.TaskExpression = cronJob.CronExpression
	t.TaskCommand = cronJob.Command
	t.TaskMethod = HTTPMethod(cronJob.Method)
	t.TaskAPIEndpoint = cronJob.APIEndpoint
	t.TaskRetries = cronJob.Retries
	t.TaskTimeout = cronJob.Timeout
	t.Active = cronJob.IsActive
	t.CreatedBy = cronJob.UserID.String()

	// Parse JSONB fields
	if cronJob.Payload != "" && cronJob.Payload != "{}" {
		var payload tp.JsonB
		if err := json.Unmarshal([]byte(cronJob.Payload), &payload); err == nil {
			t.TaskPayload = payload
		}
	}

	if cronJob.Headers != "" && cronJob.Headers != "{}" {
		var headers tp.JsonB
		if err := json.Unmarshal([]byte(cronJob.Headers), &headers); err == nil {
			t.TaskHeaders = headers
		}
	}

	return nil
}
