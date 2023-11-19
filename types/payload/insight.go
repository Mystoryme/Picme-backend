package payload

import (
	"picme-backend/types/enum"
	"time"
)

type InsightResponse struct {
	Id          *uint64       `json:"id"`
	Trigger     *ProfileInfo  `json:"trigger"`
	TriggerId   *uint64       `json:"triggerId"`
	Triggee     *ProfileInfo  `json:"triggee"`
	TriggeeId   *uint64       `json:"triggeeId"`
	InsightType *enum.Insight `json:"type"`
	CreatedAt   *time.Time    `json:"createdAt"` // Embedded field
	UpdatedAt   *time.Time    `json:"updatedAt"` // Embedded field
}

type InsightObject struct {
	views int64 `json:"views"`
	likes int64 `json:"likes"`
}
