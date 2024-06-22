package usecase_test

import (
	"context"
	"errors"

	"testing"

	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/usecase"
	"github.com/kondrushin/blog/mocks"
	"github.com/stretchr/testify/assert"
)

var mockRepository = new(mocks.IBlogRepository)
var blogUseCase = usecase.NewBlogUseCase(mockRepository)
var ctx = context.Background()
var postInRepo = &domain.Post{
	Author:  "Anton",
	Title:   "On mockery",
	Content: "qwerty"}

func Test_GetPost_ShouldReturnPostFromRepositry(t *testing.T) {
	id := int64(45)

	mockRepository.
		On("GetPost", ctx, id).
		Once().
		Return(postInRepo, nil)

	post, err := blogUseCase.GetPost(ctx, id)

	assert.EqualValues(t, postInRepo, post)
	assert.NoError(t, err)
	mockRepository.AssertExpectations(t)
}

func Test_GetPost_ShouldReturnErrorFromRepositry(t *testing.T) {
	id := int64(45)

	error := errors.New("problem")

	mockRepository.
		On("GetPost", ctx, id).
		Once().
		Return(nil, error)

	_, err := blogUseCase.GetPost(ctx, id)

	assert.ErrorIs(t, error, err)
	mockRepository.AssertExpectations(t)
}

func Test_GetPosts_ShouldReturnPostsFromRepositry(t *testing.T) {
	post1 := &domain.Post{
		Author:  "Anton1",
		Title:   "On mockery",
		Content: "qwerty"}
	post2 := &domain.Post{
		Author:  "Anton2",
		Title:   "On mockery",
		Content: "qwerty"}

	postsInRepo := []*domain.Post{post1, post2}

	mockRepository.
		On("GetPosts", ctx).
		Once().
		Return(postsInRepo)

	posts := blogUseCase.GetPosts(ctx)

	assert.Equal(t, len(postsInRepo), len(posts))
	for i := 0; i < len(posts); i++ {
		assert.EqualValues(t, postsInRepo[i], posts[i])
	}

	mockRepository.AssertExpectations(t)
}

func Test_CreatePost_ShouldReturnIdFromRepositry(t *testing.T) {
	postIdFromRepo := int64(45)

	mockRepository.
		On("CreatePost", ctx, postInRepo).
		Once().
		Return(postIdFromRepo, nil)

	id, err := blogUseCase.CreatePost(ctx, postInRepo)

	assert.Equal(t, postIdFromRepo, id)
	assert.NoError(t, err)
	mockRepository.AssertExpectations(t)
}

func Test_CreatePost_Error_ShouldReturnErrorFromRepositry(t *testing.T) {
	error := errors.New("problem")

	mockRepository.
		On("CreatePost", ctx, postInRepo).
		Once().
		Return(int64(0), error)

	_, err := blogUseCase.CreatePost(ctx, postInRepo)

	assert.ErrorIs(t, error, err)
	mockRepository.AssertExpectations(t)
}

func Test_UpdatePost_ShouldCallRepoMethodOnce(t *testing.T) {
	id := int64(45)

	mockRepository.
		On("UpdatePost", ctx, postInRepo, id).
		Once().
		Return(nil)

	err := blogUseCase.UpdatePost(ctx, postInRepo, id)

	assert.NoError(t, err)
	mockRepository.AssertExpectations(t)
}

func Test_UpdatePost_Error_ShouldReturnErrorFromRepositry(t *testing.T) {
	id := int64(45)

	error := errors.New("problem")
	mockRepository.
		On("UpdatePost", ctx, postInRepo, id).
		Once().
		Return(error)

	err := blogUseCase.UpdatePost(ctx, postInRepo, id)

	assert.ErrorIs(t, error, err)
	mockRepository.AssertExpectations(t)
}

func Test_DeleteePost_ShouldCallRepoMethodOnce(t *testing.T) {
	id := int64(45)

	mockRepository.
		On("DeletePost", ctx, id).
		Once().
		Return(nil)

	err := blogUseCase.DeletePost(ctx, id)

	assert.NoError(t, err)
	mockRepository.AssertExpectations(t)
}

func Test_DeleteePost_ShouldPassErrorFromRepo(t *testing.T) {
	id := int64(45)

	error := errors.New("problem")

	mockRepository.
		On("DeletePost", ctx, id).
		Once().
		Return(error)

	err := blogUseCase.DeletePost(ctx, id)

	assert.ErrorIs(t, error, err)
	mockRepository.AssertExpectations(t)
}
