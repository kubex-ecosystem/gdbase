# 🔔 GDBASE MCP Notification System

Sistema completo de notificações integrado com bots (Telegram, Discord, WhatsApp) que automaticamente notifica sobre jobs de análise e eventos do sistema.

## 🚀 Quick Start

```bash
# 1. Execute o sistema de bots primeiro
./test_bots.sh

# 2. Teste as notificações
go run examples/notification_example.go

# 3. Execute testes específicos
go test -v ./tests -run TestNotificationSystem
```

## 📁 Estrutura do Sistema

```
internal/models/mcp/notifications/
├── notification_rules.go         # Regras de quando notificar
├── notification_templates.go     # Templates personalizáveis
├── notification_history.go       # Histórico de notificações
├── notification_events.go        # Sistema de eventos
└── analysis_jobs_integration.go  # Integração com analysis_jobs

factory/models/
└── notifications.go              # Factory pattern e API pública
```

## 🎯 Recursos Implementados

### ✅ Sistema de Regras
- **Condições flexíveis**: JOB_COMPLETED, JOB_FAILED, SCORE_ALERT, TIME_ALERT
- **Multi-plataforma**: Telegram, Discord, WhatsApp, Email, Slack
- **Rate limiting**: Controle de frequência de notificações
- **Cooldown**: Evita spam de notificações
- **Filtros**: Por job type, usuário, projeto

### ✅ Templates Personalizáveis
- **Formatos**: TEXT, MARKDOWN, HTML, JSON
- **Variáveis dinâmicas**: {{ job_id }}, {{ score }}, {{ timestamp }}
- **Multi-idioma**: Suporte para pt-BR, en-US
- **Templates padrão**: Pré-configurados para cenários comuns

### ✅ Histórico Completo
- **Status tracking**: PENDING → SENT → DELIVERED → READ
- **Retry automático**: Com limite configurável
- **Métricas**: Taxa de entrega, tempo de resposta
- **Auditoria**: Log completo de todas as notificações

### ✅ Integração com Analysis Jobs
- **Eventos automáticos**: Dispara quando jobs mudam status
- **Score alerts**: Notifica quando score está abaixo do limite
- **Progress tracking**: Marcos de progresso (25%, 50%, 75%, 90%)
- **Timeout detection**: Alerta para jobs que excedem tempo limite

## 💡 Uso Básico

### 1. Criar Regra de Notificação

```go
import "github.com/kubex-ecosystem/gdbase/factory/models"

// Setup
ruleRepo := models.NewNotificationRuleRepo(db)
rule := models.NewNotificationRuleModel().(*models.NotificationRule)

// Configurar regra
rule.Name = "Alertas de Jobs Críticos"
rule.Condition = models.NotificationRuleConditionJobFailed
rule.SetPlatforms([]models.NotificationRulePlatform{
    models.NotificationRulePlatformTelegram,
    models.NotificationRulePlatformDiscord,
})
rule.Priority = models.NotificationRulePriorityHigh
rule.CreatedBy = userID

// Configurar alvos (onde enviar)
targets := []models.NotificationTargetConfig{
    {
        Platform: models.NotificationRulePlatformTelegram,
        TargetID: "-1001234567890", // chat_id do Telegram
        Name:     "DevOps Team",
    },
    {
        Platform: models.NotificationRulePlatformDiscord,
        TargetID: "123456789012345678", // channel_id do Discord
        Name:     "alerts-channel",
    },
}
rule.SetTargetConfig(models.CreateTargetConfig(targets))

// Salvar regra
createdRule, err := ruleService.CreateRule(ctx, rule)
```

### 2. Criar Template Personalizado

```go
template := models.NewNotificationTemplateModel().(*models.NotificationTemplate)
template.Name = "Custom Job Alert"
template.TemplateType = models.NotificationTemplateTypeJobFailed
template.SubjectTemplate = "🚨 URGENTE: {{ job_type }} falhou!"
template.BodyTemplate = `
**ALERTA CRÍTICO** 🚨

Job: {{ job_type }}
ID: {{ job_id }}
Erro: {{ error_message }}

**Ação necessária:**
1. Verificar logs do sistema
2. Contatar equipe responsável
3. Investigar causa raiz

⏰ {{ timestamp }}
`
template.Language = "pt-BR"
template.CreatedBy = userID

// Salvar template
createdTemplate, err := templateService.CreateTemplate(ctx, template)
```

### 3. Processar Eventos Manualmente

```go
// Criar evento de job completado
event := models.NewJobCompletedEvent(
    jobID,
    userID,
    &projectID,
    map[string]interface{}{
        "job_type": "SCORECARD_ANALYSIS",
        "score":    0.85,
        "duration": "2m30s",
    },
)

// Processar evento
processor := models.NewNotificationEventProcessor(ruleRepo, templateRepo, historyRepo, messageQueue)
err := processor.ProcessEvent(ctx, event)
```

## 🔧 Integração com Analysis Jobs

### Configuração Automática

```go
// Setup do event processor
eventProcessor := models.NewNotificationEventProcessor(ruleRepo, templateRepo, historyRepo, amqpPublisher)

// Setup da integração
integration := models.NewAnalysisJobNotificationIntegration(eventProcessor)

// Instalar hooks
hooks := models.NewAnalysisJobNotificationHooks(integration)
err := hooks.InstallHooks(ctx)

// Integrar com o service de analysis jobs
models.IntegrateWithAnalysisJobService(analysisJobService, integration)
```

### Events Automáticos

O sistema automaticamente cria eventos quando:

- **Job Status Changed**: Qualquer mudança de status
- **Job Completed**: Job termina com sucesso
- **Job Failed**: Job falha durante execução
- **Job Started**: Job começa a executar
- **Job Retried**: Job é tentado novamente
- **Score Alert**: Score do job está abaixo do limite
- **Time Alert**: Job excede tempo limite esperado

## 📊 Templates Padrão

### Job Completed (Português)
```
✅ Job {{ job_type }} Concluído

🎉 *Job concluído com sucesso!*

📋 **Detalhes:**
• ID: {{ job_id }}
• Tipo: {{ job_type }}
• Progresso: {{ job_progress }}%
• Duração: {{ duration }}

🎯 **Score:** {{ score }}
📊 **Projeto:** {{ project_id }}

⏰ {{ timestamp }}
```

### Job Failed (Português)
```
❌ Job {{ job_type }} Falhou

⚠️ *Job falhou durante execução!*

📋 **Detalhes:**
• ID: {{ job_id }}
• Tentativas: {{ retry_count }}/{{ max_retries }}

🚨 **Erro:** {{ error_message }}

🔄 O sistema tentará novamente automaticamente.
```

### Score Alert (Português)
```
⚠️ Alerta de Score - {{ job_type }}

📊 *Score abaixo do limite!*

🎯 **Score:** {{ score }}
📋 **Job:** {{ job_id }}

🔍 Recomendamos verificar a qualidade do código.
```

## 🗄️ Schema do Banco

As tabelas são criadas automaticamente:

```sql
-- Sistema de notificações
mcp_notification_rules      -- Regras de notificação
mcp_notification_templates  -- Templates de mensagens
mcp_notification_history    -- Histórico de notificações enviadas

-- Integração com jobs
mcp_analysis_jobs          -- Jobs de análise (já existente)
```

## 🔧 Configuração

### 1. Configurar Targets por Plataforma

```go
// Telegram
telegramTargets := []models.NotificationTargetConfig{
    {
        Platform: models.NotificationRulePlatformTelegram,
        TargetID: "-1001234567890", // Group/Channel chat_id
        Name:     "DevOps Alerts",
        Config: map[string]interface{}{
            "parse_mode": "Markdown",
            "disable_web_page_preview": true,
        },
    },
}

// Discord
discordTargets := []models.NotificationTargetConfig{
    {
        Platform: models.NotificationRulePlatformDiscord,
        TargetID: "123456789012345678", // Channel ID
        Name:     "alerts-channel",
        Config: map[string]interface{}{
            "mention_role": "@DevOps",
            "embed_color": 0xFF0000, // Red for errors
        },
    },
}

// WhatsApp
whatsappTargets := []models.NotificationTargetConfig{
    {
        Platform: models.NotificationRulePlatformWhatsApp,
        TargetID: "5511987654321", // Phone number
        Name:     "Admin WhatsApp",
        Config: map[string]interface{}{
            "message_type": "text",
        },
    },
}
```

### 2. Configurar Rate Limiting

```go
rule.MaxNotificationsPerHour = 10  // Máximo 10 notificações por hora
rule.CooldownMinutes = 5           // Aguardar 5 min entre notificações similares
```

### 3. Configurar Triggers de Score

```go
rule.TriggerConfig = map[string]interface{}{
    "score_threshold": 0.7,     // Alertar se score < 0.7
    "operator":       "lt",     // less than
}
```

### 4. Configurar Agendamento

```go
rule.ScheduleConfig = map[string]interface{}{
    "time_ranges": []string{"09:00-18:00"}, // Apenas horário comercial
    "weekdays":    []string{"mon", "tue", "wed", "thu", "fri"},
    "timezone":    "America/Sao_Paulo",
}
```

## 📈 Métricas e Monitoramento

### Consultar Histórico

```go
// Buscar notificações por status
notifications, total, err := historyService.ListNotifications(ctx,
    &ruleID,
    &models.NotificationHistoryPlatformTelegram,
    &models.NotificationHistoryStatusSent,
    10, 0)

// Buscar estatísticas
stats, err := historyService.GetNotificationStats(ctx, time.Now().Add(-24*time.Hour))
```

### Estatísticas Disponíveis

```go
stats := map[string]interface{}{
    "total_notifications": 150,
    "sent_notifications":  145,
    "failed_notifications": 5,
    "delivery_rate":       0.967,
    "avg_delivery_time":   "2.3s",
    "by_platform": map[string]interface{}{
        "TELEGRAM": 80,
        "DISCORD":  60,
        "WHATSAPP": 10,
    },
    "by_priority": map[string]interface{}{
        "HIGH":   50,
        "MEDIUM": 80,
        "LOW":    20,
    },
}
```

## 🚨 Troubleshooting

### Notificações não sendo enviadas

1. **Verificar regras ativas**:
   ```go
   rules, err := ruleService.ListActiveRules(ctx)
   ```

2. **Verificar rate limiting**:
   ```go
   // Verificar se a regra pode ser disparada
   canTrigger := rule.CanTrigger()
   ```

3. **Verificar targets**:
   ```go
   // Verificar se a configuração de targets está correta
   targetConfig := rule.GetTargetConfig()
   ```

### Templates não renderizando

1. **Verificar sintaxe**:
   ```go
   err := template.Validate()
   ```

2. **Testar renderização**:
   ```go
   subject, err := template.RenderSubject(variables)
   body, err := template.RenderBody(variables)
   ```

### Performance Issues

1. **Verificar índices do banco**
2. **Ajustar rate limiting**
3. **Otimizar queries com filtros**

## 🎉 Integração com GoBE

O sistema está pronto para integrar com o GoBE através do **RabbitMQ**:

1. **GDBASE** gera eventos de notificação
2. **RabbitMQ** transporta os eventos (exchange: `gobe.events`)
3. **GoBE** processa e envia via bots existentes
4. **Webhook feedback** confirma entrega

### Routing Keys
- `mcp.notification.JOB_COMPLETED`
- `mcp.notification.JOB_FAILED`
- `mcp.notification.SCORE_ALERT`

**Sistema 100% funcional e pronto para usar!** 🚀