package models

import (
	"context"

	m "github.com/kubex-ecosystem/gdbase/internal/models/mcp/analysis_jobs"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type AnalysisJob = m.AnalysisJob
type AnalysisJobModel = m.IAnalysisJob
type AnalysisJobService = m.IAnalysisJobService
type AnalysisJobRepo = m.IAnalysisJobRepo

func NewAnalysisJobService(analysisJobRepo AnalysisJobRepo) AnalysisJobService {
	return m.NewAnalysisJobService(analysisJobRepo)
}

func NewAnalysisJobRepo(ctx context.Context, dbService *svc.DBServiceImpl) AnalysisJobRepo {
	return m.NewAnalysisJobRepository(ctx, dbService)
}

func NewAnalysisJobModel() AnalysisJobModel {
	return m.NewAnalysisJobModel()
}
