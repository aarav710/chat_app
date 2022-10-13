package middleware

import (
	"chatapp/backend/ent"
	customErrors "chatapp/backend/errors"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		var err error = c.Errors[0]
		if errors.Is(err, customErrors.UnauthorizedError) {
            c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}
		switch err.(type) {
		case customErrors.IncorrectQueryParameterError:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
        case *ent.NotFoundError:
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		
	}
}
