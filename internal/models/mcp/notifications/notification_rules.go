package notifications

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/types"
)

// NotificationRuleStatus representa o status de uma regra de notificação
type NotificationRuleStatus string

const (
	NotificationRuleStatusActive   NotificationRuleStatus = "ACTIVE"
	NotificationRuleStatusInactive NotificationRuleStatus = "INACTIVE"
	NotificationRuleStatusPaused   NotificationRuleStatus = "PAUSED"
)

// NotificationRuleCondition representa condições que disparam notificações
type NotificationRuleCondition string

const (
	NotificationRuleConditionJobCompleted NotificationRuleCondition = "JOB_COMPLETED"
	NotificationRuleConditionJobFailed    NotificationRuleCondition = "JOB_FAILED"
	NotificationRuleConditionJobStarted   NotificationRuleCondition = "JOB_STARTED"
	NotificationRuleConditionJobRetried   NotificationRuleCondition = "JOB_RETRIED"
	NotificationRuleConditionScoreAlert   NotificationRuleCondition = "SCORE_ALERT"
	NotificationRuleConditionTimeAlert    NotificationRuleCondition = "TIME_ALERT"
)

// NotificationRulePlatform representa plataformas de notificação
type NotificationRulePlatform string

const (
	NotificationRulePlatformTelegram NotificationRulePlatform = "TELEGRAM"
	NotificationRulePlatformDiscord  NotificationRulePlatform = "DISCORD"
	NotificationRulePlatformWhatsApp NotificationRulePlatform = "WHATSAPP"
	NotificationRulePlatformEmail    NotificationRulePlatform = "EMAIL"
	NotificationRulePlatformSlack    NotificationRulePlatform = "SLACK"
)

// NotificationRulePriority representa prioridade da notificação
type NotificationRulePriority string

const (
	NotificationRulePriorityLow      NotificationRulePriority = "LOW"
	NotificationRulePriorityMedium   NotificationRulePriority = "MEDIUM"
	NotificationRulePriorityHigh     NotificationRulePriority = "HIGH"
	NotificationRulePriorityCritical NotificationRulePriority = "CRITICAL"
)

// INotificationRule interface para regras de notificação
type INotificationRule interface {
	TableName() string
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
	GetCondition() NotificationRuleCondition
	SetCondition(condition NotificationRuleCondition)
	GetPlatforms() map[string]NotificationRulePlatform
	SetPlatforms(platforms map[string]NotificationRulePlatform)
	GetJobTypes() map[string]string
	SetJobTypes(jobTypes map[string]string)
	GetUserIDs() map[string]uuid.UUID
	SetUserIDs(userIDs map[string]uuid.UUID)
	GetProjectIDs() []uuid.UUID
	SetProjectIDs(projectIDs []uuid.UUID)
	GetPriority() NotificationRulePriority
	SetPriority(priority NotificationRulePriority)
	GetStatus() NotificationRuleStatus
	SetStatus(status NotificationRuleStatus)
	GetTriggerConfig() t.JSONB
	SetTriggerConfig(config t.JSONB)
	GetTargetConfig() t.JSONB
	SetTargetConfig(config t.JSONB)
	GetScheduleConfig() t.JSONB
	SetScheduleConfig(config t.JSONB)
	GetTemplateID() *uuid.UUID
	SetTemplateID(templateID *uuid.UUID)
	GetCooldownMinutes() int
	SetCooldownMinutes(minutes int)
	GetMaxNotificationsPerHour() int
	SetMaxNotificationsPerHour(max int)
	GetIsGlobal() bool
	SetIsGlobal(isGlobal bool)
	GetCreatedBy() uuid.UUID
	SetCreatedBy(createdBy uuid.UUID)
	GetUpdatedBy() *uuid.UUID
	SetUpdatedBy(updatedBy *uuid.UUID)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetLastTriggeredAt() *time.Time
	SetLastTriggeredAt(lastTriggered *time.Time)
	GetTriggerCount() int64
	SetTriggerCount(count int64)

	// Validation methods
	Validate() error
	IsEligibleForJobType(jobType string) bool
	IsEligibleForUser(userID uuid.UUID) bool
	IsEligibleForProject(projectID uuid.UUID) bool
	CanTrigger() bool
	ShouldTriggerForScore(score float64) bool
	ShouldTriggerForDuration(duration time.Duration) bool
}

// NotificationRule implementação da regra de notificação
type NotificationRule struct {
	ID                      uuid.UUID                 `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name                    string                    `json:"name" xml:"name" yaml:"name" gorm:"column:name;not null;type:VARCHAR(255)"`
	Description             string                    `json:"description" xml:"description" yaml:"description" gorm:"column:description;type:TEXT"`
	Condition               NotificationRuleCondition `json:"condition" xml:"condition" yaml:"condition" gorm:"column:condition;not null;type:notification_rule_condition"`
	Platforms               t.JSONB                   `json:"platforms" xml:"platforms" yaml:"platforms" gorm:"column:platforms;type:jsonb"`         // Array de NotificationRulePlatform
	JobTypes                t.JSONB                   `json:"job_types" xml:"job_types" yaml:"job_types" gorm:"column:job_types;type:jsonb"`         // Array de strings
	UserIDs                 t.JSONB                   `json:"user_ids" xml:"user_ids" yaml:"user_ids" gorm:"column:user_ids;type:jsonb"`             // Array de UUIDs
	ProjectIDs              t.JSONB                   `json:"project_ids" xml:"project_ids" yaml:"project_ids" gorm:"column:project_ids;type:jsonb"` // Array de UUIDs
	Priority                NotificationRulePriority  `json:"priority" xml:"priority" yaml:"priority" gorm:"column:priority;default:'MEDIUM';type:notification_rule_priority"`
	Status                  NotificationRuleStatus    `json:"status" xml:"status" yaml:"status" gorm:"column:status;default:'ACTIVE';type:notification_rule_status"`
	TriggerConfig           t.JSONB                   `json:"trigger_config" xml:"trigger_config" yaml:"trigger_config" gorm:"column:trigger_config;type:jsonb"`     // Configurações específicas para disparo (ex: score_threshold, time_threshold)
	TargetConfig            t.JSONB                   `json:"target_config" xml:"target_config" yaml:"target_config" gorm:"column:target_config;type:jsonb"`         // Configurações de destino (chat_ids, channels, emails)
	ScheduleConfig          t.JSONB                   `json:"schedule_config" xml:"schedule_config" yaml:"schedule_config" gorm:"column:schedule_config;type:jsonb"` // Configurações de agendamento (horários, dias da semana)
	TemplateID              *uuid.UUID                `json:"template_id" xml:"template_id" yaml:"template_id" gorm:"column:template_id;type:uuid"`
	CooldownMinutes         int                       `json:"cooldown_minutes" xml:"cooldown_minutes" yaml:"cooldown_minutes" gorm:"column:cooldown_minutes;default:0"`
	MaxNotificationsPerHour int                       `json:"max_notifications_per_hour" xml:"max_notifications_per_hour" yaml:"max_notifications_per_hour" gorm:"column:max_notifications_per_hour;default:10"`
	IsGlobal                bool                      `json:"is_global" xml:"is_global" yaml:"is_global" gorm:"column:is_global;default:false"`
	CreatedBy               uuid.UUID                 `json:"created_by" xml:"created_by" yaml:"created_by" gorm:"column:created_by;type:uuid"`
	UpdatedBy               *uuid.UUID                `json:"updated_by" xml:"updated_by" yaml:"updated_by" gorm:"column:updated_by;type:uuid"`
	CreatedAt               time.Time                 `json:"created_at" xml:"created_at" yaml:"created_at" gorm:"column:created_at;default:now()"`
	UpdatedAt               time.Time                 `json:"updated_at" xml:"updated_at" yaml:"updated_at" gorm:"column:updated_at;default:now()"`
	LastTriggeredAt         *time.Time                `json:"last_triggered_at" xml:"last_triggered_at" yaml:"last_triggered_at" gorm:"column:last_triggered_at;type:timestamp"`
	TriggerCount            int64                     `json:"trigger_count" xml:"trigger_count" yaml:"trigger_count" gorm:"column:trigger_count;default:0"`
}

// NewNotificationRuleModel cria uma nova instância de regra de notificação
func NewNotificationRuleModel() INotificationRule {
	return &NotificationRule{
		Platforms:      make(t.JSONB),
		JobTypes:       make(t.JSONB),
		UserIDs:        make(t.JSONB),
		ProjectIDs:     make(t.JSONB),
		TriggerConfig:  make(t.JSONB),
		TargetConfig:   make(t.JSONB),
		ScheduleConfig: make(t.JSONB),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

// TableName retorna o nome da tabela
func (n *NotificationRule) TableName() string {
	return "mcp_notification_rules"
}

// Getters e Setters

func (n *NotificationRule) GetID() uuid.UUID                                 { return n.ID }
func (n *NotificationRule) SetID(id uuid.UUID)                               { n.ID = id }
func (n *NotificationRule) GetName() string                                  { return n.Name }
func (n *NotificationRule) SetName(name string)                              { n.Name = name }
func (n *NotificationRule) GetDescription() string                           { return n.Description }
func (n *NotificationRule) SetDescription(description string)                { n.Description = description }
func (n *NotificationRule) GetCondition() NotificationRuleCondition          { return n.Condition }
func (n *NotificationRule) SetCondition(condition NotificationRuleCondition) { n.Condition = condition }
func (n *NotificationRule) GetPriority() NotificationRulePriority            { return n.Priority }
func (n *NotificationRule) SetPriority(priority NotificationRulePriority)    { n.Priority = priority }
func (n *NotificationRule) GetStatus() NotificationRuleStatus                { return n.Status }
func (n *NotificationRule) SetStatus(status NotificationRuleStatus)          { n.Status = status }
func (n *NotificationRule) GetTriggerConfig() t.JSONB                        { return n.TriggerConfig }
func (n *NotificationRule) SetTriggerConfig(config t.JSONB)                  { n.TriggerConfig = config }
func (n *NotificationRule) GetTargetConfig() t.JSONB                         { return n.TargetConfig }
func (n *NotificationRule) SetTargetConfig(config t.JSONB)                   { n.TargetConfig = config }
func (n *NotificationRule) GetScheduleConfig() t.JSONB                       { return n.ScheduleConfig }
func (n *NotificationRule) SetScheduleConfig(config t.JSONB)                 { n.ScheduleConfig = config }
func (n *NotificationRule) GetTemplateID() *uuid.UUID                        { return n.TemplateID }
func (n *NotificationRule) SetTemplateID(templateID *uuid.UUID)              { n.TemplateID = templateID }
func (n *NotificationRule) GetCooldownMinutes() int                          { return n.CooldownMinutes }
func (n *NotificationRule) SetCooldownMinutes(minutes int)                   { n.CooldownMinutes = minutes }
func (n *NotificationRule) GetMaxNotificationsPerHour() int                  { return n.MaxNotificationsPerHour }
func (n *NotificationRule) SetMaxNotificationsPerHour(max int)               { n.MaxNotificationsPerHour = max }
func (n *NotificationRule) GetIsGlobal() bool                                { return n.IsGlobal }
func (n *NotificationRule) SetIsGlobal(isGlobal bool)                        { n.IsGlobal = isGlobal }
func (n *NotificationRule) GetCreatedBy() uuid.UUID                          { return n.CreatedBy }
func (n *NotificationRule) SetCreatedBy(createdBy uuid.UUID)                 { n.CreatedBy = createdBy }
func (n *NotificationRule) GetUpdatedBy() *uuid.UUID                         { return n.UpdatedBy }
func (n *NotificationRule) SetUpdatedBy(updatedBy *uuid.UUID)                { n.UpdatedBy = updatedBy }
func (n *NotificationRule) GetCreatedAt() time.Time                          { return n.CreatedAt }
func (n *NotificationRule) SetCreatedAt(createdAt time.Time)                 { n.CreatedAt = createdAt }
func (n *NotificationRule) GetUpdatedAt() time.Time                          { return n.UpdatedAt }
func (n *NotificationRule) SetUpdatedAt(updatedAt time.Time)                 { n.UpdatedAt = updatedAt }
func (n *NotificationRule) GetLastTriggeredAt() *time.Time                   { return n.LastTriggeredAt }
func (n *NotificationRule) SetLastTriggeredAt(lastTriggered *time.Time) {
	n.LastTriggeredAt = lastTriggered
}
func (n *NotificationRule) GetTriggerCount() int64      { return n.TriggerCount }
func (n *NotificationRule) SetTriggerCount(count int64) { n.TriggerCount = count }

// GetPlatforms para slices (implementação específica para JSONB)
func (n *NotificationRule) GetPlatforms() map[string]NotificationRulePlatform {
	if n.Platforms.IsNil() || n.Platforms.IsEmpty() {
		return map[string]NotificationRulePlatform{}
	}

	var platforms map[string]NotificationRulePlatform
	if !n.Platforms.IsNil() {
		platforms = make(map[string]NotificationRulePlatform)
		for _, p := range n.Platforms {
			if platformStr, ok := p.(string); ok {
				platforms[platformStr] = NotificationRulePlatform(platformStr)
			}
		}
	}
	return platforms
}

func (n *NotificationRule) SetPlatforms(platforms map[string]NotificationRulePlatform) {
	platformStrs := make(map[string]interface{})
	for i, p := range platforms {
		platformStrs[i] = string(p)
	}
	n.Platforms = platformStrs
}

func (n *NotificationRule) GetJobTypes() map[string]string {
	if n.JobTypes.IsNil() || n.JobTypes.IsEmpty() {
		return map[string]string{}
	}

	var jobTypes map[string]string
	if !n.JobTypes.IsNil() {
		jobTypes = make(map[string]string)
		for _, jt := range n.JobTypes {
			if jobTypeStr, ok := jt.(string); ok {
				jobTypes[jobTypeStr] = jobTypeStr
			}
		}
	}
	return jobTypes
}

func (n *NotificationRule) SetJobTypes(jobTypes map[string]string) {
	jobTypeStrs := make(map[string]interface{})
	for i, jt := range jobTypes {
		jobTypeStrs[i] = jt
	}
	n.JobTypes = jobTypeStrs
}

func (n *NotificationRule) GetUserIDs() map[string]uuid.UUID {
	if n.UserIDs.IsNil() || n.UserIDs.IsEmpty() {
		return map[string]uuid.UUID{}
	}

	userIDs := make(map[string]uuid.UUID)
	if !n.UserIDs.IsNil() {
		for _, uid := range n.UserIDs {
			if userIDStr, ok := uid.(string); ok {
				if userID, err := uuid.Parse(userIDStr); err == nil {
					userIDs[userIDStr] = userID
				}
			}
		}
	}
	return userIDs
}

func (n *NotificationRule) SetUserIDs(userIDs map[string]uuid.UUID) {
	userIDStrs := make(map[string]interface{})
	for _, uid := range userIDs {
		userIDStrs[uid.String()] = uid
	}
	n.UserIDs = userIDStrs
}

func (n *NotificationRule) GetProjectIDs() []uuid.UUID {
	if n.ProjectIDs.IsNil() || n.ProjectIDs.IsEmpty() {
		return []uuid.UUID{}
	}

	var projectIDs map[string]uuid.UUID
	if !n.ProjectIDs.IsNil() {
		projectIDs = make(map[string]uuid.UUID)
		for _, pid := range n.ProjectIDs {
			if projectIDStr, ok := pid.(string); ok {
				if projectID, err := uuid.Parse(projectIDStr); err == nil {
					projectIDs[projectIDStr] = projectID
				}
			}
		}
	}

	ids := make([]uuid.UUID, 0, len(projectIDs))
	for _, id := range projectIDs {
		ids = append(ids, id)
	}
	return ids
}

func (n *NotificationRule) SetProjectIDs(projectIDs []uuid.UUID) {
	projectIDStrs := make(map[string]interface{})
	for _, pid := range projectIDs {
		projectIDStrs[pid.String()] = pid
	}
	n.ProjectIDs = projectIDStrs
}

// Validation methods

func (n *NotificationRule) Validate() error {
	if n.Name == "" {
		return fmt.Errorf("notification rule name is required")
	}

	if n.Condition == "" {
		return fmt.Errorf("notification rule condition is required")
	}

	platforms := n.GetPlatforms()
	if len(platforms) == 0 {
		return fmt.Errorf("at least one platform is required")
	}

	if n.CreatedBy == uuid.Nil {
		return fmt.Errorf("created_by is required")
	}

	if n.CooldownMinutes < 0 {
		return fmt.Errorf("cooldown_minutes cannot be negative")
	}

	if n.MaxNotificationsPerHour <= 0 {
		return fmt.Errorf("max_notifications_per_hour must be positive")
	}

	return nil
}

func (n *NotificationRule) IsEligibleForJobType(jobType string) bool {
	jobTypes := n.GetJobTypes()
	if len(jobTypes) == 0 {
		return true // Se não especificado, aplica para todos
	}

	for _, jt := range jobTypes {
		if jt == jobType {
			return true
		}
	}
	return false
}

func (n *NotificationRule) IsEligibleForUser(userID uuid.UUID) bool {
	if n.IsGlobal {
		return true
	}

	userIDs := n.GetUserIDs()
	if len(userIDs) == 0 {
		return true // Se não especificado, aplica para todos
	}

	for _, uid := range userIDs {
		if uid == userID {
			return true
		}
	}
	return false
}

func (n *NotificationRule) IsEligibleForProject(projectID uuid.UUID) bool {
	projectIDs := n.GetProjectIDs()
	if len(projectIDs) == 0 {
		return true // Se não especificado, aplica para todos
	}

	for _, pid := range projectIDs {
		if pid == projectID {
			return true
		}
	}
	return false
}

func (n *NotificationRule) CanTrigger() bool {
	if n.Status != NotificationRuleStatusActive {
		return false
	}

	if n.LastTriggeredAt != nil && n.CooldownMinutes > 0 {
		cooldownEnd := n.LastTriggeredAt.Add(time.Duration(n.CooldownMinutes) * time.Minute)
		if time.Now().Before(cooldownEnd) {
			return false
		}
	}

	return true
}

func (n *NotificationRule) ShouldTriggerForScore(score float64) bool {
	if n.Condition != NotificationRuleConditionScoreAlert {
		return false
	}

	if n.TriggerConfig.IsNil() || n.TriggerConfig.IsEmpty() {
		return false
	}

	var configMap = make(map[string]interface{})
	if !n.TriggerConfig.IsNil() {
		for k, v := range n.TriggerConfig {
			configMap[k] = v
		}
	}

	threshold, ok := configMap["score_threshold"].(float64)
	if !ok {
		return false
	}

	operator, ok := configMap["operator"].(string)
	if !ok {
		operator = "lt" // default: less than
	}

	switch operator {
	case "lt":
		return score < threshold
	case "lte":
		return score <= threshold
	case "gt":
		return score > threshold
	case "gte":
		return score >= threshold
	case "eq":
		return score == threshold
	default:
		return score < threshold
	}
}

func (n *NotificationRule) ShouldTriggerForDuration(duration time.Duration) bool {
	if n.Condition != NotificationRuleConditionTimeAlert {
		return false
	}

	if n.TriggerConfig.IsNil() || n.TriggerConfig.IsEmpty() {
		return false
	}

	var configMap = make(map[string]interface{})
	if !n.TriggerConfig.IsNil() {
		for k, v := range n.TriggerConfig {
			configMap[k] = v
		}
	}

	thresholdMinutes, ok := configMap["duration_threshold_minutes"].(float64)
	if !ok {
		return false
	}

	threshold := time.Duration(thresholdMinutes) * time.Minute

	operator, ok := configMap["operator"].(string)
	if !ok {
		operator = "gt" // default: greater than
	}

	switch operator {
	case "lt":
		return duration < threshold
	case "lte":
		return duration <= threshold
	case "gt":
		return duration > threshold
	case "gte":
		return duration >= threshold
	default:
		return duration > threshold
	}
}
