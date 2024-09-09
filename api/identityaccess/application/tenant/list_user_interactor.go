package tenant

import (
	"api/identityaccess/domain/repository"
	"api/identityaccess/usecase/tenant/list_user"
	"context"
)

type ListUserInteractor struct {
	tenantRepository repository.TenantRepository
}

func NewListUserInteractor(tenantRepository repository.TenantRepository) *ListUserInteractor {
	return &ListUserInteractor{tenantRepository: tenantRepository}
}

func (l *ListUserInteractor) ListUsers(ctx context.Context, input list_user.ListUserInputData) (*list_user.ListUserOutputData, error) {
	outputData := &list_user.ListUserOutputData{}

	tenant, err := l.tenantRepository.TenantOfId(ctx, input.TenantId)
	if err != nil {
		return nil, err
	}

	for _, user := range tenant.TenantMembers {
		outputData.Users = append(outputData.Users, list_user.User{
			Id: user.UserId.String(),
			// Name: xxxxxx
		})
	}

	return outputData, nil
}
