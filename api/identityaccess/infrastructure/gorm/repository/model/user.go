package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	TenantId uuid.UUID `gorm:"type:uuid;"`
	Person   Person
}
