package query_service

import (
	"api/identityaccess/infrastructure/query/gorm/connector"
	"api/identityaccess/infrastructure/query/gorm/query_model"
	"context"
)

type ListTenantMembersQueryService struct {
	connector connector.Connector
}

func NewListTenantMembersQueryService(connector connector.Connector) *ListTenantMembersQueryService {
	return &ListTenantMembersQueryService{connector: connector}
}

func (q *ListTenantMembersQueryService) ListTenantMembers(ctx context.Context, tenantId string) ([]query_model.User, error) {
	db := q.connector.GetDB()

	var users []query_model.User
	result := db.Find(&users, "tenant_id = ?", tenantId)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
