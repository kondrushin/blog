package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/server/response"
)

type IBlogUseCase interface {
	GetPost(ctx context.Context, id int64) (*domain.Post, error)
	GetPosts(ctx context.Context) []*domain.Post
	CreatePost(ctx context.Context, p *domain.Post) (int64, error)
	UpdatePost(ctx context.Context, post *domain.Post, id int64) error
	DeletePost(ctx context.Context, id int64) error
}

type Controller struct {
	UseCase IBlogUseCase
}

func (ctr *Controller) GetPost(c *gin.Context) {
	var reqModel postIdRequest

	if err := readPathParameters(c, &reqModel); err != nil {
		c.Error(err)
		return
	}

	post, err := ctr.UseCase.GetPost(c.Request.Context(), reqModel.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (ctr *Controller) GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"posts": ctr.UseCase.GetPosts(c)})
}

func (ctr *Controller) CreatePost(c *gin.Context) {
	var reqModel postRequest
	if err := readJSON(c, &reqModel); err != nil {
		c.Error(err)
		return
	}

	id, err := ctr.UseCase.CreatePost(c.Request.Context(), reqModel.toDomainModel())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Id": id})
}

func (ctr *Controller) UpdatePost(c *gin.Context) {
	var idReqModel postIdRequest
	if err := readPathParameters(c, &idReqModel); err != nil {
		c.Error(err)
		return
	}

	var reqModel postRequest
	if err := readJSON(c, &reqModel); err != nil {
		c.Error(err)
		return
	}
	reqModel.ID = idReqModel.ID

	err := ctr.UseCase.UpdatePost(c.Request.Context(), reqModel.toDomainModel(), reqModel.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Id": reqModel.ID})
}

func (ctr *Controller) DeletePost(c *gin.Context) {
	var reqModel postIdRequest

	if err := readPathParameters(c, &reqModel); err != nil {
		c.Error(err)
		return
	}

	if err := ctr.UseCase.DeletePost(c.Request.Context(), reqModel.ID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"status": "success"})
}

func readPathParameters(c *gin.Context, dst any) error {
	err := c.BindUri(dst)
	if err != nil {
		err = response.SetHttpStatusCode(err, http.StatusBadRequest)
		return err
	}

	return nil
}

func readJSON(c *gin.Context, dst any) error {
	err := c.BindJSON(dst)
	if err != nil {
		err = response.SetHttpStatusCode(err, http.StatusBadRequest)
		return err
	}

	return nil
}

type postRequest struct {
	ID      int64  `json:"-" uri:"id"`
	Author  string `json:"author" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type postIdRequest struct {
	ID int64 `uri:"id"`
}

func (p *postRequest) toDomainModel() *domain.Post {
	return &domain.Post{
		ID:      p.ID,
		Author:  p.Author,
		Title:   p.Title,
		Content: p.Content,
	}
}
