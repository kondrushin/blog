package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kondrushin/blog/internal/server/middleware"
)

func RegisterHandlers(r *gin.Engine, blogUseCase IBlogUseCase) {
	s := Controller{UseCase: blogUseCase}

	blogGroup := r.Group("/v1/api/blog")
	{
		blogGroup.GET("/posts/:id", s.GetPost)
		blogGroup.GET("/posts", s.GetPosts)
		blogGroup.POST("/posts", s.CreatePost)
		blogGroup.DELETE("/posts/:id", s.DeletePost)
		blogGroup.PUT("/posts/:id", s.UpdatePost)
	}
}

func SetupMiddleware(r *gin.Engine) {
	r.Use(middleware.HttpErrorHandlerMiddleware())
	r.Use(gin.Recovery())
}
