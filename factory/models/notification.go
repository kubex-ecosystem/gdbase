package models

import (
	m "github.com/kubex-ecosystem/gdbase/internal/models/notification"
	"gorm.io/gorm"
)

type MemoryNotificationRepo = m.MemoryNotificationRepo
type NotificationModel = m.Notification
type NotificationService = m.NotificationService
type NotificationRepo = m.NotificationRepo

func NewNotificationService(notificationRepo NotificationRepo) NotificationService {
	return *m.NewNotificationService(notificationRepo)
}

func NewNotificationRepo(db *gorm.DB) NotificationRepo {
	return m.NewMemoryNotificationRepo()
}
