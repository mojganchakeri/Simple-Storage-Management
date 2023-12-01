package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckUser(c *gin.Context) {

	// if c.Request.Header.Get("user_id") == "" {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Request does not contain any user-id"})
	// }

	c.Next()
}
