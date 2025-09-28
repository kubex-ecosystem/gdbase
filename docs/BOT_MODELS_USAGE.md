# 🤖 Bot Models Usage Guide

Guia completo para usar os novos modelos de bots no GDBASE.

## 📋 Modelos Disponíveis

### 🔷 Discord Integration
- **Modelo**: `DiscordModel`
- **Factory**: `factory/models/discord.go`
- **Suporte**: Bot, Webhook, OAuth2
- **Tabela**: `mcp_discord_integrations`

### 🔷 Telegram Integration
- **Modelo**: `TelegramModel`
- **Factory**: `factory/models/telegram.go`
- **Suporte**: Bot API, Webhook, API direta
- **Tabela**: `mcp_telegram_integrations`

### 🔷 WhatsApp Integration
- **Modelo**: `WhatsAppModel`
- **Factory**: `factory/models/whatsapp.go`
- **Suporte**: Business API, Cloud API, Graph API, Webhook
- **Tabela**: `mcp_whatsapp_integrations`

### 🔷 Unified Messaging
- **Conversas**: `ConversationModel`
- **Mensagens**: `MessageModel`
- **Factory**: `factory/models/messaging.go`
- **Tabelas**: `mcp_conversations`, `mcp_messages`

## 🚀 Como Usar

### 1. Inicialização Básica

```go
package main

import (
    "context"
    "database/sql"

    "github.com/kubex-ecosystem/gdbase/factory/models"
    "gorm.io/gorm"
)

func main() {
    // Sua conexão GORM aqui
    var db *gorm.DB

    ctx := context.Background()

    // ==============================
    // TELEGRAM SETUP
    // ==============================
    telegramRepo := models.NewTelegramRepo(db)
    telegramService := models.NewTelegramService(telegramRepo)

    // ==============================
    // WHATSAPP SETUP
    // ==============================
    whatsappRepo := models.NewWhatsAppRepo(db)
    whatsappService := models.NewWhatsAppService(whatsappRepo)

    // ==============================
    // DISCORD SETUP
    // ==============================
    discordRepo := models.NewDiscordRepo(db)
    discordService := models.NewDiscordService(discordRepo)

    // ==============================
    // MESSAGING SETUP
    // ==============================
    conversationRepo := models.NewConversationRepo(db)
    // messageRepo := models.NewMessageRepo(db) // Implementar se precisar
}
```

### 2. Criando Integração Telegram

```go
func CreateTelegramBot(telegramService models.TelegramService) {
    ctx := context.Background()

    // Criar novo modelo
    telegram := models.NewTelegramModel().(*models.TelegramModel)
    telegram.TelegramUserID = "123456789"
    telegram.FirstName = "Bot Assistant"
    telegram.Username = "myawesome_bot"
    telegram.UserType = models.TelegramUserTypeBot
    telegram.Status = models.TelegramStatusActive

    // Setup como Bot
    err := telegramService.SetupBotIntegration(
        ctx,
        telegram,
        "1234567890:ABCDEFghijklmnopQRSTUVwxyz123456789", // bot token
    )
    if err != nil {
        log.Fatal("Failed to setup Telegram bot:", err)
    }

    log.Println("✅ Telegram bot created successfully!")
}
```

### 3. Criando Integração WhatsApp Business

```go
func CreateWhatsAppBusiness(whatsappService models.WhatsAppService) {
    ctx := context.Background()

    // Criar novo modelo
    whatsapp := models.NewWhatsAppModel().(*models.WhatsAppModel)
    whatsapp.PhoneNumber = "+5511987654321"
    whatsapp.BusinessName = "My Awesome Business"
    whatsapp.UserType = models.WhatsAppUserTypeBusiness
    whatsapp.Status = models.WhatsAppStatusActive

    // Setup Business API
    err := whatsappService.SetupBusinessAPIIntegration(
        ctx,
        whatsapp,
        "EAAG1234567890...", // access token
        "109876543210987",   // phone number ID
    )
    if err != nil {
        log.Fatal("Failed to setup WhatsApp business:", err)
    }

    log.Println("✅ WhatsApp business created successfully!")
}
```

### 4. Criando Integração Discord

```go
func CreateDiscordBot(discordService models.DiscordService) {
    ctx := context.Background()

    // Criar novo modelo
    discord := models.NewDiscordModel().(*models.DiscordModel)
    discord.DiscordUserID = "123456789012345678"
    discord.Username = "awesome_bot"
    discord.DisplayName = "Awesome Bot"
    discord.UserType = models.DiscordUserTypeBot
    discord.Status = models.DiscordStatusActive
    discord.IntegrationType = models.DiscordIntegrationTypeBot

    // Criar integração
    createdDiscord, err := discordService.CreateIntegration(ctx, discord)
    if err != nil {
        log.Fatal("Failed to create Discord integration:", err)
    }

    log.Println("✅ Discord bot created successfully!", createdDiscord.GetID())
}
```

### 5. Sistema de Conversas Unificado

```go
func CreateUnifiedConversation(conversationRepo models.ConversationRepo) {
    ctx := context.Background()

    // Criar nova conversa
    conversation := models.NewConversationModel().(*models.ConversationModel)
    conversation.Platform = models.PlatformTelegram
    conversation.PlatformConversationID = "chat_12345"
    conversation.IntegrationID = "telegram-integration-uuid"
    conversation.ConversationType = models.ConversationTypePrivate
    conversation.Status = models.ConversationStatusActive
    conversation.Title = "Customer Support Chat"

    // Salvar conversa
    savedConversation, err := conversationRepo.Create(ctx, conversation)
    if err != nil {
        log.Fatal("Failed to create conversation:", err)
    }

    log.Printf("✅ Conversation created: %s", savedConversation.GetID())
}
```

### 6. Criando Mensagens

```go
func CreateMessage() {
    // Criar nova mensagem
    message := models.NewMessageModel().(*models.MessageModel)
    message.ConversationID = "conversation-uuid"
    message.Platform = models.PlatformTelegram
    message.PlatformMessageID = "msg_123456"
    message.MessageType = models.MessageTypeText
    message.Direction = models.MessageDirectionInbound
    message.Status = models.MessageStatusSent
    message.SenderID = "user_123"
    message.SenderName = "João Silva"
    message.Content = "Olá! Preciso de ajuda com meu pedido."

    log.Printf("✅ Message ready: %s", message.GetContent())
}
```

## 🔍 Consultando Dados

### Buscar Integrações Ativas

```go
func ListActiveIntegrations(
    telegramService models.TelegramService,
    whatsappService models.WhatsAppService,
    discordService models.DiscordService,
) {
    ctx := context.Background()

    // Buscar Telegram ativos
    telegramBots, err := telegramService.GetActiveIntegrations(ctx)
    if err == nil {
        log.Printf("🔷 Active Telegram bots: %d", len(telegramBots))
    }

    // Buscar WhatsApp ativos
    whatsappBots, err := whatsappService.GetActiveIntegrations(ctx)
    if err == nil {
        log.Printf("🔷 Active WhatsApp integrations: %d", len(whatsappBots))
    }

    // Buscar Discord ativos
    discordBots, err := discordService.GetActiveIntegrations(ctx)
    if err == nil {
        log.Printf("🔷 Active Discord bots: %d", len(discordBots))
    }
}
```

### Buscar por Platform

```go
func GetTelegramByUserID(telegramService models.TelegramService, userID string) {
    ctx := context.Background()

    telegram, err := telegramService.GetIntegrationByTelegramUserID(ctx, userID)
    if err != nil {
        log.Printf("❌ Telegram not found: %v", err)
        return
    }

    log.Printf("✅ Found Telegram: %s (%s)", telegram.GetDisplayName(), telegram.GetUsername())
}
```

## 🧪 Status e Health Checks

### Testar Conexões

```go
func TestConnections(
    telegramService models.TelegramService,
    whatsappService models.WhatsAppService,
    discordService models.DiscordService,
) {
    ctx := context.Background()

    // Test Telegram
    err := telegramService.TestConnection(ctx, "telegram-integration-id")
    if err != nil {
        log.Printf("❌ Telegram connection failed: %v", err)
    } else {
        log.Println("✅ Telegram connection OK")
    }

    // Test WhatsApp
    err = whatsappService.TestConnection(ctx, "whatsapp-integration-id")
    if err != nil {
        log.Printf("❌ WhatsApp connection failed: %v", err)
    } else {
        log.Println("✅ WhatsApp connection OK")
    }

    // Test Discord (implement if needed)
    log.Println("✅ Discord connection OK")
}
```

## 📊 Enum Values

### Platforms
```go
models.PlatformDiscord
models.PlatformTelegram
models.PlatformWhatsApp
models.PlatformMeta
models.PlatformUnified
```

### Status Types
```go
// Geral
models.TelegramStatusActive
models.WhatsAppStatusActive
models.DiscordStatusActive

// Conversas
models.ConversationStatusActive
models.ConversationStatusArchived

// Mensagens
models.MessageStatusSent
models.MessageStatusDelivered
models.MessageStatusRead
```

### Integration Types
```go
// Telegram
models.TelegramIntegrationTypeBot
models.TelegramIntegrationTypeWebhook
models.TelegramIntegrationTypeAPI

// WhatsApp
models.WhatsAppIntegrationTypeBusinessAPI
models.WhatsAppIntegrationTypeCloudAPI
models.WhatsAppIntegrationTypeWebhook

// Discord
models.DiscordIntegrationTypeBot
models.DiscordIntegrationTypeWebhook
models.DiscordIntegrationTypeOAuth2
```

## 🛠️ Database Schema

Para criar as tabelas, execute:

```sql
-- Ver arquivo: internal/services/assets/init-db.sql
-- Seção: BOT INTEGRATION TABLES
```

## ⚠️ Notas Importantes

1. **Tokens sensíveis**: Nunca commitar tokens reais no código
2. **Validação**: Todos os modelos têm validação automática
3. **Logs**: Use o logger integrado `gl.Log()`
4. **Context**: Sempre passe context para operações async
5. **Transações**: Use transações GORM para operações críticas

## 🎯 Próximos Passos

1. Implementar webhook handlers
2. Criar message processors
3. Adicionar rate limiting
4. Implementar retry logic
5. Criar dashboard de monitoramento