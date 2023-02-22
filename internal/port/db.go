package port

import (
	"context"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
)

type DB interface {
	CreateUser(ctx context.Context, userRequest models.UserRequest) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	ComputeHash(password, salt string) string
}

type Rates interface {
}
