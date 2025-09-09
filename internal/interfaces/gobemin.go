package interfaces

import (
	"net/http"

	l "github.com/kubex-ecosystem/logz"
)

type IGoBE interface {
	StartGoBE()
	HandleValidate(w http.ResponseWriter, r *http.Request)
	HandleContact(w http.ResponseWriter, r *http.Request)
	RateLimit(w http.ResponseWriter, r *http.Request) bool
	Initialize() error
	GetLogFilePath() string
	GetConfigFilePath() string
	GetLogger() l.Logger
	Mu() IMutexes
	GetReference() IReference
	Environment() IEnvironment
}
