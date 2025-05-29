package services

import (
	jobqueue "github.com/rafa-mori/gdbase/internal/models/job_queue"
)

// ExecutionLogService implements the IExecutionLogService interface.
type ExecutionLogService struct {
	// Add dependencies like database connection or logger if needed.
}

// CreateLog creates a new execution log entry.
func (s *ExecutionLogService) CreateLog(log jobqueue.ExecutionLog) error {
	// Implement the logic to save the log to the database.
	return nil
}
