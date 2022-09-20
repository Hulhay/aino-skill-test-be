package middleware

import (
	"log"
	"net/http"

	"github.com/Hulhay/aino-skill-test-be/shared"
	"github.com/Hulhay/aino-skill-test-be/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(u usecase.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := shared.BuildErrorResponse("Failed", "No token found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := u.ValidateToken(authHeader)
		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
		} else {
			log.Println(err)
			response := shared.BuildErrorResponse("Token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
