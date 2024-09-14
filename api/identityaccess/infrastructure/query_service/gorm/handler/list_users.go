package query_service

import (
	"api/identityaccess/infrastructure/query_service/gorm/connector"
	"api/identityaccess/usecase/query/query_service"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	ID   uuid.UUID
	Name string
}

type User struct {
	gorm.Model
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	TenantId  uuid.UUID
	Tenant    Tenant `gorm:"foreignKey:TenantId"`
}

type ListUsers struct {
	connector connector.Connector
}

func NewListUsers(connector connector.Connector) *ListUsers {
	return &ListUsers{connector: connector}
}

func (q *ListUsers) Handle(ctx context.Context, request query_service.ListUsersRequest) ([]query_service.UserDto, error) {
	db := q.connector.GetDB()

	var users []User
	result := db.Find(&users, "tenant_id = ?", request.TenantId)
	if result.Error != nil {
		return nil, result.Error
	}

	userDtos := make([]query_service.UserDto, len(users))

	for i, user := range users {
		userDtos[i] = query_service.UserDto{
			Id:       user.ID,
			FistName: user.FirstName,
			LastName: user.LastName,
			TenantId: user.TenantId,
			Email:    user.Email,
		}
	}

	return userDtos, nil
}
