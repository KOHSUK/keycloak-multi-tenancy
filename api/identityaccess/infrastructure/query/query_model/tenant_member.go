package query_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
    ID uuid.UUID
    Name string
}

type User struct {
	gorm.Model
	ID uuid.UUID
	Name string
	TenantId string
	Tenant Tenant `gorm:"foreignKey:TenantId"`
}