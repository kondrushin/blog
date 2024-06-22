package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kondrushin/blog/internal/controller"
)

func RegisterHandlers(r *gin.Engine, blogUseCase controller.IBlogUseCase) {
	s := controller.Controller{UseCase: blogUseCase}

	blogGroup := r.Group("/v1/api/blog")
	{
		blogGroup.GET("/posts/:id", s.GetPost)
		blogGroup.GET("/posts", s.GetPosts)
		blogGroup.POST("/posts", s.CreatePost)
		blogGroup.DELETE("/posts/:id", s.DeletePost)
		blogGroup.PUT("/posts/:id", s.UpdatePost)
	}
}
