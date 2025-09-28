package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	n "github.com/kubex-ecosystem/gdbase/internal/models/mcp/notifications"
	t "github.com/kubex-ecosystem/gdbase/types"
)

// Notification Rules
type NotificationRule = n.NotificationRule
type NotificationRuleModel = n.INotificationRule
type NotificationRuleStatus = n.NotificationRuleStatus
type NotificationRuleCondition = n.NotificationRuleCondition
type NotificationRulePlatform = n.NotificationRulePlatform
type NotificationRulePriority = n.NotificationRulePriority

// Notification Rules Constants
const (
	// Status
	NotificationRuleStatusActive   = n.NotificationRuleStatusActive
	NotificationRuleStatusInactive = n.NotificationRuleStatusInactive
	NotificationRuleStatusPaused   = n.NotificationRuleStatusPaused

	// Conditions
	NotificationRuleConditionJobCompleted = n.NotificationRuleConditionJobCompleted
	NotificationRuleConditionJobFailed    = n.NotificationRuleConditionJobFailed
	NotificationRuleConditionJobStarted   = n.NotificationRuleConditionJobStarted
	NotificationRuleConditionJobRetried   = n.NotificationRuleConditionJobRetried
	NotificationRuleConditionScoreAlert   = n.NotificationRuleConditionScoreAlert
	NotificationRuleConditionTimeAlert    = n.NotificationRuleConditionTimeAlert

	// Platforms
	NotificationRulePlatformTelegram = n.NotificationRulePlatformTelegram
	NotificationRulePlatformDiscord  = n.NotificationRulePlatformDiscord
	NotificationRulePlatformWhatsApp = n.NotificationRulePlatformWhatsApp
	NotificationRulePlatformEmail    = n.NotificationRulePlatformEmail
	NotificationRulePlatformSlack    = n.NotificationRulePlatformSlack

	// Priority
	NotificationRulePriorityLow      = n.NotificationRulePriorityLow
	NotificationRulePriorityMedium   = n.NotificationRulePriorityMedium
	NotificationRulePriorityHigh     = n.NotificationRulePriorityHigh
	NotificationRulePriorityCritical = n.NotificationRulePriorityCritical
)

// Notification Templates
type NotificationTemplate = n.NotificationTemplate
type NotificationTemplateModel = n.INotificationTemplate
type NotificationTemplateType = n.NotificationTemplateType
type NotificationTemplateFormat = n.NotificationTemplateFormat
type NotificationTemplateStatus = n.NotificationTemplateStatus

// Notification Templates Constants
const (
	// Types
	NotificationTemplateTypeJobCompleted = n.NotificationTemplateTypeJobCompleted
	NotificationTemplateTypeJobFailed    = n.NotificationTemplateTypeJobFailed
	NotificationTemplateTypeJobStarted   = n.NotificationTemplateTypeJobStarted
	NotificationTemplateTypeJobRetried   = n.NotificationTemplateTypeJobRetried
	NotificationTemplateTypeScoreAlert   = n.NotificationTemplateTypeScoreAlert
	NotificationTemplateTypeTimeAlert    = n.NotificationTemplateTypeTimeAlert
	NotificationTemplateTypeCustom       = n.NotificationTemplateTypeCustom

	// Formats
	NotificationTemplateFormatText     = n.NotificationTemplateFormatText
	NotificationTemplateFormatMarkdown = n.NotificationTemplateFormatMarkdown
	NotificationTemplateFormatHTML     = n.NotificationTemplateFormatHTML
	NotificationTemplateFormatJSON     = n.NotificationTemplateFormatJSON

	// Status
	NotificationTemplateStatusActive   = n.NotificationTemplateStatusActive
	NotificationTemplateStatusInactive = n.NotificationTemplateStatusInactive
	NotificationTemplateStatusDraft    = n.NotificationTemplateStatusDraft
)

// Notification History
type NotificationHistory = n.NotificationHistory
type NotificationHistoryModel = n.INotificationHistory
type NotificationHistoryStatus = n.NotificationHistoryStatus
type NotificationHistoryPlatform = n.NotificationHistoryPlatform

// Notification History Constants
const (
	// Status
	NotificationHistoryStatusPending   = n.NotificationHistoryStatusPending
	NotificationHistoryStatusSent      = n.NotificationHistoryStatusSent
	NotificationHistoryStatusDelivered = n.NotificationHistoryStatusDelivered
	NotificationHistoryStatusFailed    = n.NotificationHistoryStatusFailed
	NotificationHistoryStatusRetrying  = n.NotificationHistoryStatusRetrying
	NotificationHistoryStatusCancelled = n.NotificationHistoryStatusCancelled
	NotificationHistoryStatusRead      = n.NotificationHistoryStatusRead

	// Platforms
	NotificationHistoryPlatformTelegram = n.NotificationHistoryPlatformTelegram
	NotificationHistoryPlatformDiscord  = n.NotificationHistoryPlatformDiscord
	NotificationHistoryPlatformWhatsApp = n.NotificationHistoryPlatformWhatsApp
	NotificationHistoryPlatformEmail    = n.NotificationHistoryPlatformEmail
	NotificationHistoryPlatformSlack    = n.NotificationHistoryPlatformSlack
	NotificationHistoryPlatformWebhook  = n.NotificationHistoryPlatformWebhook
)

// Notification Events
type NotificationEvent = n.NotificationEvent
type NotificationEventType = n.NotificationEventType
type NotificationEventProcessor = n.INotificationEventProcessor
type NotificationEventHandler = n.EventHandler

// Notification Events Constants
const (
	NotificationEventTypeJobStatusChanged = n.NotificationEventTypeJobStatusChanged
	NotificationEventTypeJobCompleted     = n.NotificationEventTypeJobCompleted
	NotificationEventTypeJobFailed        = n.NotificationEventTypeJobFailed
	NotificationEventTypeJobStarted       = n.NotificationEventTypeJobStarted
	NotificationEventTypeJobRetried       = n.NotificationEventTypeJobRetried
	NotificationEventTypeScoreAlert       = n.NotificationEventTypeScoreAlert
	NotificationEventTypeTimeAlert        = n.NotificationEventTypeTimeAlert
	NotificationEventTypeSystemAlert      = n.NotificationEventTypeSystemAlert
)

// Integration
type AnalysisJobNotificationIntegration = n.AnalysisJobNotificationIntegration
type AnalysisJobNotificationHooks = n.AnalysisJobNotificationHooks

// Factory functions for Notification Rules
func NewNotificationRuleModel() NotificationRuleModel {
	return n.NewNotificationRuleModel()
}

// Repository and Service interfaces
type NotificationRuleRepo = n.INotificationRuleRepo
type NotificationRuleService interface {
	CreateRule(ctx context.Context, rule NotificationRuleModel) (NotificationRuleModel, error)
	GetRuleByID(ctx context.Context, id uuid.UUID) (NotificationRuleModel, error)
	UpdateRule(ctx context.Context, rule NotificationRuleModel) (NotificationRuleModel, error)
	DeleteRule(ctx context.Context, id uuid.UUID) error
	ListRules(ctx context.Context, userID uuid.UUID) ([]NotificationRuleModel, error)
	ListActiveRules(ctx context.Context) ([]NotificationRuleModel, error)
	EnableRule(ctx context.Context, id uuid.UUID) error
	DisableRule(ctx context.Context, id uuid.UUID) error
	GetRulesForJobEvent(ctx context.Context, eventType NotificationEventType, jobType string, userID, projectID uuid.UUID) ([]NotificationRuleModel, error)
}

// Factory functions for Notification Templates
func NewNotificationTemplateModel() NotificationTemplateModel {
	return n.NewNotificationTemplateModel()
}

func NewJobCompletedTemplate(name, language string, createdBy uuid.UUID) NotificationTemplateModel {
	return n.NewJobCompletedTemplate(name, language, createdBy)
}

func NewJobFailedTemplate(name, language string, createdBy uuid.UUID) NotificationTemplateModel {
	return n.NewJobFailedTemplate(name, language, createdBy)
}

func NewScoreAlertTemplate(name, language string, createdBy uuid.UUID) NotificationTemplateModel {
	return n.NewScoreAlertTemplate(name, language, createdBy)
}

// Repository and Service interfaces for Templates
type NotificationTemplateRepo = n.INotificationTemplateRepo
type NotificationTemplateService interface {
	CreateTemplate(ctx context.Context, template NotificationTemplateModel) (NotificationTemplateModel, error)
	GetTemplateByID(ctx context.Context, id uuid.UUID) (NotificationTemplateModel, error)
	UpdateTemplate(ctx context.Context, template NotificationTemplateModel) (NotificationTemplateModel, error)
	DeleteTemplate(ctx context.Context, id uuid.UUID) error
	ListTemplates(ctx context.Context, templateType NotificationTemplateType, language string) ([]NotificationTemplateModel, error)
	GetDefaultTemplate(ctx context.Context, templateType NotificationTemplateType, language string) (NotificationTemplateModel, error)
	RenderTemplate(ctx context.Context, templateID uuid.UUID, variables map[string]interface{}) (string, string, error) // subject, body, error
}

// Factory functions for Notification History
func NewNotificationHistoryModel() NotificationHistoryModel {
	return n.NewNotificationHistoryModel()
}

func NewNotificationHistoryFromJob(ruleID uuid.UUID, jobID uuid.UUID, platform NotificationHistoryPlatform, targetID, targetName, subject, message string) NotificationHistoryModel {
	return n.NewNotificationHistoryFromJob(ruleID, jobID, platform, targetID, targetName, subject, message)
}

func NewScheduledNotification(ruleID uuid.UUID, platform NotificationHistoryPlatform, targetID, targetName, subject, message string, scheduledFor time.Time) NotificationHistoryModel {
	return n.NewScheduledNotification(ruleID, platform, targetID, targetName, subject, message, scheduledFor)
}

// Repository and Service interfaces for History
type NotificationHistoryRepo = n.INotificationHistoryRepo
type NotificationHistoryService interface {
	CreateNotification(ctx context.Context, notification NotificationHistoryModel) (NotificationHistoryModel, error)
	GetNotificationByID(ctx context.Context, id uuid.UUID) (NotificationHistoryModel, error)
	UpdateNotification(ctx context.Context, notification NotificationHistoryModel) (NotificationHistoryModel, error)
	ListNotifications(ctx context.Context, ruleID *uuid.UUID, platform *NotificationHistoryPlatform, status *NotificationHistoryStatus, limit, offset int) ([]NotificationHistoryModel, int, error)
	GetNotificationStats(ctx context.Context, since time.Time) (map[string]interface{}, error)
	MarkAsSent(ctx context.Context, id uuid.UUID) error
	MarkAsDelivered(ctx context.Context, id uuid.UUID) error
	MarkAsFailed(ctx context.Context, id uuid.UUID, errorMessage string) error
	RetryNotification(ctx context.Context, id uuid.UUID) error
}

// Factory functions for Event Processing
func NewNotificationEventProcessor(
	ruleRepo NotificationRuleRepo,
	templateRepo NotificationTemplateRepo,
	historyRepo NotificationHistoryRepo,
	messageQueue n.MessageQueuePublisher,
) NotificationEventProcessor {
	return n.NewNotificationEventProcessor(ruleRepo, templateRepo, historyRepo, messageQueue)
}

// Factory functions for Integration
func NewAnalysisJobNotificationIntegration(eventProcessor NotificationEventProcessor) *AnalysisJobNotificationIntegration {
	return n.NewAnalysisJobNotificationIntegration(eventProcessor)
}

func NewAnalysisJobNotificationHooks(integration *AnalysisJobNotificationIntegration) *AnalysisJobNotificationHooks {
	return n.NewAnalysisJobNotificationHooks(integration)
}

// Event creation helpers
func NewJobStatusChangedEvent(jobID, userID uuid.UUID, projectID *uuid.UUID, oldStatus, newStatus string, jobData map[string]interface{}) *NotificationEvent {
	return n.NewJobStatusChangedEvent(jobID, userID, projectID, oldStatus, newStatus, jobData)
}

func NewJobCompletedEvent(jobID, userID uuid.UUID, projectID *uuid.UUID, jobData map[string]interface{}) *NotificationEvent {
	return n.NewJobCompletedEvent(jobID, userID, projectID, jobData)
}

func NewJobFailedEvent(jobID, userID uuid.UUID, projectID *uuid.UUID, errorMessage string, jobData map[string]interface{}) *NotificationEvent {
	return n.NewJobFailedEvent(jobID, userID, projectID, errorMessage, jobData)
}

func NewScoreAlertEvent(jobID, userID uuid.UUID, projectID *uuid.UUID, score float64, threshold float64, jobData map[string]interface{}) *NotificationEvent {
	return n.NewScoreAlertEvent(jobID, userID, projectID, score, threshold, jobData)
}

// Notification Management Service - High-level interface for managing all notification functionality
type NotificationManagementService interface {
	// Rules management
	CreateRule(ctx context.Context, rule NotificationRuleModel) (NotificationRuleModel, error)
	UpdateRule(ctx context.Context, rule NotificationRuleModel) (NotificationRuleModel, error)
	DeleteRule(ctx context.Context, id uuid.UUID) error
	GetRuleByID(ctx context.Context, id uuid.UUID) (NotificationRuleModel, error)
	ListUserRules(ctx context.Context, userID uuid.UUID) ([]NotificationRuleModel, error)

	// Templates management
	CreateTemplate(ctx context.Context, template NotificationTemplateModel) (NotificationTemplateModel, error)
	UpdateTemplate(ctx context.Context, template NotificationTemplateModel) (NotificationTemplateModel, error)
	GetTemplateByID(ctx context.Context, id uuid.UUID) (NotificationTemplateModel, error)
	ListTemplates(ctx context.Context, filters map[string]interface{}) ([]NotificationTemplateModel, error)

	// Event processing
	ProcessEvent(ctx context.Context, event *NotificationEvent) error

	// History and stats
	GetNotificationHistory(ctx context.Context, filters map[string]interface{}, limit, offset int) ([]NotificationHistoryModel, int, error)
	GetNotificationStats(ctx context.Context, since time.Time) (map[string]interface{}, error)

	// Integration management
	EnableNotifications(ctx context.Context) error
	DisableNotifications(ctx context.Context) error
	IsEnabled(ctx context.Context) bool
}

// CreateDefaultNotificationRules Helper function to create default notification rules for a user
func CreateDefaultNotificationRules(userID uuid.UUID, platforms map[string]NotificationRulePlatform) []NotificationRuleModel {
	rules := make([]NotificationRuleModel, 0)

	// Rule 1: Job Completion Notifications
	completedRule := NewNotificationRuleModel().(*NotificationRule)
	completedRule.Name = "Job Completed Notifications"
	completedRule.Description = "Notify when analysis jobs complete successfully"
	completedRule.Condition = NotificationRuleConditionJobCompleted
	completedRule.SetPlatforms(platforms)
	completedRule.Priority = NotificationRulePriorityMedium
	completedRule.Status = NotificationRuleStatusActive
	completedRule.CreatedBy = userID
	completedRule.MaxNotificationsPerHour = 20
	rules = append(rules, completedRule)

	// Rule 2: Job Failure Notifications
	failedRule := NewNotificationRuleModel().(*NotificationRule)
	failedRule.Name = "Job Failed Notifications"
	failedRule.Description = "Notify immediately when analysis jobs fail"
	failedRule.Condition = NotificationRuleConditionJobFailed
	failedRule.SetPlatforms(platforms)
	failedRule.Priority = NotificationRulePriorityHigh
	failedRule.Status = NotificationRuleStatusActive
	failedRule.CreatedBy = userID
	failedRule.MaxNotificationsPerHour = 10
	rules = append(rules, failedRule)

	// Rule 3: Score Alert Notifications
	scoreRule := NewNotificationRuleModel().(*NotificationRule)
	scoreRule.Name = "Low Score Alerts"
	scoreRule.Description = "Notify when analysis scores are below threshold"
	scoreRule.Condition = NotificationRuleConditionScoreAlert
	scoreRule.SetPlatforms(platforms)
	scoreRule.Priority = NotificationRulePriorityMedium
	scoreRule.Status = NotificationRuleStatusActive
	scoreRule.CreatedBy = userID
	scoreRule.MaxNotificationsPerHour = 5
	scoreRule.CooldownMinutes = 60 // Don't spam with score alerts

	// Set trigger config for score threshold
	scoreRule.TriggerConfig = map[string]interface{}{
		"score_threshold": 0.7,
		"operator":        "lt", // less than
	}
	rules = append(rules, scoreRule)

	return rules
}

// Helper function to create default notification templates
func CreateDefaultNotificationTemplates(createdBy uuid.UUID, language string) []NotificationTemplateModel {
	templates := make([]NotificationTemplateModel, 0)

	// Job Completed Template
	completedTemplate := NewJobCompletedTemplate("Default Job Completed", language, createdBy)
	templates = append(templates, completedTemplate)

	// Job Failed Template
	failedTemplate := NewJobFailedTemplate("Default Job Failed", language, createdBy)
	templates = append(templates, failedTemplate)

	// Score Alert Template
	scoreTemplate := NewScoreAlertTemplate("Default Score Alert", language, createdBy)
	templates = append(templates, scoreTemplate)

	return templates
}

// Configuration helper for target setup
type NotificationTargetConfig struct {
	Platform NotificationRulePlatform `json:"platform"`
	TargetID string                   `json:"target_id"`
	Name     string                   `json:"name"`
	Config   map[string]interface{}   `json:"config,omitempty"`
}

// Helper function to create target configuration for rules
func CreateTargetConfig(targets []NotificationTargetConfig) t.JSONB {
	config := make(map[string]interface{})

	for _, target := range targets {
		platformKey := string(target.Platform)

		// Initialize platform config if not exists
		if config[platformKey] == nil {
			config[platformKey] = make([]interface{}, 0)
		}

		// Add target to platform config
		platformTargets := config[platformKey].([]interface{})
		targetConfig := map[string]interface{}{
			"id":   target.TargetID,
			"name": target.Name,
		}

		// Add additional config if provided
		if target.Config != nil {
			for k, v := range target.Config {
				targetConfig[k] = v
			}
		}

		platformTargets = append(platformTargets, targetConfig)
		config[platformKey] = platformTargets
	}

	return config
}
