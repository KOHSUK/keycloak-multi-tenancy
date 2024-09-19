package user

import (
	"api/identityaccess/infrastructure/gorm/repository/person"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	TenantId uuid.UUID `gorm:"type:uuid;"`
	Person   person.Person
	Active   bool
}
