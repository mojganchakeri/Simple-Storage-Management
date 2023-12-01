package middleware

import (
	"net/http"
	"retreival_service/internal"
	"retreival_service/internal/models"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		authToken := internal.ExtractAuthToken(authHeader)
		if authToken != "" {
			if internal.IsWhitelisted(authToken) {
				authorized, _ := internal.IsAuthorized(authToken, secret)
				if authorized {
					userID, err := internal.ExtractIDFromToken(authToken, secret)
					if err != nil {
						c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
						c.Abort()
						return
					}
					c.Set("x-user-id", userID)
					c.Next()
					return
				}
			}

			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Not authorized"})
			c.Abort()
			return
		}

		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Not authorized"})
		c.Abort()
	}
}
