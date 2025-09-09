package tests

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/kubex-ecosystem/gdbase/internal/controllers"
// 	jobqueue "github.com/kubex-ecosystem/gdbase/internal/models/job_queue"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// type MockExecutionLogService struct{}

// func (m *MockExecutionLogService) CreateLog(log jobqueue.ExecutionLog) error {
// 	return nil
// }

// func TestCreateLog(t *testing.T) {
// 	service := &MockExecutionLogService{}
// 	controller := controllers.NewExecutionLogController(service)

// 	router := gin.Default()
// 	router.POST("/logs", controller.CreateLog)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/logs", nil)

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// }
