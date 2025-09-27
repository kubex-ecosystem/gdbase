# 🤖 GDBASE Bot Integration System

Sistema completo de integração com bots para Discord, Telegram e WhatsApp/Meta.

## 🚀 Quick Start

```bash
# 1. Execute os testes
./test_bots.sh

# 2. Rode o exemplo
go run examples/bot_models_example.go

# 3. Execute testes específicos
go test -v ./tests -run TestBotModelsQuickStart
```

## 📁 Estrutura dos Modelos

```
internal/models/
├── telegram/           # Telegram Bot API
├── whatsapp/          # WhatsApp Business API
├── discord/           # Discord Bot API
└── messaging/         # Sistema unificado

factory/models/
├── telegram.go        # Factory Telegram
├── whatsapp.go        # Factory WhatsApp
├── discord.go         # Factory Discord
└── messaging.go       # Factory Messaging
```

## 🎯 Recursos Implementados

### ✅ Telegram Integration
- Bot API, Webhook, API direta
- Chats, grupos, canais
- Gerenciamento de permissões
- Tokens e autenticação

### ✅ WhatsApp Integration
- Business API, Cloud API
- Graph API, Webhooks
- Multi-tipos de mensagem
- Configurações de negócio

### ✅ Discord Integration
- Bot, Webhook, OAuth2
- Servidores e canais
- Permissões e roles
- Integrações customizadas

### ✅ Unified Messaging
- Conversas cross-platform
- Mensagens com attachments
- Status tracking
- Threading e replies

## 💡 Uso Básico

```go
import "github.com/kubex-ecosystem/gdbase/factory/models"

// Setup Telegram
telegramRepo := models.NewTelegramRepo(db)
telegramService := models.NewTelegramService(telegramRepo)

// Criar bot
telegram := models.NewTelegramModel().(*models.TelegramModel)
telegram.TelegramUserID = "123456789"
telegram.FirstName = "My Bot"

err := telegramService.SetupBotIntegration(ctx, telegram, "bot_token")
```

## 📚 Documentação

- **[📖 Guia Completo](docs/BOT_MODELS_USAGE.md)** - Documentação detalhada
- **[🧪 Testes](tests/)** - Suíte de testes completa
- **[📝 Exemplo](examples/bot_models_example.go)** - Exemplo prático

## 🗄️ Database Schema

As tabelas são criadas automaticamente:

```sql
-- Integrações
mcp_telegram_integrations
mcp_whatsapp_integrations
mcp_discord_integrations

-- Sistema unificado
mcp_conversations
mcp_messages

-- MCP expandido
mcp_analysis_jobs
```

## 🔧 Comandos Úteis

```bash
# Rodar todos os testes
go test ./tests -v

# Teste específico
go test ./tests -run TestTelegramModelIntegration -v

# Compilar
go build -o gdbase .

# Migrar banco (automático no GORM)
# As tabelas são criadas automaticamente
```

## 🚨 Importante

1. **Tokens**: Nunca commitar tokens reais
2. **Database**: Configure sua conexão GORM
3. **Context**: Sempre use context.Context
4. **Validação**: Modelos têm validação automática
5. **Logger**: Use o logger integrado `gl.Log()`

## 🎉 Status

- ✅ **Modelos**: Telegram, WhatsApp, Discord
- ✅ **Messaging**: Sistema unificado
- ✅ **Factory**: Exposição via factory
- ✅ **Tests**: Suíte completa
- ✅ **Docs**: Documentação detalhada
- ✅ **SQL**: Schema do banco
- ✅ **Examples**: Exemplos práticos

**Tudo pronto para usar!** 🚀