package notifications

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

// NotificationHistoryStatus representa o status de uma notificação enviada
type NotificationHistoryStatus string

const (
	NotificationHistoryStatusPending   NotificationHistoryStatus = "PENDING"
	NotificationHistoryStatusSent      NotificationHistoryStatus = "SENT"
	NotificationHistoryStatusDelivered NotificationHistoryStatus = "DELIVERED"
	NotificationHistoryStatusFailed    NotificationHistoryStatus = "FAILED"
	NotificationHistoryStatusRetrying  NotificationHistoryStatus = "RETRYING"
	NotificationHistoryStatusCancelled NotificationHistoryStatus = "CANCELLED"
	NotificationHistoryStatusRead      NotificationHistoryStatus = "READ" // Para plataformas que suportam
)

// NotificationHistoryPlatform representa a plataforma de notificação
type NotificationHistoryPlatform string

const (
	NotificationHistoryPlatformTelegram NotificationHistoryPlatform = "TELEGRAM"
	NotificationHistoryPlatformDiscord  NotificationHistoryPlatform = "DISCORD"
	NotificationHistoryPlatformWhatsApp NotificationHistoryPlatform = "WHATSAPP"
	NotificationHistoryPlatformEmail    NotificationHistoryPlatform = "EMAIL"
	NotificationHistoryPlatformSlack    NotificationHistoryPlatform = "SLACK"
	NotificationHistoryPlatformWebhook  NotificationHistoryPlatform = "WEBHOOK"
)

// INotificationHistory interface para histórico de notificações
type INotificationHistory interface {
	TableName() string
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	GetRuleID() uuid.UUID
	SetRuleID(ruleID uuid.UUID)
	GetTemplateID() *uuid.UUID
	SetTemplateID(templateID *uuid.UUID)
	GetAnalysisJobID() *uuid.UUID
	SetAnalysisJobID(analysisJobID *uuid.UUID)
	GetPlatform() NotificationHistoryPlatform
	SetPlatform(platform NotificationHistoryPlatform)
	GetStatus() NotificationHistoryStatus
	SetStatus(status NotificationHistoryStatus)
	GetSubject() string
	SetSubject(subject string)
	GetMessage() string
	SetMessage(message string)
	GetTargetID() string
	SetTargetID(targetID string)
	GetTargetName() string
	SetTargetName(targetName string)
	GetPlatformConfig() t.JSONBImpl
	SetPlatformConfig(config t.JSONBImpl)
	GetResponse() t.JSONBImpl
	SetResponse(response t.JSONBImpl)
	GetErrorMessage() string
	SetErrorMessage(errorMessage string)
	GetRetryCount() int
	SetRetryCount(retryCount int)
	GetMaxRetries() int
	SetMaxRetries(maxRetries int)
	GetPriority() NotificationRulePriority
	SetPriority(priority NotificationRulePriority)
	GetScheduledFor() *time.Time
	SetScheduledFor(scheduledFor *time.Time)
	GetSentAt() *time.Time
	SetSentAt(sentAt *time.Time)
	GetDeliveredAt() *time.Time
	SetDeliveredAt(deliveredAt *time.Time)
	GetReadAt() *time.Time
	SetReadAt(readAt *time.Time)
	GetExpiresAt() *time.Time
	SetExpiresAt(expiresAt *time.Time)
	GetMetadata() t.JSONBImpl
	SetMetadata(metadata t.JSONBImpl)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)

	// Status management methods
	Validate() error
	CanRetry() bool
	ShouldExpire() bool
	GetDeliveryDuration() *time.Duration
	GetDeliveryStatus() string
	MarkAsSent() error
	MarkAsDelivered() error
	MarkAsFailed(errorMessage string) error
	MarkAsRead() error
	IncrementRetryCount() error
}

// NotificationHistory implementação do histórico de notificações
type NotificationHistory struct {
	ID             uuid.UUID                   `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	RuleID         uuid.UUID                   `json:"rule_id" xml:"rule_id" yaml:"rule_id" gorm:"column:rule_id;not null;type:uuid;index"`
	TemplateID     *uuid.UUID                  `json:"template_id" xml:"template_id" yaml:"template_id" gorm:"column:template_id;type:uuid;index"`
	AnalysisJobID  *uuid.UUID                  `json:"analysis_job_id" xml:"analysis_job_id" yaml:"analysis_job_id" gorm:"column:analysis_job_id;type:uuid;index"`
	Platform       NotificationHistoryPlatform `json:"platform" xml:"platform" yaml:"platform" gorm:"column:platform;not null;type:notification_history_platform;index"`
	Status         NotificationHistoryStatus   `json:"status" xml:"status" yaml:"status" gorm:"column:status;not null;default:'PENDING';type:notification_history_status;index"`
	Subject        string                      `json:"subject" xml:"subject" yaml:"subject" gorm:"column:subject;type:VARCHAR(500)"`
	Message        string                      `json:"message" xml:"message" yaml:"message" gorm:"column:message;not null;type:TEXT"`
	TargetID       string                      `json:"target_id" xml:"target_id" yaml:"target_id" gorm:"column:target_id;not null;type:VARCHAR(255);index"` // chat_id, channel_id, email, etc.
	TargetName     string                      `json:"target_name" xml:"target_name" yaml:"target_name" gorm:"column:target_name;type:VARCHAR(255)"`
	PlatformConfig t.JSONBImpl                 `json:"platform_config" xml:"platform_config" yaml:"platform_config" gorm:"column:platform_config;type:jsonb"` // Configurações específicas da plataforma
	Response       t.JSONBImpl                 `json:"response" xml:"response" yaml:"response" gorm:"column:response;type:jsonb"`                             // Resposta da API da plataforma
	ErrorMessage   string                      `json:"error_message" xml:"error_message" yaml:"error_message" gorm:"column:error_message;type:TEXT"`
	RetryCount     int                         `json:"retry_count" xml:"retry_count" yaml:"retry_count" gorm:"column:retry_count;default:0"`
	MaxRetries     int                         `json:"max_retries" xml:"max_retries" yaml:"max_retries" gorm:"column:max_retries;default:3"`
	Priority       NotificationRulePriority    `json:"priority" xml:"priority" yaml:"priority" gorm:"column:priority;default:'MEDIUM';type:notification_rule_priority;index"`
	ScheduledFor   *time.Time                  `json:"scheduled_for" xml:"scheduled_for" yaml:"scheduled_for" gorm:"column:scheduled_for;type:timestamp;index"`
	SentAt         *time.Time                  `json:"sent_at" xml:"sent_at" yaml:"sent_at" gorm:"column:sent_at;type:timestamp"`
	DeliveredAt    *time.Time                  `json:"delivered_at" xml:"delivered_at" yaml:"delivered_at" gorm:"column:delivered_at;type:timestamp"`
	ReadAt         *time.Time                  `json:"read_at" xml:"read_at" yaml:"read_at" gorm:"column:read_at;type:timestamp"`
	ExpiresAt      *time.Time                  `json:"expires_at" xml:"expires_at" yaml:"expires_at" gorm:"column:expires_at;type:timestamp;index"`
	Metadata       t.JSONBImpl                 `json:"metadata" xml:"metadata" yaml:"metadata" gorm:"column:metadata;type:jsonb"`
	CreatedAt      time.Time                   `json:"created_at" xml:"created_at" yaml:"created_at" gorm:"column:created_at;default:now();index"`
	UpdatedAt      time.Time                   `json:"updated_at" xml:"updated_at" yaml:"updated_at" gorm:"column:updated_at;default:now()"`
}

// NewNotificationHistoryModel cria uma nova instância de histórico de notificação
func NewNotificationHistoryModel() INotificationHistory {
	return &NotificationHistory{
		PlatformConfig: make(t.JSONBImpl),
		Response:       make(t.JSONBImpl),
		Metadata:       make(t.JSONBImpl),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

// TableName retorna o nome da tabela
func (n *NotificationHistory) TableName() string {
	return "mcp_notification_history"
}

// Getters e Setters

func (n *NotificationHistory) GetID() uuid.UUID                    { return n.ID }
func (n *NotificationHistory) SetID(id uuid.UUID)                  { n.ID = id }
func (n *NotificationHistory) GetRuleID() uuid.UUID                { return n.RuleID }
func (n *NotificationHistory) SetRuleID(ruleID uuid.UUID)          { n.RuleID = ruleID }
func (n *NotificationHistory) GetTemplateID() *uuid.UUID           { return n.TemplateID }
func (n *NotificationHistory) SetTemplateID(templateID *uuid.UUID) { n.TemplateID = templateID }
func (n *NotificationHistory) GetAnalysisJobID() *uuid.UUID        { return n.AnalysisJobID }
func (n *NotificationHistory) SetAnalysisJobID(analysisJobID *uuid.UUID) {
	n.AnalysisJobID = analysisJobID
}
func (n *NotificationHistory) GetPlatform() NotificationHistoryPlatform { return n.Platform }
func (n *NotificationHistory) SetPlatform(platform NotificationHistoryPlatform) {
	n.Platform = platform
}
func (n *NotificationHistory) GetStatus() NotificationHistoryStatus          { return n.Status }
func (n *NotificationHistory) SetStatus(status NotificationHistoryStatus)    { n.Status = status }
func (n *NotificationHistory) GetSubject() string                            { return n.Subject }
func (n *NotificationHistory) SetSubject(subject string)                     { n.Subject = subject }
func (n *NotificationHistory) GetMessage() string                            { return n.Message }
func (n *NotificationHistory) SetMessage(message string)                     { n.Message = message }
func (n *NotificationHistory) GetTargetID() string                           { return n.TargetID }
func (n *NotificationHistory) SetTargetID(targetID string)                   { n.TargetID = targetID }
func (n *NotificationHistory) GetTargetName() string                         { return n.TargetName }
func (n *NotificationHistory) SetTargetName(targetName string)               { n.TargetName = targetName }
func (n *NotificationHistory) GetPlatformConfig() t.JSONBImpl                { return n.PlatformConfig }
func (n *NotificationHistory) SetPlatformConfig(config t.JSONBImpl)          { n.PlatformConfig = config }
func (n *NotificationHistory) GetResponse() t.JSONBImpl                      { return n.Response }
func (n *NotificationHistory) SetResponse(response t.JSONBImpl)              { n.Response = response }
func (n *NotificationHistory) GetErrorMessage() string                       { return n.ErrorMessage }
func (n *NotificationHistory) SetErrorMessage(errorMessage string)           { n.ErrorMessage = errorMessage }
func (n *NotificationHistory) GetRetryCount() int                            { return n.RetryCount }
func (n *NotificationHistory) SetRetryCount(retryCount int)                  { n.RetryCount = retryCount }
func (n *NotificationHistory) GetMaxRetries() int                            { return n.MaxRetries }
func (n *NotificationHistory) SetMaxRetries(maxRetries int)                  { n.MaxRetries = maxRetries }
func (n *NotificationHistory) GetPriority() NotificationRulePriority         { return n.Priority }
func (n *NotificationHistory) SetPriority(priority NotificationRulePriority) { n.Priority = priority }
func (n *NotificationHistory) GetScheduledFor() *time.Time                   { return n.ScheduledFor }
func (n *NotificationHistory) SetScheduledFor(scheduledFor *time.Time)       { n.ScheduledFor = scheduledFor }
func (n *NotificationHistory) GetSentAt() *time.Time                         { return n.SentAt }
func (n *NotificationHistory) SetSentAt(sentAt *time.Time)                   { n.SentAt = sentAt }
func (n *NotificationHistory) GetDeliveredAt() *time.Time                    { return n.DeliveredAt }
func (n *NotificationHistory) SetDeliveredAt(deliveredAt *time.Time)         { n.DeliveredAt = deliveredAt }
func (n *NotificationHistory) GetReadAt() *time.Time                         { return n.ReadAt }
func (n *NotificationHistory) SetReadAt(readAt *time.Time)                   { n.ReadAt = readAt }
func (n *NotificationHistory) GetExpiresAt() *time.Time                      { return n.ExpiresAt }
func (n *NotificationHistory) SetExpiresAt(expiresAt *time.Time)             { n.ExpiresAt = expiresAt }
func (n *NotificationHistory) GetMetadata() t.JSONBImpl                      { return n.Metadata }
func (n *NotificationHistory) SetMetadata(metadata t.JSONBImpl)              { n.Metadata = metadata }
func (n *NotificationHistory) GetCreatedAt() time.Time                       { return n.CreatedAt }
func (n *NotificationHistory) SetCreatedAt(createdAt time.Time)              { n.CreatedAt = createdAt }
func (n *NotificationHistory) GetUpdatedAt() time.Time                       { return n.UpdatedAt }
func (n *NotificationHistory) SetUpdatedAt(updatedAt time.Time)              { n.UpdatedAt = updatedAt }

// Status management methods

func (n *NotificationHistory) Validate() error {
	if n.RuleID == uuid.Nil {
		return fmt.Errorf("rule_id is required")
	}

	if n.Message == "" {
		return fmt.Errorf("message is required")
	}

	if n.TargetID == "" {
		return fmt.Errorf("target_id is required")
	}

	if n.Platform == "" {
		return fmt.Errorf("platform is required")
	}

	if n.RetryCount < 0 {
		return fmt.Errorf("retry_count cannot be negative")
	}

	if n.MaxRetries < 0 {
		return fmt.Errorf("max_retries cannot be negative")
	}

	return nil
}

func (n *NotificationHistory) CanRetry() bool {
	if n.Status == NotificationHistoryStatusSent ||
		n.Status == NotificationHistoryStatusDelivered ||
		n.Status == NotificationHistoryStatusCancelled {
		return false
	}

	return n.RetryCount < n.MaxRetries
}

func (n *NotificationHistory) ShouldExpire() bool {
	if n.ExpiresAt == nil {
		return false
	}

	return time.Now().After(*n.ExpiresAt)
}

func (n *NotificationHistory) GetDeliveryDuration() *time.Duration {
	if n.SentAt == nil || n.DeliveredAt == nil {
		return nil
	}

	duration := n.DeliveredAt.Sub(*n.SentAt)
	return &duration
}

func (n *NotificationHistory) GetDeliveryStatus() string {
	switch n.Status {
	case NotificationHistoryStatusPending:
		return "⏳ Pendente"
	case NotificationHistoryStatusSent:
		return "📤 Enviado"
	case NotificationHistoryStatusDelivered:
		return "✅ Entregue"
	case NotificationHistoryStatusFailed:
		return "❌ Falhou"
	case NotificationHistoryStatusRetrying:
		return "🔄 Tentando novamente"
	case NotificationHistoryStatusCancelled:
		return "🚫 Cancelado"
	case NotificationHistoryStatusRead:
		return "👁️ Lida"
	default:
		return "❓ Desconhecido"
	}
}

func (n *NotificationHistory) MarkAsSent() error {
	if n.Status != NotificationHistoryStatusPending && n.Status != NotificationHistoryStatusRetrying {
		return fmt.Errorf("cannot mark as sent from status: %s", n.Status)
	}

	n.Status = NotificationHistoryStatusSent
	now := time.Now()
	n.SentAt = &now
	n.UpdatedAt = now

	return nil
}

func (n *NotificationHistory) MarkAsDelivered() error {
	if n.Status != NotificationHistoryStatusSent {
		return fmt.Errorf("cannot mark as delivered from status: %s", n.Status)
	}

	n.Status = NotificationHistoryStatusDelivered
	now := time.Now()
	n.DeliveredAt = &now
	n.UpdatedAt = now

	return nil
}

func (n *NotificationHistory) MarkAsFailed(errorMessage string) error {
	n.Status = NotificationHistoryStatusFailed
	n.ErrorMessage = errorMessage
	n.UpdatedAt = time.Now()

	return nil
}

func (n *NotificationHistory) MarkAsRead() error {
	if n.Status != NotificationHistoryStatusDelivered {
		return fmt.Errorf("cannot mark as read from status: %s", n.Status)
	}

	n.Status = NotificationHistoryStatusRead
	now := time.Now()
	n.ReadAt = &now
	n.UpdatedAt = now

	return nil
}

func (n *NotificationHistory) IncrementRetryCount() error {
	if !n.CanRetry() {
		return fmt.Errorf("cannot retry: retry count (%d) >= max retries (%d)", n.RetryCount, n.MaxRetries)
	}

	n.RetryCount++
	n.Status = NotificationHistoryStatusRetrying
	n.UpdatedAt = time.Now()

	return nil
}

// Helper functions para criar notificações específicas

func NewNotificationHistoryFromJob(ruleID uuid.UUID, jobID uuid.UUID, platform NotificationHistoryPlatform, targetID, targetName, subject, message string) INotificationHistory {
	history := NewNotificationHistoryModel().(*NotificationHistory)
	history.RuleID = ruleID
	history.AnalysisJobID = &jobID
	history.Platform = platform
	history.TargetID = targetID
	history.TargetName = targetName
	history.Subject = subject
	history.Message = message
	history.Status = NotificationHistoryStatusPending

	// Set expiration to 24 hours from now
	expiration := time.Now().Add(24 * time.Hour)
	history.ExpiresAt = &expiration

	return history
}

func NewScheduledNotification(ruleID uuid.UUID, platform NotificationHistoryPlatform, targetID, targetName, subject, message string, scheduledFor time.Time) INotificationHistory {
	history := NewNotificationHistoryModel().(*NotificationHistory)
	history.RuleID = ruleID
	history.Platform = platform
	history.TargetID = targetID
	history.TargetName = targetName
	history.Subject = subject
	history.Message = message
	history.Status = NotificationHistoryStatusPending
	history.ScheduledFor = &scheduledFor

	// Set expiration to 48 hours from scheduled time
	expiration := scheduledFor.Add(48 * time.Hour)
	history.ExpiresAt = &expiration

	return history
}
