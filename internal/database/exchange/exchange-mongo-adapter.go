package exchange

import (
	"context"
	"github.com/kjasuquo/usdngn-exchange/config"
	appmongo "github.com/kjasuquo/usdngn-exchange/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	timeout            = 10
	databaseName       = "payourse"
	exchangeCollection = "exchange"
	accountID          = "account_id"
	notificationStatus = "notification_status"
	createdAt          = "created_at"
	documentID         = "_id"
)

type Adapter struct {
	exchangeCol *mongo.Collection
}

func NewExchangeMongoDatabaseAdapter(config config.Config) (*Adapter, error) {

	db, err := appmongo.NewDriver(appmongo.Config{
		URI:     config.MongoURI,
		Timeout: timeout,
	})
	if err != nil {
		return nil, err
	}

	return &Adapter{
		exchangeCol: db.Database(databaseName).Collection(exchangeCollection),
	}, nil

}

func (a *Adapter) SetNotificationStatus(ctx context.Context) {

}
