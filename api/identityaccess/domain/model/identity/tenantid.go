package identity

import "github.com/google/uuid"

type TenantId struct {
	Value uuid.UUID
}

// create from UUID
func NewTenantId(value uuid.UUID) TenantId {
	return TenantId{Value: value}
}
