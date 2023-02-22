package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	StatusCode int64  `json:"statusCode"`
	Message    string `json:"message"`
	ActivityId string `json:"activityId"`
}

// Error implements an error interface.
func (e ErrorResponse) Error() string {
	return strconv.FormatInt(e.StatusCode, 10) + ":" + e.Message
}
func JSON(c *gin.Context, message string, status int, data interface{}, err error) {
	errMessage := ""
	if err != nil {
		errMessage = err.Error()
	}
	responsedata := gin.H{
		"message": message,
		"data":    data,
		"errors":  errMessage,
		"status":  http.StatusText(status),
	}

	c.JSON(status, responsedata)
}

func New(message string, status int) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: int64(status),
	}
}
