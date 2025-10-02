package models

import (
	w "github.com/kubex-ecosystem/gdbase/internal/models/whatsapp"
	"gorm.io/gorm"
)

type WhatsAppModel = w.WhatsAppModel
type WhatsAppModelInterface = w.IWhatsAppModel
type WhatsAppService = w.IWhatsAppService
type WhatsAppRepo = w.IWhatsAppRepo

// Type aliases for enums

type WhatsAppUserType = w.WhatsAppUserType
type WhatsAppStatus = w.WhatsAppStatus
type WhatsAppIntegrationType = w.WhatsAppIntegrationType
type WhatsAppMessageType = w.WhatsAppMessageType

// Enum constants
const (
	WhatsAppUserTypeBusiness = w.WhatsAppUserTypeBusiness
	WhatsAppUserTypeUser     = w.WhatsAppUserTypeUser
	WhatsAppUserTypeBot      = w.WhatsAppUserTypeBot
	WhatsAppUserTypeSystem   = w.WhatsAppUserTypeSystem

	WhatsAppStatusActive       = w.WhatsAppStatusActive
	WhatsAppStatusInactive     = w.WhatsAppStatusInactive
	WhatsAppStatusDisconnected = w.WhatsAppStatusDisconnected
	WhatsAppStatusError        = w.WhatsAppStatusError
	WhatsAppStatusBlocked      = w.WhatsAppStatusBlocked
	WhatsAppStatusPending      = w.WhatsAppStatusPending

	WhatsAppIntegrationTypeBusinessAPI = w.WhatsAppIntegrationTypeBusinessAPI
	WhatsAppIntegrationTypeCloudAPI    = w.WhatsAppIntegrationTypeCloudAPI
	WhatsAppIntegrationTypeWebhook     = w.WhatsAppIntegrationTypeWebhook
	WhatsAppIntegrationTypeGraph       = w.WhatsAppIntegrationTypeGraph

	WhatsAppMessageTypeText     = w.WhatsAppMessageTypeText
	WhatsAppMessageTypeImage    = w.WhatsAppMessageTypeImage
	WhatsAppMessageTypeVideo    = w.WhatsAppMessageTypeVideo
	WhatsAppMessageTypeAudio    = w.WhatsAppMessageTypeAudio
	WhatsAppMessageTypeDocument = w.WhatsAppMessageTypeDocument
	WhatsAppMessageTypeLocation = w.WhatsAppMessageTypeLocation
	WhatsAppMessageTypeTemplate = w.WhatsAppMessageTypeTemplate
	WhatsAppMessageTypeButton   = w.WhatsAppMessageTypeButton
	WhatsAppMessageTypeList     = w.WhatsAppMessageTypeList
)

func NewWhatsAppService(whatsappRepo WhatsAppRepo) WhatsAppService {
	return w.NewWhatsAppService(whatsappRepo)
}

func NewWhatsAppRepo(db *gorm.DB) WhatsAppRepo {
	return w.NewWhatsAppRepository(db)
}

func NewWhatsAppModel() WhatsAppModelInterface {
	return w.NewWhatsAppModel()
}
