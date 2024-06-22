package repository_test

import (
	"context"
	"testing"

	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/repository"
	"github.com/kondrushin/blog/mocks"
	"github.com/stretchr/testify/assert"
)

var mockRepository = new(mocks.IBlogRepository)
var ctx = context.Background()

func Test_CreatePostAndGetPost_ShouldAddPostToRepo(t *testing.T) {
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	postId, err := repo.CreatePost(ctx, post)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	foundPost, err := repo.GetPost(ctx, postId)
	assert.NoError(t, err)

	assert.EqualValues(t, post, foundPost)
}

func Test_CreatePost_ShouldSetIdSequentially(t *testing.T) {
	repo := repository.NewRepository()

	var post1 = &domain.Post{
		Author:  "Anton1",
		Title:   "On mockery1",
		Content: "qwerty1",
	}
	postId, err := repo.CreatePost(ctx, post1)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	var post2 = &domain.Post{
		Author:  "Anton2",
		Title:   "On mockery2",
		Content: "qwerty2",
	}
	postId, err = repo.CreatePost(ctx, post2)
	assert.EqualValues(t, 2, postId)
	assert.NoError(t, err)

	var post3 = &domain.Post{
		Author:  "Anton3",
		Title:   "On mockery3",
		Content: "qwerty3",
	}
	postId, err = repo.CreatePost(ctx, post3)
	assert.EqualValues(t, 3, postId)
	assert.NoError(t, err)
}

func Test_CreatePost_ShouldNotReuseIdAfterDelete(t *testing.T) {
	repo := repository.NewRepository()

	var post1 = &domain.Post{
		Author:  "Anton1",
		Title:   "On mockery1",
		Content: "qwerty1",
	}
	postId, err := repo.CreatePost(ctx, post1)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	err = repo.DeletePost(ctx, postId)
	assert.NoError(t, err)

	var post2 = &domain.Post{
		Author:  "Anton2",
		Title:   "On mockery2",
		Content: "qwerty2",
	}
	postId, err = repo.CreatePost(ctx, post2)
	assert.EqualValues(t, 2, postId)
	assert.NoError(t, err)
}

func Test_DeletePost_ShouldDeletePostFromRepo(t *testing.T) {
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	postId, err := repo.CreatePost(ctx, post)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	foundPost, err := repo.GetPost(ctx, postId)
	assert.EqualValues(t, post, foundPost)

	err = repo.DeletePost(ctx, postId)
	assert.NoError(t, err)

	foundPost, err = repo.GetPost(ctx, postId)
	assert.Nil(t, foundPost)
	assert.ErrorIs(t, domain.ErrorPostNotFound, err)
}

func Test_DeletePost_NoItemToDelete_ShouldNotReturnError(t *testing.T) {
	repo := repository.NewRepository()

	err := repo.DeletePost(ctx, int64(63))
	assert.NoError(t, err)
}

func Test_UpdatePost_ShouldUpdateRepo(t *testing.T) {
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	postId, err := repo.CreatePost(ctx, post)
	assert.EqualValues(t, 1, postId)
	assert.NoError(t, err)

	foundPost, err := repo.GetPost(ctx, postId)
	assert.EqualValues(t, post, foundPost)

	post.Author = "Anton2"
	post.Title = "New title"
	post.Content = "www"

	err = repo.UpdatePost(ctx, post, postId)
	assert.NoError(t, err)

	foundPost, err = repo.GetPost(ctx, postId)
	assert.EqualValues(t, post, foundPost)
}

func Test_UpdatePost_NoItemToUpdate_ShouldReturnError(t *testing.T) {
	repo := repository.NewRepository()

	var post = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	err := repo.UpdatePost(ctx, post, int64(34))
	assert.ErrorIs(t, domain.ErrorPostNotFound, err)
}
