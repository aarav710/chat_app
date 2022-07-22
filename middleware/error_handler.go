package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"chatapp/backend/errors"
)


func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		var err error = c.Errors[0]
		switch err.(type) {
		  case errors.NotFoundError:
			  c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
	}
}