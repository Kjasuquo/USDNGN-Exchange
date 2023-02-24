package tests

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kjasuquo/usdngn-exchange/cmd/server"
	"github.com/kjasuquo/usdngn-exchange/config"
	"github.com/kjasuquo/usdngn-exchange/internal/controller"
	mock_database "github.com/kjasuquo/usdngn-exchange/internal/database/mocks"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
	token "github.com/kjasuquo/usdngn-exchange/internal/services/jwt"
	"github.com/kjasuquo/usdngn-exchange/internal/services/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDB(ctrl)
	mockRate := mock_database.NewMockRates(ctrl)
	h := &controller.Handler{
		DB:     mockDb,
		Rates:  mockRate,
		Config: config.Config{},
	}
	r := server.SetupRouter(h)

	login := models.LoginRequest{
		Email:    "okoasuquo@yahoo.com",
		Password: "joseph",
	}

	objId, err := primitive.ObjectIDFromHex(utils.ComputeHash(login.Email, ""))
	if err != nil {
		t.Fail()
	}

	salt := "758682606881766088"

	user := &models.User{
		ID:           objId,
		FullName:     "Joseph Asuquo",
		PhoneNumber:  "08133477843",
		Email:        "okoasuquo@yahoo.com",
		PasswordHash: utils.ComputeHash(login.Password, salt),
		Salt:         salt,
		USDBalance:   100,
		NGNBalance:   0,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	rate := models.Rates{
		Success: true,
		Message: "Current rates",
	}

	request := models.TransactionRequest{Amount: 49}

	sellUSD := struct {
		Rate float64 `json:"rate"`
		Key  string  `json:"key"`
	}{
		Rate: 738.39,
		Key:  "738.39-USDCNGN-d7516d12e151-1677219252394",
	}

	buyUSD := struct {
		Rate float64 `json:"rate"`
		Key  string  `json:"key"`
	}{
		Rate: 758.39,
		Key:  "758.39-USDCNGN_-51f2658cb793-1677219252394",
	}

	accToken, err := token.GenerateToken(user.Email, token.AccessTokenValidity)
	if err != nil {
		t.Fail()
	}

	transactionSell := models.Transactions{
		Type:             models.SellUSD,
		UserId:           user.ID,
		Rate:             sellUSD.Rate,
		RequestCurrency:  "USD",
		RequestAmount:    request.Amount,
		ReceivedCurrency: "NGN",
		ReceivedAmount:   sellUSD.Rate * request.Amount,
	}

	transactionBuy := models.Transactions{
		Type:             models.BuyUSD,
		UserId:           user.ID,
		Rate:             buyUSD.Rate,
		RequestCurrency:  "NGN",
		RequestAmount:    request.Amount,
		ReceivedCurrency: "USD",
		ReceivedAmount:   request.Amount / buyUSD.Rate,
	}

	transactions := []models.Transactions{transactionSell, transactionBuy}

	t.Run("successful CustomerSellUSDForNGN", func(t *testing.T) {
		mockDb.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(user, nil)
		mockRate.EXPECT().GetRates(gomock.Any()).Return(rate, nil)
		mockDb.EXPECT().UpdateBalances(gomock.Any(), user.Email, gomock.Any(), gomock.Any()).Return(nil)
		mockDb.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(nil)
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/transaction/sellUSD", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accToken))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "transaction successful")
	})

	t.Run("successful CustomerBuyUSDWithNGN", func(t *testing.T) {
		mockDb.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(user, nil)
		mockRate.EXPECT().GetRates(gomock.Any()).Return(rate, nil)
		mockDb.EXPECT().UpdateBalances(gomock.Any(), user.Email, gomock.Any(), gomock.Any()).Return(nil)
		mockDb.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(nil)
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/transaction/buyUSD", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accToken))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "transaction successful")
	})

	t.Run("successful CustomerBuyUSDWithNGN", func(t *testing.T) {
		mockDb.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(user, nil)
		mockRate.EXPECT().GetRates(gomock.Any()).Return(rate, nil)
		mockDb.EXPECT().UpdateBalances(gomock.Any(), user.Email, gomock.Any(), gomock.Any()).Return(nil)
		mockDb.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(nil)
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/transaction/buyUSD", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accToken))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "transaction successful")
	})

	t.Run("successful GetTransaction", func(t *testing.T) {
		mockDb.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(user, nil)
		mockDb.EXPECT().GetTransaction(gomock.Any(), user.Email).Return(transactions, nil)

		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/transaction/transactions", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accToken))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "successfully found")
	})

}
