package middleware

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	r.Use(HttpErrorHandlerMiddleware())
	r.Use(gin.Recovery())
}
