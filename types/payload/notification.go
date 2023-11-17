package payload

import (
	"picme-backend/types/enum"
	"time"
)

type NotificationResponse struct {
	Id               *uint64            `json:"id"`
	Trigger          *ProfileInfo       `json:"trigger"`
	TriggerId        *uint64            `json:"triggerId"`
	Triggee          *ProfileInfo       `json:"triggee"`
	TriggeeId        *uint64            `json:"triggeeId"`
	Post             *PostResponse      `json:"post"`
	PostId           *uint64            `json:"postId"`
	NotificationType *enum.Notification `json:"type"`
	CreatedAt        *time.Time         `json:"createdAt"` // Embedded field
	UpdatedAt        *time.Time         `json:"updatedAt"` // Embedded field
}
