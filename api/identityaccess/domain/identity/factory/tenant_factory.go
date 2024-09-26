package factory

import (
	"api/identityaccess/domain/identity/model"
)

type TenantFactory interface {
	NewTenant(id model.TenantId, name string) (*model.Tenant, error)
}
