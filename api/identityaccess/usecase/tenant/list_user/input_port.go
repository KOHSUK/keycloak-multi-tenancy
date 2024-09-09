package list_user

import "context"

type ListUserInputPort interface {
	ListUser(ctx context.Context, input ListUserInputData) (*ListUserOutputData, error)
}
