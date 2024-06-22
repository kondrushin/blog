package usecase_test

import (
	"context"
	"errors"

	"testing"

	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/usecase"
	"github.com/kondrushin/blog/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

type UseCaseTestSuite struct {
	mockRepository *mocks.IBlogRepository
	blogUseCase    *usecase.BlogUseCase
	ctx            context.Context
	postInRepo     *domain.Post
}

func SetSuite() *UseCaseTestSuite {
	var suite = UseCaseTestSuite{}
	suite.mockRepository = new(mocks.IBlogRepository)
	suite.blogUseCase = usecase.NewBlogUseCase(suite.mockRepository)
	suite.ctx = context.Background()
	suite.postInRepo = &domain.Post{
		Author:  "Anton",
		Title:   "On mockery",
		Content: "qwerty",
	}

	return &suite
}

func Test_GetPost_ShouldReturnPostFromRepositry(t *testing.T) {
	suite := SetSuite()
	id := int64(45)

	suite.mockRepository.
		On("GetPost", suite.ctx, id).
		Once().
		Return(suite.postInRepo, nil)

	post, err := suite.blogUseCase.GetPost(suite.ctx, id)

	assert.EqualValues(t, suite.postInRepo, post)
	assert.NoError(t, err)
	suite.mockRepository.AssertExpectations(t)
}

func Test_GetPost_ShouldReturnErrorFromRepositry(t *testing.T) {
	suite := SetSuite()
	id := int64(45)

	error := errors.New("problem")

	suite.mockRepository.
		On("GetPost", suite.ctx, id).
		Once().
		Return(nil, error)

	_, err := suite.blogUseCase.GetPost(suite.ctx, id)

	assert.ErrorIs(t, error, err)
	suite.mockRepository.AssertExpectations(t)
}

func Test_GetPosts_ShouldReturnPostsFromRepositry(t *testing.T) {
	suite := SetSuite()
	post1 := &domain.Post{
		Author:  "Anton1",
		Title:   "On mockery",
		Content: "qwerty",
	}
	post2 := &domain.Post{
		Author:  "Anton2",
		Title:   "On mockery",
		Content: "qwerty",
	}

	postsInRepo := []*domain.Post{post1, post2}

	suite.mockRepository.
		On("GetPosts", suite.ctx).
		Once().
		Return(postsInRepo)

	posts := suite.blogUseCase.GetPosts(suite.ctx)

	assert.Equal(t, len(postsInRepo), len(posts))
	for i := 0; i < len(posts); i++ {
		assert.EqualValues(t, postsInRepo[i], posts[i])
	}

	suite.mockRepository.AssertExpectations(t)
}

func Test_CreatePost_ShouldReturnIdFromRepositry(t *testing.T) {
	suite := SetSuite()
	postIdFromRepo := int64(45)

	suite.mockRepository.
		On("CreatePost", suite.ctx, suite.postInRepo).
		Once().
		Return(postIdFromRepo, nil)

	id, err := suite.blogUseCase.CreatePost(suite.ctx, suite.postInRepo)

	assert.Equal(t, postIdFromRepo, id)
	assert.NoError(t, err)
	suite.mockRepository.AssertExpectations(t)
}

func Test_CreatePost_Error_ShouldReturnErrorFromRepositry(t *testing.T) {
	suite := SetSuite()
	error := errors.New("problem")

	suite.mockRepository.
		On("CreatePost", suite.ctx, suite.postInRepo).
		Once().
		Return(int64(0), error)

	_, err := suite.blogUseCase.CreatePost(suite.ctx, suite.postInRepo)

	assert.ErrorIs(t, error, err)
	suite.mockRepository.AssertExpectations(t)
}

func Test_UpdatePost_ShouldCallRepoMethodOnce(t *testing.T) {
	suite := SetSuite()
	id := int64(45)

	suite.mockRepository.
		On("UpdatePost", suite.ctx, suite.postInRepo, id).
		Once().
		Return(nil)

	err := suite.blogUseCase.UpdatePost(suite.ctx, suite.postInRepo, id)

	assert.NoError(t, err)
	suite.mockRepository.AssertExpectations(t)
}

func Test_UpdatePost_Error_ShouldReturnErrorFromRepositry(t *testing.T) {
	suite := SetSuite()
	id := int64(45)

	error := errors.New("problem")
	suite.mockRepository.
		On("UpdatePost", suite.ctx, suite.postInRepo, id).
		Once().
		Return(error)

	err := suite.blogUseCase.UpdatePost(suite.ctx, suite.postInRepo, id)

	assert.ErrorIs(t, error, err)
	suite.mockRepository.AssertExpectations(t)
}

func Test_DeleteePost_ShouldCallRepoMethodOnce(t *testing.T) {
	suite := SetSuite()
	id := int64(45)

	suite.mockRepository.
		On("DeletePost", suite.ctx, id).
		Once().
		Return(nil)

	err := suite.blogUseCase.DeletePost(suite.ctx, id)

	assert.NoError(t, err)
	suite.mockRepository.AssertExpectations(t)
}

func Test_DeleteePost_ShouldPassErrorFromRepo(t *testing.T) {
	suite := SetSuite()
	id := int64(45)

	error := errors.New("problem")

	suite.mockRepository.
		On("DeletePost", suite.ctx, id).
		Once().
		Return(error)

	err := suite.blogUseCase.DeletePost(suite.ctx, id)

	assert.ErrorIs(t, error, err)
	suite.mockRepository.AssertExpectations(t)
}
