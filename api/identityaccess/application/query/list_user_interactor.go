package tenant

import (
	"api/identityaccess/usecase/query"
	"api/identityaccess/usecase/query/query_service"
	"context"
)

type ListUserInteractor struct {
	queryService query_service.ListUsersQueryService
}

func NewListUserInteractor(queryService query_service.ListUsersQueryService) *ListUserInteractor {
	return &ListUserInteractor{queryService: queryService}
}

func (l *ListUserInteractor) Handle(ctx context.Context, input query.ListUserInputData) (*query.ListUserOutputData, error) {
	users, err := l.queryService.Handle(ctx, query_service.ListUsersRequest{
		TenantId: input.TenantId,
	})

	if err != nil {
		return nil, err
	}

	outputData := &query.ListUserOutputData{
		Users: make([]query.User, len(users)),
	}

	return outputData, nil
}
