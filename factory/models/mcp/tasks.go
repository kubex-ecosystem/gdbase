package mcp

import (
	"time"

	"github.com/google/uuid"
	m "github.com/kubex-ecosystem/gdbase/internal/models/mcp/tasks"
	tp "github.com/kubex-ecosystem/gdbase/internal/types"
	"gorm.io/gorm"
)

type TaskSearchOptions = m.TaskSearchOptions
type JobScheduleType = m.JobScheduleType
type CronJobIntegration = m.CronJobIntegration
type HTTPMethod = m.HTTPMethod
type TaskType = m.TaskType
type TaskStatus = m.TaskStatus
type TasksModelType = m.TasksModel
type TasksModel = m.ITasksModel
type TasksService = m.ITasksService
type TasksRepo = m.ITasksRepo

func NewTasksService(tasksRepo TasksRepo) TasksService {
	return m.NewTasksService(tasksRepo)
}

func NewTasksRepo(db *gorm.DB) TasksRepo {
	return m.NewTasksRepo(db)
}

func NewTasksModel(
	provider string,
	target string,
	taskType TaskType,
	taskSchedule JobScheduleType,
	taskExpression string,
	taskCommandType string,
	taskMethod HTTPMethod,
	taskAPIEndpoint string,
	taskPayload tp.JSONB,
	taskHeaders tp.JSONB,
	taskRetries int,
	taskTimeout int,
	taskStatus TaskStatus,
	taskNextRun *time.Time,
	taskLastRun *time.Time,
	taskLastRunStatus string,
	taskLastRunMessage string,
	taskCommand string,
	taskActivated bool,
	taskConfig tp.JSONB,
	taskTags []string,
	taskPriority int,
	taskNotes string,
	taskCreatedAt string,
	taskUpdatedAt string,
	taskCreatedBy string,
	taskUpdatedBy string,
	taskLastExecutedBy string,
	taskLastExecutedAt *time.Time,
	config tp.JSONB,
	active bool,
) TasksModel {
	return &m.TasksModel{
		ID:                 uuid.New().String(),
		MCPProvider:        provider,
		TargetTask:         target,
		LastSynced:         nil,
		CreatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		CreatedBy:          "",
		UpdatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		UpdatedBy:          "",
		TaskType:           taskType,
		TaskSchedule:       taskSchedule,
		TaskExpression:     taskExpression,
		TaskCommandType:    taskCommandType,
		TaskMethod:         taskMethod,
		TaskAPIEndpoint:    taskAPIEndpoint,
		TaskPayload:        taskPayload,
		TaskHeaders:        taskHeaders,
		TaskRetries:        taskRetries,
		TaskTimeout:        taskTimeout,
		TaskStatus:         taskStatus,
		TaskNextRun:        taskNextRun,
		TaskLastRun:        taskLastRun,
		TaskLastRunStatus:  taskLastRunStatus,
		TaskLastRunMessage: taskLastRunMessage,
		TaskCommand:        taskCommand,
		TaskActivated:      taskActivated,
		TaskConfig:         taskConfig,
		TaskTags:           taskTags,
		TaskPriority:       taskPriority,
		TaskNotes:          taskNotes,
		TaskCreatedAt:      taskCreatedAt,
		TaskUpdatedAt:      taskUpdatedAt,
		TaskCreatedBy:      taskCreatedBy,
		TaskUpdatedBy:      taskUpdatedBy,
		TaskLastExecutedBy: taskLastExecutedBy,
		TaskLastExecutedAt: taskLastExecutedAt,
		Config:             config,
		Active:             active,
	}
}
