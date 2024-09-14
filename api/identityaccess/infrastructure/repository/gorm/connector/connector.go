package connector

import (
	"gorm.io/gorm"
)

type Connector interface {
	GetDB() *gorm.DB
}
