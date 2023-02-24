package tests

import (
	"encoding/json"
	"errors"
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

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDB(ctrl)
	mockRate := mock_database.NewMockRates(ctrl)
	h := &controller.Handler{
		DB:     mockDb,
		Rates:  mockRate,
		Config: config.Config{},
	}
	r := server.SetupRouter(h)

	userRequest := models.UserRequest{
		FullName:    "Joseph Asuquo",
		PhoneNumber: "08133477843",
		Email:       "okoasuquo@yahoo.com",
		Password:    "joseph",
	}

	objId, err := primitive.ObjectIDFromHex(utils.ComputeHash(userRequest.Email, ""))
	if err != nil {
		t.Fail()
	}

	salt := "758682606881766088"

	user := &models.User{
		ID:           objId,
		FullName:     "Joseph Asuquo",
		PhoneNumber:  "08133477843",
		Email:        "okoasuquo@yahoo.com",
		PasswordHash: utils.ComputeHash(userRequest.Password, salt),
		Salt:         salt,
		USDBalance:   100,
		NGNBalance:   0,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	login := models.LoginRequest{
		Email:    "okoasuquo@yahoo.com",
		Password: "joseph",
	}

	accToken, err := token.GenerateToken(user.Email, token.AccessTokenValidity)
	if err != nil {
		t.Fail()
	}

	t.Run("successful signup", func(t *testing.T) {
		mockDb.EXPECT().CreateUser(gomock.Any(), userRequest).Return(nil)
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(userRequest)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(string(bytes)))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusCreated, rw.Code)
		assert.Contains(t, rw.Body.String(), "user created")
	})

	t.Run("not successful signup", func(t *testing.T) {
		mockDb.EXPECT().CreateUser(gomock.Any(), userRequest).Return(errors.New("error Exist"))
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(userRequest)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(string(bytes)))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "cannot create user")

	})

	t.Run("successful login", func(t *testing.T) {
		mockDb.EXPECT().GetUserByEmail(gomock.Any(), login.Email).Return(user, nil)
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(login)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(string(bytes)))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), user.Email)
	})

	t.Run("successful user profile", func(t *testing.T) {
		mockDb.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(user, nil)
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/profile", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accToken))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), user.Email)
	})

	t.Run("successful user balance", func(t *testing.T) {
		mockDb.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(user, nil)
		rw := httptest.NewRecorder()
		bytes, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/balances", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accToken))
		r.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "successful")
	})

}
