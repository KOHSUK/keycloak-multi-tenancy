package repository

import (
	"api/identityaccess/domain/model/identity"
	"context"
)

type TenantMemberRepository interface {
	Save(ctx context.Context, tenantId identity.TenantId, members []identity.TenantMember) error
}