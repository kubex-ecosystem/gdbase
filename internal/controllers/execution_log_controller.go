package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jobqueue "github.com/kubex-ecosystem/gdbase/internal/models/job_queue"
	"github.com/kubex-ecosystem/gdbase/internal/services"
)

// ExecutionLogController handles HTTP requests for execution logs.
type ExecutionLogController struct {
	Service services.ExecutionLogService
}

// NewExecutionLogController creates a new instance of ExecutionLogController.
func NewExecutionLogController(service services.ExecutionLogService) *ExecutionLogController {
	return &ExecutionLogController{Service: service}
}

// CreateLog handles the creation of a new execution log.
func (c *ExecutionLogController) CreateLog(ctx *gin.Context) {
	var log jobqueue.ExecutionLog
	if err := ctx.ShouldBindJSON(&log); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateLog(log); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Log created successfully"})
}
