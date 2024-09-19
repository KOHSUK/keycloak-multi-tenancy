package factory

import (
	"api/identityaccess/domain/identity/model"

	"github.com/google/uuid"
)

type TenantFactory interface {
	NewTenant(id uuid.UUID, name string) *model.Tenant
}
