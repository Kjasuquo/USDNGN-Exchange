package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	SellUSD = "sell_usd"
	BuyUSD  = "buy_usd"
)

type Rates struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Rates struct {
			SellUSD struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDCNGN"`
			BuyUSD struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDCNGN_"`
		} `json:"rates"`
	} `json:"data"`
}

type TransactionRequest struct {
	Amount float64 `json:"amount"`
}

type Transactions struct {
	ID               primitive.ObjectID `bson:"_id"`
	UserId           primitive.ObjectID `bson:"user_id"`
	Type             string             `bson:"type"`
	Rate             float64            `bson:"rate"`
	RequestCurrency  string             `bson:"request_currency"`
	RequestAmount    float64            `bson:"request_amount"`
	ReceivedCurrency string             `bson:"received_currency"`
	ReceivedAmount   float64            `bson:"received_amount"`
	CreatedAt        primitive.DateTime `bson:"created_at"`
}
