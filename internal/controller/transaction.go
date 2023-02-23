package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
	"github.com/kjasuquo/usdngn-exchange/internal/services/web"
	"net/http"
)

func (h *Handler) CustomerSellUSDForNGN(c *gin.Context) {
	user, err := h.GetUserFromContext(c)
	if err != nil {
		web.JSON(c, "invalid access token", http.StatusUnauthorized, nil, errors.New("invalid access_token"))
		return
	}

	request := models.TransactionRequest{}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		web.JSON(c, "bad request", http.StatusBadRequest, nil, err)
		return
	}

	if request.Amount > user.USDBalance {
		web.JSON(c, "insufficient funds", http.StatusBadRequest, nil, err)
		return
	}

	USDBal := user.USDBalance - request.Amount

	rates, err := h.Rates.GetRates(c)
	if err != nil {
		web.JSON(c, "cannot get rates", http.StatusInternalServerError, nil, err)
		return
	}

	NGNBal := (rates.Data.Rates.SellUSD.Rate * request.Amount) + user.NGNBalance

	err = h.DB.UpdateBalances(c, user.Email, USDBal, NGNBal)
	if err != nil {
		web.JSON(c, "cannot update balances", http.StatusInternalServerError, nil, err)
		return
	}

	transaction := models.Transactions{
		Type:             models.SellUSD,
		UserId:           user.ID,
		Rate:             rates.Data.Rates.SellUSD.Rate,
		RequestCurrency:  "USD",
		RequestAmount:    request.Amount,
		ReceivedCurrency: "NGN",
		ReceivedAmount:   rates.Data.Rates.SellUSD.Rate * request.Amount,
	}

	err = h.DB.CreateTransaction(c, transaction)
	if err != nil {
		web.JSON(c, "cannot create transaction", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "transaction successful", http.StatusOK, nil, nil)

}

func (h *Handler) CustomerBuyUSDWithNGN(c *gin.Context) {
	user, err := h.GetUserFromContext(c)
	if err != nil {
		web.JSON(c, "invalid access token", http.StatusUnauthorized, nil, errors.New("invalid access_token"))
		return
	}

	request := models.TransactionRequest{}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		web.JSON(c, "bad request", http.StatusBadRequest, nil, err)
		return
	}

	if request.Amount > user.NGNBalance {
		web.JSON(c, "insufficient funds", http.StatusBadRequest, nil, err)
		return
	}

	NGNBal := user.NGNBalance - request.Amount

	rates, err := h.Rates.GetRates(c)
	if err != nil {
		web.JSON(c, "cannot get rates", http.StatusInternalServerError, nil, err)
		return
	}

	USDBal := (request.Amount / rates.Data.Rates.BuyUSD.Rate) + user.USDBalance

	err = h.DB.UpdateBalances(c, user.Email, USDBal, NGNBal)
	if err != nil {
		web.JSON(c, "cannot update balances", http.StatusInternalServerError, nil, err)
		return
	}

	transaction := models.Transactions{
		Type:             models.BuyUSD,
		UserId:           user.ID,
		Rate:             rates.Data.Rates.BuyUSD.Rate,
		RequestCurrency:  "NGN",
		RequestAmount:    request.Amount,
		ReceivedCurrency: "USD",
		ReceivedAmount:   request.Amount / rates.Data.Rates.BuyUSD.Rate,
	}

	err = h.DB.CreateTransaction(c, transaction)
	if err != nil {
		web.JSON(c, "cannot create transaction", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "transaction successful", http.StatusOK, nil, nil)

}

func (h *Handler) GetTransaction(c *gin.Context) {
	user, err := h.GetUserFromContext(c)
	if err != nil {
		web.JSON(c, "invalid access token", http.StatusUnauthorized, nil, errors.New("invalid access_token"))
		return
	}

	transactions, err := h.DB.GetTransaction(c, user.Email)
	if err != nil {
		web.JSON(c, "cannot get transactions", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "successfully found", http.StatusOK, transactions, nil)
}
