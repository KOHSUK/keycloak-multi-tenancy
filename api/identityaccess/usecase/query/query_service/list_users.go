package query_service

import (
	"context"

	"github.com/google/uuid"
)

type ListUsersRequest struct {
	TenantId uuid.UUID
}

type UserDto struct {
	Id       uuid.UUID
	FistName string
	LastName string
	Email    string
	TenantId uuid.UUID
}

type ListUsersQueryService interface {
	Handle(ctx context.Context, request ListUsersRequest) ([]UserDto, error)
}
