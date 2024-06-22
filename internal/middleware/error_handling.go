package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/response"
)

func HttpErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var errInfo errorInfo
		for _, err := range c.Errors {

			var errorWithCode *response.HttpError
			if errors.As(err, &errorWithCode) {
				errInfo = errorInfo{code: errorWithCode.StatusCode, message: err.Error()}
			} else if errors.Is(err, domain.ErrorPostNotFound) {
				errInfo = errorInfo{code: http.StatusNotFound, message: err.Error()}
			} else {
				errInfo = errorInfo{code: http.StatusInternalServerError, message: err.Error()}
			}
		}

		if errInfo.code != 0 {
			c.JSON(errInfo.code, gin.H{"error": errInfo.message})
		}
	}
}

type errorInfo struct {
	code    int
	message string
}
