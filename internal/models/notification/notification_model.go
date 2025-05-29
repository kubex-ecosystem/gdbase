package notification

import (
	"time"
)

type INotification interface {
	TableName() string
	GetID() int
	GetTitle() string
	GetMessage() string
	GetScheduledAt() time.Time
	GetChannel() string
	GetStatus() string
}

// Model: Notification define a estrutura básica de uma notificação.
type Notification struct {
	ID          int       `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title       string    `json:"title" xml:"title" yaml:"title" gorm:"column:title"`
	Message     string    `json:"message" xml:"message" yaml:"message" gorm:"column:message"`
	ScheduledAt time.Time `json:"scheduled_at" xml:"scheduled_at" yaml:"scheduled_at" gorm:"column:scheduled_at"`
	Channel     string    `json:"channel" xml:"channel" yaml:"channel" gorm:"column:channel"`
	Status      string    `json:"status" xml:"status" yaml:"status" gorm:"column:status"`
}

func (n *Notification) TableName() string                    { return "notifications" }
func (n *Notification) GetID() int                           { return n.ID }
func (n *Notification) GetTitle() string                     { return n.Title }
func (n *Notification) GetMessage() string                   { return n.Message }
func (n *Notification) GetScheduledAt() time.Time            { return n.ScheduledAt }
func (n *Notification) GetChannel() string                   { return n.Channel }
func (n *Notification) GetStatus() string                    { return n.Status }
func (n *Notification) SetID(id int)                         { n.ID = id }
func (n *Notification) SetTitle(title string)                { n.Title = title }
func (n *Notification) SetMessage(message string)            { n.Message = message }
func (n *Notification) SetScheduledAt(scheduledAt time.Time) { n.ScheduledAt = scheduledAt }
func (n *Notification) SetChannel(channel string)            { n.Channel = channel }
func (n *Notification) SetStatus(status string)              { n.Status = status }
