package repository

import (
	"api/identityaccess/domain/identity/model"
	"context"
)

type TenantMemberRepository interface {
	Save(ctx context.Context, tenantId model.TenantId, members []model.TenantMember) error
}
