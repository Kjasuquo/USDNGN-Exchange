package exchange

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/kjasuquo/usdngn-exchange/config"
	"github.com/kjasuquo/usdngn-exchange/internal/database"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
)

const (
	timeout               = 10
	databaseName          = "exchange"
	userCollection        = "users"
	transactionCollection = "transactions"
)

type Adapter struct {
	userCol        *mongo.Collection
	transactionCol *mongo.Collection
}

func NewExchangeMongoDatabaseAdapter(config config.Config) (*Adapter, error) {

	db, err := database.NewDriver(database.Config{
		URI:     config.MongoURI,
		Timeout: timeout,
	})
	if err != nil {
		return nil, err
	}

	return &Adapter{
		userCol:        db.Database(databaseName).Collection(userCollection),
		transactionCol: db.Database(databaseName).Collection(transactionCollection),
	}, nil

}

func (a *Adapter) CreateUser(ctx context.Context, userRequest models.UserRequest) error {
	salt := generateSalt()

	objId, err := primitive.ObjectIDFromHex(a.ComputeHash(userRequest.Email, ""))
	if err != nil {
		return err
	}

	user := models.User{
		ID:           objId,
		FullName:     userRequest.FullName,
		PhoneNumber:  userRequest.PhoneNumber,
		Email:        userRequest.Email,
		PasswordHash: a.ComputeHash(userRequest.Password, salt),
		Salt:         salt,
		USDBalance:   100,
		NGNBalance:   0,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	_, err = a.userCol.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	objId, err := primitive.ObjectIDFromHex(a.ComputeHash(email, ""))
	if err != nil {
		return nil, err
	}

	var user models.User
	err = a.userCol.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *Adapter) ComputeHash(password, salt string) string {
	hasher := sha512.New()
	// TODO: we should throw this error
	_, _ = hasher.Write([]byte(password + salt))
	result := hex.EncodeToString(hasher.Sum(nil))
	return result[:24]
}

func generateSalt() string {
	rand.Seed(time.Now().Unix())
	result := ""

	for i := 0; i <= 8; i++ {
		result += fmt.Sprint('0' + rand.Intn(41))
	}
	return result
}
