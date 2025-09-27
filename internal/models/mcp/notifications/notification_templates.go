package notifications

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/types"
)

// NotificationTemplateType representa o tipo de template de notificação
type NotificationTemplateType string

const (
	NotificationTemplateTypeJobCompleted NotificationTemplateType = "JOB_COMPLETED"
	NotificationTemplateTypeJobFailed    NotificationTemplateType = "JOB_FAILED"
	NotificationTemplateTypeJobStarted   NotificationTemplateType = "JOB_STARTED"
	NotificationTemplateTypeJobRetried   NotificationTemplateType = "JOB_RETRIED"
	NotificationTemplateTypeScoreAlert   NotificationTemplateType = "SCORE_ALERT"
	NotificationTemplateTypeTimeAlert    NotificationTemplateType = "TIME_ALERT"
	NotificationTemplateTypeCustom       NotificationTemplateType = "CUSTOM"
)

// NotificationTemplateFormat representa o formato do template
type NotificationTemplateFormat string

const (
	NotificationTemplateFormatText     NotificationTemplateFormat = "TEXT"
	NotificationTemplateFormatMarkdown NotificationTemplateFormat = "MARKDOWN"
	NotificationTemplateFormatHTML     NotificationTemplateFormat = "HTML"
	NotificationTemplateFormatJSON     NotificationTemplateFormat = "JSON"
)

// NotificationTemplateStatus representa o status do template
type NotificationTemplateStatus string

const (
	NotificationTemplateStatusActive   NotificationTemplateStatus = "ACTIVE"
	NotificationTemplateStatusInactive NotificationTemplateStatus = "INACTIVE"
	NotificationTemplateStatusDraft    NotificationTemplateStatus = "DRAFT"
)

// INotificationTemplate interface para templates de notificação
type INotificationTemplate interface {
	TableName() string
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
	GetTemplateType() NotificationTemplateType
	SetTemplateType(templateType NotificationTemplateType)
	GetFormat() NotificationTemplateFormat
	SetFormat(format NotificationTemplateFormat)
	GetStatus() NotificationTemplateStatus
	SetStatus(status NotificationTemplateStatus)
	GetSubjectTemplate() string
	SetSubjectTemplate(subject string)
	GetBodyTemplate() string
	SetBodyTemplate(body string)
	GetPlatformConfigs() t.JSONB
	SetPlatformConfigs(configs t.JSONB)
	GetVariables() t.JSONB
	SetVariables(variables t.JSONB)
	GetIsDefault() bool
	SetIsDefault(isDefault bool)
	GetLanguage() string
	SetLanguage(language string)
	GetTags() t.JSONB
	SetTags(tags t.JSONB)
	GetCreatedBy() uuid.UUID
	SetCreatedBy(createdBy uuid.UUID)
	GetUpdatedBy() *uuid.UUID
	SetUpdatedBy(updatedBy *uuid.UUID)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)

	// Template processing methods
	Validate() error
	RenderSubject(variables map[string]interface{}) (string, error)
	RenderBody(variables map[string]interface{}) (string, error)
	GetAvailableVariables() []string
	GetPlatformConfig(platform string) map[string]interface{}
	IsCompatibleWithPlatform(platform string) bool
}

// NotificationTemplate implementação do template de notificação
type NotificationTemplate struct {
	ID              uuid.UUID                   `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name            string                      `json:"name" xml:"name" yaml:"name" gorm:"column:name;not null;type:VARCHAR(255)"`
	Description     string                      `json:"description" xml:"description" yaml:"description" gorm:"column:description;type:TEXT"`
	TemplateType    NotificationTemplateType    `json:"template_type" xml:"template_type" yaml:"template_type" gorm:"column:template_type;not null;type:notification_template_type"`
	Format          NotificationTemplateFormat  `json:"format" xml:"format" yaml:"format" gorm:"column:format;default:'TEXT';type:notification_template_format"`
	Status          NotificationTemplateStatus  `json:"status" xml:"status" yaml:"status" gorm:"column:status;default:'ACTIVE';type:notification_template_status"`
	SubjectTemplate string                      `json:"subject_template" xml:"subject_template" yaml:"subject_template" gorm:"column:subject_template;type:TEXT"`
	BodyTemplate    string                      `json:"body_template" xml:"body_template" yaml:"body_template" gorm:"column:body_template;not null;type:TEXT"`
	PlatformConfigs t.JSONB                     `json:"platform_configs" xml:"platform_configs" yaml:"platform_configs" gorm:"column:platform_configs;type:jsonb"` // Configurações específicas por plataforma
	Variables       t.JSONB                     `json:"variables" xml:"variables" yaml:"variables" gorm:"column:variables;type:jsonb"` // Definição de variáveis disponíveis
	IsDefault       bool                        `json:"is_default" xml:"is_default" yaml:"is_default" gorm:"column:is_default;default:false"`
	Language        string                      `json:"language" xml:"language" yaml:"language" gorm:"column:language;default:'pt-BR';type:VARCHAR(10)"`
	Tags            t.JSONB                     `json:"tags" xml:"tags" yaml:"tags" gorm:"column:tags;type:jsonb"` // Tags para organização
	CreatedBy       uuid.UUID                   `json:"created_by" xml:"created_by" yaml:"created_by" gorm:"column:created_by;type:uuid"`
	UpdatedBy       *uuid.UUID                  `json:"updated_by" xml:"updated_by" yaml:"updated_by" gorm:"column:updated_by;type:uuid"`
	CreatedAt       time.Time                   `json:"created_at" xml:"created_at" yaml:"created_at" gorm:"column:created_at;default:now()"`
	UpdatedAt       time.Time                   `json:"updated_at" xml:"updated_at" yaml:"updated_at" gorm:"column:updated_at;default:now()"`
}

// NewNotificationTemplateModel cria uma nova instância de template de notificação
func NewNotificationTemplateModel() INotificationTemplate {
	return &NotificationTemplate{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// TableName retorna o nome da tabela
func (n *NotificationTemplate) TableName() string {
	return "mcp_notification_templates"
}

// Getters e Setters
func (n *NotificationTemplate) GetID() uuid.UUID                                   { return n.ID }
func (n *NotificationTemplate) SetID(id uuid.UUID)                                 { n.ID = id }
func (n *NotificationTemplate) GetName() string                                    { return n.Name }
func (n *NotificationTemplate) SetName(name string)                                { n.Name = name }
func (n *NotificationTemplate) GetDescription() string                             { return n.Description }
func (n *NotificationTemplate) SetDescription(description string)                  { n.Description = description }
func (n *NotificationTemplate) GetTemplateType() NotificationTemplateType          { return n.TemplateType }
func (n *NotificationTemplate) SetTemplateType(templateType NotificationTemplateType) { n.TemplateType = templateType }
func (n *NotificationTemplate) GetFormat() NotificationTemplateFormat              { return n.Format }
func (n *NotificationTemplate) SetFormat(format NotificationTemplateFormat)        { n.Format = format }
func (n *NotificationTemplate) GetStatus() NotificationTemplateStatus              { return n.Status }
func (n *NotificationTemplate) SetStatus(status NotificationTemplateStatus)        { n.Status = status }
func (n *NotificationTemplate) GetSubjectTemplate() string                         { return n.SubjectTemplate }
func (n *NotificationTemplate) SetSubjectTemplate(subject string)                  { n.SubjectTemplate = subject }
func (n *NotificationTemplate) GetBodyTemplate() string                            { return n.BodyTemplate }
func (n *NotificationTemplate) SetBodyTemplate(body string)                        { n.BodyTemplate = body }
func (n *NotificationTemplate) GetPlatformConfigs() t.JSONB                        { return n.PlatformConfigs }
func (n *NotificationTemplate) SetPlatformConfigs(configs t.JSONB)                 { n.PlatformConfigs = configs }
func (n *NotificationTemplate) GetVariables() t.JSONB                              { return n.Variables }
func (n *NotificationTemplate) SetVariables(variables t.JSONB)                     { n.Variables = variables }
func (n *NotificationTemplate) GetIsDefault() bool                                 { return n.IsDefault }
func (n *NotificationTemplate) SetIsDefault(isDefault bool)                        { n.IsDefault = isDefault }
func (n *NotificationTemplate) GetLanguage() string                                { return n.Language }
func (n *NotificationTemplate) SetLanguage(language string)                        { n.Language = language }
func (n *NotificationTemplate) GetTags() t.JSONB                                   { return n.Tags }
func (n *NotificationTemplate) SetTags(tags t.JSONB)                               { n.Tags = tags }
func (n *NotificationTemplate) GetCreatedBy() uuid.UUID                            { return n.CreatedBy }
func (n *NotificationTemplate) SetCreatedBy(createdBy uuid.UUID)                   { n.CreatedBy = createdBy }
func (n *NotificationTemplate) GetUpdatedBy() *uuid.UUID                           { return n.UpdatedBy }
func (n *NotificationTemplate) SetUpdatedBy(updatedBy *uuid.UUID)                  { n.UpdatedBy = updatedBy }
func (n *NotificationTemplate) GetCreatedAt() time.Time                            { return n.CreatedAt }
func (n *NotificationTemplate) SetCreatedAt(createdAt time.Time)                   { n.CreatedAt = createdAt }
func (n *NotificationTemplate) GetUpdatedAt() time.Time                            { return n.UpdatedAt }
func (n *NotificationTemplate) SetUpdatedAt(updatedAt time.Time)                   { n.UpdatedAt = updatedAt }

// Template processing methods
func (n *NotificationTemplate) Validate() error {
	if n.Name == "" {
		return fmt.Errorf("template name is required")
	}

	if n.BodyTemplate == "" {
		return fmt.Errorf("template body is required")
	}

	if n.TemplateType == "" {
		return fmt.Errorf("template type is required")
	}

	if n.CreatedBy == uuid.Nil {
		return fmt.Errorf("created_by is required")
	}

	// Validate template syntax (basic check for variables)
	if err := n.validateTemplateSyntax(); err != nil {
		return fmt.Errorf("template syntax error: %w", err)
	}

	return nil
}

func (n *NotificationTemplate) validateTemplateSyntax() error {
	// Check if template has valid variable syntax
	body := n.BodyTemplate
	if n.SubjectTemplate != "" {
		body = n.SubjectTemplate + " " + body
	}

	// Basic validation for {{ variable }} syntax
	openCount := strings.Count(body, "{{")
	closeCount := strings.Count(body, "}}")

	if openCount != closeCount {
		return fmt.Errorf("mismatched template brackets: %d opening, %d closing", openCount, closeCount)
	}

	return nil
}

func (n *NotificationTemplate) RenderSubject(variables map[string]interface{}) (string, error) {
	if n.SubjectTemplate == "" {
		return "", nil
	}

	return n.renderTemplate(n.SubjectTemplate, variables)
}

func (n *NotificationTemplate) RenderBody(variables map[string]interface{}) (string, error) {
	return n.renderTemplate(n.BodyTemplate, variables)
}

func (n *NotificationTemplate) renderTemplate(template string, variables map[string]interface{}) (string, error) {
	result := template

	// Simple template engine - replace {{ variable }} with values
	for key, value := range variables {
		placeholder := fmt.Sprintf("{{ %s }}", key)
		placeholderAlt := fmt.Sprintf("{{%s}}", key) // Without spaces

		valueStr := n.formatValue(value)
		result = strings.ReplaceAll(result, placeholder, valueStr)
		result = strings.ReplaceAll(result, placeholderAlt, valueStr)
	}

	// Add default variables if not provided
	defaultVars := n.getDefaultVariables()
	for key, value := range defaultVars {
		placeholder := fmt.Sprintf("{{ %s }}", key)
		placeholderAlt := fmt.Sprintf("{{%s}}", key)

		if !strings.Contains(variables[key].(string), key) { // Only if not already provided
			valueStr := n.formatValue(value)
			result = strings.ReplaceAll(result, placeholder, valueStr)
			result = strings.ReplaceAll(result, placeholderAlt, valueStr)
		}
	}

	return result, nil
}

func (n *NotificationTemplate) formatValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case bool:
		if v {
			return "✅"
		}
		return "❌"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (n *NotificationTemplate) getDefaultVariables() map[string]interface{} {
	return map[string]interface{}{
		"timestamp":    time.Now().Format("2006-01-02 15:04:05"),
		"date":         time.Now().Format("2006-01-02"),
		"time":         time.Now().Format("15:04:05"),
		"system_name":  "GDBASE MCP",
		"environment":  "production", // Could be configurable
	}
}

func (n *NotificationTemplate) GetAvailableVariables() []string {
	variables := []string{
		"job_id", "job_type", "job_status", "job_progress",
		"project_id", "user_id", "source_url", "source_type",
		"started_at", "completed_at", "duration", "error_message",
		"score", "retry_count", "max_retries",
		"timestamp", "date", "time", "system_name", "environment",
	}

	// Add variables from template definition
	if n.Variables != nil {
		if varsMap, ok := n.Variables.(map[string]interface{}); ok {
			for key := range varsMap {
				variables = append(variables, key)
			}
		}
	}

	return variables
}

func (n *NotificationTemplate) GetPlatformConfig(platform string) map[string]interface{} {
	if n.PlatformConfigs == nil {
		return map[string]interface{}{}
	}

	configsMap, ok := n.PlatformConfigs.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}

	platformConfig, ok := configsMap[platform].(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}

	return platformConfig
}

func (n *NotificationTemplate) IsCompatibleWithPlatform(platform string) bool {
	config := n.GetPlatformConfig(platform)

	// If no specific config, it's compatible with basic text
	if len(config) == 0 {
		return n.Format == NotificationTemplateFormatText
	}

	// Check format compatibility
	supportedFormats, ok := config["supported_formats"].([]interface{})
	if !ok {
		return true // Default to compatible
	}

	currentFormat := string(n.Format)
	for _, format := range supportedFormats {
		if formatStr, ok := format.(string); ok && formatStr == currentFormat {
			return true
		}
	}

	return false
}

// Predefined template builders
func NewJobCompletedTemplate(name, language string, createdBy uuid.UUID) INotificationTemplate {
	template := NewNotificationTemplateModel().(*NotificationTemplate)
	template.Name = name
	template.TemplateType = NotificationTemplateTypeJobCompleted
	template.Language = language
	template.CreatedBy = createdBy
	template.IsDefault = true

	if language == "pt-BR" {
		template.SubjectTemplate = "✅ Job {{ job_type }} Concluído"
		template.BodyTemplate = `🎉 *Job concluído com sucesso!*

📋 **Detalhes:**
• ID: {{ job_id }}
• Tipo: {{ job_type }}
• Status: {{ job_status }}
• Progresso: {{ job_progress }}%
• Duração: {{ duration }}

🎯 **Score:** {{ score }}
📊 **Projeto:** {{ project_id }}
🔗 **Fonte:** {{ source_url }}

⏰ {{ timestamp }}`
	} else {
		template.SubjectTemplate = "✅ Job {{ job_type }} Completed"
		template.BodyTemplate = `🎉 *Job completed successfully!*

📋 **Details:**
• ID: {{ job_id }}
• Type: {{ job_type }}
• Status: {{ job_status }}
• Progress: {{ job_progress }}%
• Duration: {{ duration }}

🎯 **Score:** {{ score }}
📊 **Project:** {{ project_id }}
🔗 **Source:** {{ source_url }}

⏰ {{ timestamp }}`
	}

	return template
}

func NewJobFailedTemplate(name, language string, createdBy uuid.UUID) INotificationTemplate {
	template := NewNotificationTemplateModel().(*NotificationTemplate)
	template.Name = name
	template.TemplateType = NotificationTemplateTypeJobFailed
	template.Language = language
	template.CreatedBy = createdBy
	template.IsDefault = true

	if language == "pt-BR" {
		template.SubjectTemplate = "❌ Job {{ job_type }} Falhou"
		template.BodyTemplate = `⚠️ *Job falhou durante execução!*

📋 **Detalhes:**
• ID: {{ job_id }}
• Tipo: {{ job_type }}
• Status: {{ job_status }}
• Progresso: {{ job_progress }}%
• Tentativas: {{ retry_count }}/{{ max_retries }}

🚨 **Erro:** {{ error_message }}

📊 **Projeto:** {{ project_id }}
🔗 **Fonte:** {{ source_url }}

⏰ {{ timestamp }}

🔄 O sistema tentará novamente automaticamente.`
	} else {
		template.SubjectTemplate = "❌ Job {{ job_type }} Failed"
		template.BodyTemplate = `⚠️ *Job failed during execution!*

📋 **Details:**
• ID: {{ job_id }}
• Type: {{ job_type }}
• Status: {{ job_status }}
• Progress: {{ job_progress }}%
• Attempts: {{ retry_count }}/{{ max_retries }}

🚨 **Error:** {{ error_message }}

📊 **Project:** {{ project_id }}
🔗 **Source:** {{ source_url }}

⏰ {{ timestamp }}

🔄 System will retry automatically.`
	}

	return template
}

func NewScoreAlertTemplate(name, language string, createdBy uuid.UUID) INotificationTemplate {
	template := NewNotificationTemplateModel().(*NotificationTemplate)
	template.Name = name
	template.TemplateType = NotificationTemplateTypeScoreAlert
	template.Language = language
	template.CreatedBy = createdBy
	template.IsDefault = true

	if language == "pt-BR" {
		template.SubjectTemplate = "⚠️ Alerta de Score - {{ job_type }}"
		template.BodyTemplate = `📊 *Alerta de Score Detectado!*

🎯 **Score:** {{ score }} (abaixo do limite)
📋 **Job:** {{ job_type }}
📋 **ID:** {{ job_id }}

📊 **Projeto:** {{ project_id }}
🔗 **Fonte:** {{ source_url }}

⏰ {{ timestamp }}

🔍 Recomendamos verificar a qualidade do código e dependências.`
	} else {
		template.SubjectTemplate = "⚠️ Score Alert - {{ job_type }}"
		template.BodyTemplate = `📊 *Score Alert Detected!*

🎯 **Score:** {{ score }} (below threshold)
📋 **Job:** {{ job_type }}
📋 **ID:** {{ job_id }}

📊 **Project:** {{ project_id }}
🔗 **Source:** {{ source_url }}

⏰ {{ timestamp }}

🔍 We recommend checking code quality and dependencies.`
	}

	return template
}