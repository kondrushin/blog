package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/server"
	"github.com/kondrushin/blog/internal/server/mocks"
	"github.com/stretchr/testify/mock"
)

func SetupServer(t *testing.T, useCase *mocks.IBlogUseCase) *httpexpect.Expect {
	gin.SetMode(gin.TestMode)
	ginRouter := gin.Default()
	server.SetupMiddleware(ginRouter)

	server.RegisterHandlers(ginRouter, useCase)
	server := httptest.NewServer(ginRouter)
	expect := httpexpect.Default(t, server.URL)

	return expect
}

func Test_GetPost_ShouldReturnPost(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	blogUseCaseMock.
		On("GetPost", mock.Anything, int64(1)).
		Return(&domain.Post{
			ID:      int64(1),
			Author:  "Anton",
			Title:   "Big post",
			Content: "something",
		}, nil)

	expect.GET("/v1/api/blog/posts/1").
		Expect().
		Status(http.StatusOK).
		Body().IsEqual("{\"ID\":1,\"Author\":\"Anton\",\"Title\":\"Big post\",\"Content\":\"something\"}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_GetPost_InvalidId_ShouldReturnBadRequestStatus(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	expect.GET("/v1/api/blog/posts/a").
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_GetPost_NoPost_ShouldReturnNotFountStatus(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	blogUseCaseMock.
		On("GetPost", mock.Anything, int64(1)).
		Return(nil, domain.ErrorPostNotFound)

	expect.GET("/v1/api/blog/posts/1").
		Expect().
		Status(http.StatusNotFound)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_GetPost_Error_ShouldReturn500Status(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	blogUseCaseMock.
		On("GetPost", mock.Anything, int64(1)).
		Return(nil, errors.New("DB error"))

	expect.GET("/v1/api/blog/posts/1").
		Expect().
		Status(http.StatusInternalServerError).
		Body().IsEqual("{\"error\":\"DB error\"}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_GetPosts_ShouldReturnPosts(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	post1 := &domain.Post{
		ID:      int64(1),
		Author:  "Anton",
		Title:   "Big post",
		Content: "something",
	}

	post2 := &domain.Post{
		ID:      int64(2),
		Author:  "Jonny",
		Title:   "Another post",
		Content: "something but different",
	}

	blogUseCaseMock.
		On("GetPosts", mock.Anything).
		Return([]*domain.Post{post1, post2})

	expect.GET("/v1/api/blog/posts").
		Expect().
		Status(http.StatusOK).
		Body().IsEqual("{\"posts\":[{\"ID\":1,\"Author\":\"Anton\",\"Title\":\"Big post\",\"Content\":\"something\"},{\"ID\":2,\"Author\":\"Jonny\",\"Title\":\"Another post\",\"Content\":\"something but different\"}]}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_CreatePost_ShouldReturnPost(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	post := domain.Post{
		Author:  "Anton",
		Title:   "Big post",
		Content: "something",
	}

	blogUseCaseMock.
		On("CreatePost", mock.Anything, &post).
		Return(int64(1), nil)

	expect.POST("/v1/api/blog/posts").
		WithJSON(post).
		Expect().
		Status(http.StatusCreated).
		Body().IsEqual("{\"Id\":1}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_CreatePost_Error_ShouldReturn500Status(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	post := domain.Post{
		Author:  "Anton",
		Title:   "Big post",
		Content: "something",
	}

	blogUseCaseMock.
		On("CreatePost", mock.Anything, &post).
		Return(int64(1), errors.New("DB error"))

	expect.POST("/v1/api/blog/posts").
		WithJSON(post).
		Expect().
		Status(http.StatusInternalServerError).
		Body().IsEqual("{\"error\":\"DB error\"}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_CreatePost_RequestDoesNotHaveTitle_ShouldBadRequest(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	invalidPost := noTitlePost{
		Author:  "Anton",
		Content: "something",
	}

	expect.POST("/v1/api/blog/posts").
		WithJSON(invalidPost).
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_CreatePost_RequestDoesNotHaveAuthor_ShouldBadRequest(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	invalidPost := noAuthorPost{
		Title:   "Title 1",
		Content: "something",
	}

	expect.POST("/v1/api/blog/posts").
		WithJSON(invalidPost).
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_CreatePost_RequestDoesNotHaveContent_ShouldBadRequest(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	invalidPost := noContentPost{
		Author: "Anton",
		Title:  "Title 1",
	}

	expect.POST("/v1/api/blog/posts").
		WithJSON(invalidPost).
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

type noTitlePost struct {
	Id      int64
	Author  string
	Content string
}

type noAuthorPost struct {
	Id      int64
	Title   string
	Content string
}

type noContentPost struct {
	Id     int64
	Author string
	Title  string
}

func Test_UpdatePost_ShouldReturnOk(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	post := domain.Post{
		ID:      int64(1),
		Author:  "Anton",
		Title:   "New title",
		Content: "something",
	}

	blogUseCaseMock.
		On("UpdatePost", mock.Anything, &post, int64(1)).
		Return(nil)

	expect.PUT("/v1/api/blog/posts/1").
		WithJSON(post).
		Expect().
		Status(http.StatusOK).
		Body().IsEqual("{\"Id\":1}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_UpdatePost_Error_ShouldReturn500Status(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	post := domain.Post{
		ID:      int64(1),
		Author:  "Anton",
		Title:   "New title",
		Content: "something",
	}

	blogUseCaseMock.
		On("UpdatePost", mock.Anything, &post, int64(1)).
		Return(errors.New("DB error"))

	expect.PUT("/v1/api/blog/posts/1").
		WithJSON(post).
		Expect().
		Status(http.StatusInternalServerError).
		Body().IsEqual("{\"error\":\"DB error\"}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_UpdatePost_NoAuthor_ShouldReturnBadRequest(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	invalidPost := noAuthorPost{
		Id:      int64(1),
		Title:   "New title",
		Content: "something",
	}

	expect.PUT("/v1/api/blog/posts/1").
		WithJSON(invalidPost).
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_UpdatePost_NoTitle_ShouldReturnBadRequest(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	invalidPost := noTitlePost{
		Id:      int64(1),
		Author:  "Anton",
		Content: "something",
	}

	expect.PUT("/v1/api/blog/posts/1").
		WithJSON(invalidPost).
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_UpdatePost_NoContent_ShouldReturnBadRequest(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	invalidPost := noContentPost{
		Id:     int64(1),
		Author: "Anton",
		Title:  "New title",
	}

	expect.PUT("/v1/api/blog/posts/1").
		WithJSON(invalidPost).
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_UpdatePost_IncorrectId_ShouldReturnBadRequest(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	post := domain.Post{
		ID:      int64(1),
		Author:  "Anton",
		Title:   "New title",
		Content: "something",
	}

	expect.PUT("/v1/api/blog/posts/a").
		WithJSON(post).
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_DeletePost_ShouldReturnNoContentStatus(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	blogUseCaseMock.
		On("DeletePost", mock.Anything, int64(1)).
		Return(nil)

	expect.DELETE("/v1/api/blog/posts/1").
		Expect().
		Status(http.StatusNoContent)

	blogUseCaseMock.AssertExpectations(t)
}

func Test_DeletePost_Error_ShouldReturn500Status(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	blogUseCaseMock.
		On("DeletePost", mock.Anything, int64(1)).
		Return(errors.New("DB error"))

	expect.DELETE("/v1/api/blog/posts/1").
		Expect().
		Status(http.StatusInternalServerError).
		Body().IsEqual("{\"error\":\"DB error\"}")

	blogUseCaseMock.AssertExpectations(t)
}

func Test_DeletePost_IncorrectId_ShouldReturnBadRequestStatus(t *testing.T) {
	var blogUseCaseMock = new(mocks.IBlogUseCase)
	expect := SetupServer(t, blogUseCaseMock)

	expect.DELETE("/v1/api/blog/posts/a").
		Expect().
		Status(http.StatusBadRequest)

	blogUseCaseMock.AssertExpectations(t)
}
