package middleware

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt"
	"github.com/kjasuquo/usdngn-exchange/internal/controller"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
	"github.com/kjasuquo/usdngn-exchange/internal/services/jwt"
	"github.com/kjasuquo/usdngn-exchange/internal/services/web"
	"net/http"
	"os"
)

// AuthorizeUser authorizes a request
func AuthorizeUser(s *controller.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {

		secretKey := os.Getenv("JWT_SECRET")
		accessToken := getTokenFromHeader(c)
		accessClaims, err := jwt.ValidateAndGetClaims(accessToken, secretKey)
		if err != nil {
			respondAndAbort(c, "", http.StatusUnauthorized, nil, web.New("unauthorized", http.StatusUnauthorized))
			return
		}

		email, ok := accessClaims["email"].(string)
		if !ok {
			respondAndAbort(c, "", http.StatusInternalServerError, nil, web.New("internal server error", http.StatusInternalServerError))
			return
		}

		var user *models.User
		if user, err = s.DB.GetUserByEmail(c, email); err != nil {
			respondAndAbort(c, "bad request", http.StatusNotFound, nil, web.New(err.Error(), http.StatusNotFound))
			return
		}

		c.Set("access_token", accessToken)
		c.Set("user", user)

		c.Next()
	}
}

// respondAndAbort calls response.JSON and aborts the Context
func respondAndAbort(c *gin.Context, message string, status int, data interface{}, e *web.ErrorResponse) {
	web.JSON(c, message, status, data, e)
	c.Abort()
}

// getTokenFromHeader returns the token string in the authorization header
func getTokenFromHeader(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) > 8 {
		return authHeader[7:]
	}
	return ""
}
