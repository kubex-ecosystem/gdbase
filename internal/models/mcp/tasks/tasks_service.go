package tasks

import (
	"errors"
	"fmt"
	"time"

	t "github.com/rafa-mori/gdbase/types"
)

type TaskSearchOptions struct {
	Active       bool
	OutOfDate    bool
	Running      bool
	Provider     string
	Target       string
	TaskType     string
	TaskStatus   string
	UserID       string
	IncludeCron  bool
	IncludeTasks bool
}

type ITasksService interface {
	CreateTask(task ITasksModel) (ITasksModel, error)
	GetTaskByID(id string) (ITasksModel, error)
	UpdateTask(task ITasksModel) (ITasksModel, error)
	DeleteTask(id string) error
	ListTasks(opts *TaskSearchOptions) ([]ITasksModel, error)
	GetTasksByProvider(provider string) ([]ITasksModel, error)
	GetTasksByTarget(target string) ([]ITasksModel, error)
	GetTasksByProviderAndTarget(provider, target string) ([]ITasksModel, error)
	GetTasksByType(taskType TaskType) ([]ITasksModel, error)
	GetTasksByStatus(status TaskStatus) ([]ITasksModel, error)
	GetActiveTasks() ([]ITasksModel, error)
	GetTasksDueForExecution() ([]ITasksModel, error)
	GetTasksByUserID(userID string) ([]ITasksModel, error)

	// Task execution methods
	MarkTaskAsRunning(taskID string) error
	MarkTaskAsCompleted(taskID string, message string) error
	MarkTaskAsFailed(taskID string, message string) error
	UpdateTaskNextRun(taskID string, nextRun time.Time) error

	// CronJob integration methods
	ConvertTaskToCronJob(taskID string) (*CronJobIntegration, error)
	SyncTaskWithCronJob(taskID string, cronJob *CronJobIntegration) error

	GetContextDBService() t.IDBService
}

type TasksService struct {
	repo ITasksRepo
}

func NewTasksService(repo ITasksRepo) ITasksService {
	return &TasksService{repo: repo}
}

func (ts *TasksService) CreateTask(task ITasksModel) (ITasksModel, error) {
	if task.GetMCPProvider() == "" || task.GetTargetTask() == "" {
		return nil, errors.New("missing required fields: MCP provider and target task are required")
	}

	if task.GetTaskType() == "" {
		return nil, errors.New("missing required field: task type is required")
	}

	if task.GetTaskExpression() == "" {
		return nil, errors.New("missing required field: task expression is required")
	}

	// Validate the model
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	createdTask, err := ts.repo.Create(task)
	if err != nil {
		return nil, fmt.Errorf("error creating task: %w", err)
	}
	return createdTask, nil
}

func (ts *TasksService) GetTaskByID(id string) (ITasksModel, error) {
	task, err := ts.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching task: %w", err)
	}
	return task, nil
}

func (ts *TasksService) UpdateTask(task ITasksModel) (ITasksModel, error) {
	// Validate the model before updating
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	updatedTask, err := ts.repo.Update(task)
	if err != nil {
		return nil, fmt.Errorf("error updating task: %w", err)
	}
	return updatedTask, nil
}

func (ts *TasksService) DeleteTask(id string) error {
	err := ts.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}

func (ts *TasksService) ListTasks(opts *TaskSearchOptions) ([]ITasksModel, error) {

	whereClause := map[string]interface{}{"id != ?": "0000"} // Default to true for all tasks
	if opts != nil {
		if opts.Active {
			whereClause["active"] = true
		}
		if opts.OutOfDate {
			whereClause["task_next_run < ?"] = time.Now().Format(time.RFC3339)
		}
		if opts.Running {
			whereClause["task_status"] = string(TaskStatusRunning)
		}
	}

	tasks, err := ts.repo.FindAll(whereClause)
	if err != nil {
		return nil, fmt.Errorf("error listing tasks: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetTasksByProvider(provider string) ([]ITasksModel, error) {
	tasks, err := ts.repo.FindAll("provider = ?", provider)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks by provider: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetTasksByTarget(target string) ([]ITasksModel, error) {
	tasks, err := ts.repo.FindAll("target = ?", target)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks by target: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetTasksByProviderAndTarget(provider, target string) ([]ITasksModel, error) {
	tasks, err := ts.repo.FindAll("provider = ? AND target = ?", provider, target)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks by provider and target: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetTasksByType(taskType TaskType) ([]ITasksModel, error) {
	tasks, err := ts.repo.FindAll("task_type = ?", string(taskType))
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks by type: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetTasksByStatus(status TaskStatus) ([]ITasksModel, error) {
	tasks, err := ts.repo.FindAll("task_status = ?", string(status))
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks by status: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetActiveTasks() ([]ITasksModel, error) {
	tasks, err := ts.repo.FindAll("active = ? AND task_activated = ?", true, true)
	if err != nil {
		return nil, fmt.Errorf("error fetching active tasks: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetTasksDueForExecution() ([]ITasksModel, error) {
	now := time.Now()
	tasks, err := ts.repo.FindAll("active = ? AND task_activated = ? AND (task_next_run IS NULL OR task_next_run <= ?)", true, true, now)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks due for execution: %w", err)
	}
	return tasks, nil
}

func (ts *TasksService) GetTasksByUserID(userID string) ([]ITasksModel, error) {
	tasks, err := ts.repo.FindAll("created_by = ? OR updated_by = ?", userID, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks by user ID: %w", err)
	}
	return tasks, nil
}

// MarkTaskAsRunning marks a task as running
func (ts *TasksService) MarkTaskAsRunning(taskID string) error {
	task, err := ts.GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("error fetching task: %w", err)
	}

	// Update the task model directly
	if tModel, ok := task.(*TasksModel); ok {
		now := time.Now()
		tModel.TaskStatus = TaskStatusRunning
		tModel.TaskLastRun = &now
		tModel.TaskLastRunStatus = "running"
		tModel.TaskLastRunMessage = "Task execution started"
		tModel.Sanitize()

		_, err = ts.repo.Update(tModel)
		return err
	}
	return fmt.Errorf("invalid task model type")
}

func (ts *TasksService) MarkTaskAsCompleted(taskID string, message string) error {
	task, err := ts.GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("error fetching task: %w", err)
	}

	if tModel, ok := task.(*TasksModel); ok {
		now := time.Now()
		tModel.TaskStatus = TaskStatusCompleted
		tModel.LastSynced = &now
		tModel.TaskLastRunStatus = "success"
		tModel.TaskLastRunMessage = message
		tModel.Sanitize()

		_, err = ts.repo.Update(tModel)
		return err
	}
	return fmt.Errorf("invalid task model type")
}

func (ts *TasksService) MarkTaskAsFailed(taskID string, message string) error {
	task, err := ts.GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("error fetching task: %w", err)
	}

	if tModel, ok := task.(*TasksModel); ok {
		tModel.TaskStatus = TaskStatusFailed
		tModel.TaskLastRunStatus = "failed"
		tModel.TaskLastRunMessage = message
		tModel.Sanitize()

		_, err = ts.repo.Update(tModel)
		return err
	}
	return fmt.Errorf("invalid task model type")
}

func (ts *TasksService) UpdateTaskNextRun(taskID string, nextRun time.Time) error {
	task, err := ts.GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("error fetching task: %w", err)
	}

	if tModel, ok := task.(*TasksModel); ok {
		tModel.TaskNextRun = &nextRun
		tModel.Sanitize()

		_, err = ts.repo.Update(tModel)
		return err
	}
	return fmt.Errorf("invalid task model type")
}

// ConvertTaskToCronJob converts a task to a CronJob integration
func (ts *TasksService) ConvertTaskToCronJob(taskID string) (*CronJobIntegration, error) {
	task, err := ts.GetTaskByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("error fetching task: %w", err)
	}

	cronJob, err := task.ToCronJob()
	if err != nil {
		return nil, fmt.Errorf("error converting task to cronjob: %w", err)
	}

	return cronJob, nil
}

func (ts *TasksService) SyncTaskWithCronJob(taskID string, cronJob *CronJobIntegration) error {
	task, err := ts.GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("error fetching task: %w", err)
	}

	err = task.FromCronJob(cronJob)
	if err != nil {
		return fmt.Errorf("error syncing task with cronjob: %w", err)
	}

	_, err = ts.UpdateTask(task)
	if err != nil {
		return fmt.Errorf("error updating task after sync: %w", err)
	}

	return nil
}

func (ts *TasksService) GetContextDBService() t.IDBService {
	return ts.repo.GetContextDBService()
}
