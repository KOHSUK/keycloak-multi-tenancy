package query

import (
	"context"

	"github.com/google/uuid"
)

type ListUserInputData struct {
	TenantId uuid.UUID
}

type User struct {
	Id   string
	Name string
}

type ListUserOutputData struct {
	Users []User
}

type ListUserUseCase interface {
	Handle(ctx context.Context, input ListUserInputData) (*ListUserOutputData, error)
}
