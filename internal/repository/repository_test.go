package repository_test

import (
	"context"
	"testing"

	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/repository"
	"github.com/stretchr/testify/assert"
)

type UseCaseTestSuite struct {
	ctx context.Context
}

func SetSuite() *UseCaseTestSuite {
	var suite = UseCaseTestSuite{}
	suite.ctx = context.Background()
	return &suite
}

func Test_CreatePostAndGetPost_ShouldAddPostToRepo(t *testing.T) {
	suite := SetSuite()
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	postId, err := repo.CreatePost(suite.ctx, post)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	foundPost, err := repo.GetPost(suite.ctx, postId)
	assert.NoError(t, err)

	assert.EqualValues(t, post, foundPost)
}

func Test_CreatePost_ShouldSetIdSequentially(t *testing.T) {
	suite := SetSuite()
	repo := repository.NewRepository()

	var post1 = &domain.Post{
		Author:  "Anton1",
		Title:   "On mockery1",
		Content: "qwerty1",
	}
	postId, err := repo.CreatePost(suite.ctx, post1)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	var post2 = &domain.Post{
		Author:  "Anton2",
		Title:   "On mockery2",
		Content: "qwerty2",
	}
	postId, err = repo.CreatePost(suite.ctx, post2)
	assert.EqualValues(t, 2, postId)
	assert.NoError(t, err)

	var post3 = &domain.Post{
		Author:  "Anton3",
		Title:   "On mockery3",
		Content: "qwerty3",
	}
	postId, err = repo.CreatePost(suite.ctx, post3)
	assert.EqualValues(t, 3, postId)
	assert.NoError(t, err)
}

func Test_CreatePost_ShouldNotReuseIdAfterDelete(t *testing.T) {
	suite := SetSuite()
	repo := repository.NewRepository()

	var post1 = &domain.Post{
		Author:  "Anton1",
		Title:   "On mockery1",
		Content: "qwerty1",
	}
	postId, err := repo.CreatePost(suite.ctx, post1)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	err = repo.DeletePost(suite.ctx, postId)
	assert.NoError(t, err)

	var post2 = &domain.Post{
		Author:  "Anton2",
		Title:   "On mockery2",
		Content: "qwerty2",
	}
	postId, err = repo.CreatePost(suite.ctx, post2)
	assert.EqualValues(t, 2, postId)
	assert.NoError(t, err)
}

func Test_DeletePost_ShouldDeletePostFromRepo(t *testing.T) {
	suite := SetSuite()
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	postId, err := repo.CreatePost(suite.ctx, post)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	foundPost, err := repo.GetPost(suite.ctx, postId)
	assert.EqualValues(t, post, foundPost)

	err = repo.DeletePost(suite.ctx, postId)
	assert.NoError(t, err)

	foundPost, err = repo.GetPost(suite.ctx, postId)
	assert.Nil(t, foundPost)
	assert.ErrorIs(t, domain.ErrorPostNotFound, err)
}

func Test_DeletePost_NoItemToDelete_ShouldNotReturnError(t *testing.T) {
	suite := SetSuite()
	repo := repository.NewRepository()

	err := repo.DeletePost(suite.ctx, int64(63))
	assert.NoError(t, err)
}

func Test_UpdatePost_ShouldUpdateRepo(t *testing.T) {
	suite := SetSuite()
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	postId, err := repo.CreatePost(suite.ctx, post)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	foundPost, err := repo.GetPost(suite.ctx, postId)
	assert.EqualValues(t, post, foundPost)

	post.Author = "Anton2"
	post.Title = "New title"
	post.Content = "www"

	err = repo.UpdatePost(suite.ctx, post, postId)
	assert.NoError(t, err)

	foundPost, err = repo.GetPost(suite.ctx, postId)
	assert.EqualValues(t, post, foundPost)
}

func Test_UpdatePost_NoItemToUpdate_ShouldReturnError(t *testing.T) {
	suite := SetSuite()
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	err := repo.UpdatePost(suite.ctx, post, int64(34))
	assert.ErrorIs(t, domain.ErrorPostNotFound, err)
}
