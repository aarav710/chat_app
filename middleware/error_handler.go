package middleware

import (
	"chatapp/backend/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		var err error = c.Errors[0]
		switch err.(type) {
		  case *ent.NotFoundError:
			  c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
	}
}