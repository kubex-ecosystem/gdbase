package models

import (
	m "github.com/kubex-ecosystem/gdbase/internal/models/mcp/analysis_jobs"
	"gorm.io/gorm"
)

type AnalysisJob = m.AnalysisJob
type AnalysisJobModel = m.IAnalysisJob
type AnalysisJobService = m.IAnalysisJobService
type AnalysisJobRepo = m.IAnalysisJobRepo

func NewAnalysisJobService(analysisJobRepo AnalysisJobRepo) AnalysisJobService {
	return m.NewAnalysisJobService(analysisJobRepo)
}

func NewAnalysisJobRepo(db *gorm.DB) AnalysisJobRepo {
	return m.NewAnalysisJobRepository(db)
}

func NewAnalysisJobModel() AnalysisJobModel {
	return m.NewAnalysisJobModel()
}