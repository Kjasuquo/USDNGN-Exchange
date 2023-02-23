package port

import (
	"context"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
)

type DB interface {
	CreateUser(ctx context.Context, userRequest models.UserRequest) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateBalances(ctx context.Context, email string, usdBalance, ngnBalance float64) error
	CreateTransaction(ctx context.Context, transaction models.Transactions) error
	GetTransaction(ctx context.Context, email string) ([]models.Transactions, error)
	ComputeHash(password, salt string) string
}

type Rates interface {
	GetRates(ctx context.Context) (models.Rates, error)
}
