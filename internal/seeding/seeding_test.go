package seeding_test

import (
	"context"
	"encoding/json"
	"errors"

	"os"
	"testing"

	"github.com/kondrushin/blog/internal/domain"
	"github.com/kondrushin/blog/internal/seeding"
	"github.com/kondrushin/blog/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Seed_Success(t *testing.T) {
	ctx := context.Background()

	repositoryMock := new(mocks.IBlogRepository)

	blog := seeding.BlogFileModel{
		Posts: []seeding.PostFileModel{
			{Id: 1, Author: "Anton", Title: "Big title", Content: "Big Content"},
			{Id: 2, Author: "Jonny", Title: "Small title", Content: "Small Content"},
		}}

	tempFile := writeDataToTestFile(blog)
	defer os.Remove(tempFile.Name())

	post1 := &domain.Post{
		Author:  "Anton",
		Title:   "Big title",
		Content: "Big Content",
	}
	post2 := &domain.Post{
		Author:  "Jonny",
		Title:   "Small title",
		Content: "Small Content",
	}

	repositoryMock.
		On("CreatePost", mock.Anything, post1).
		Once().
		Return(int64(1), nil)

	repositoryMock.
		On("CreatePost", mock.Anything, post2).
		Once().
		Return(int64(2), nil)

	err := seeding.Seed(ctx, tempFile.Name(), repositoryMock)
	assert.NoError(t, err)
	repositoryMock.AssertExpectations(t)
}

func Test_Seed_ErrorWhileCreatingPostInRepo(t *testing.T) {
	ctx := context.Background()

	repositoryMock := new(mocks.IBlogRepository)

	blog := seeding.BlogFileModel{
		Posts: []seeding.PostFileModel{
			{Id: 1, Author: "Anton", Title: "Big title", Content: "Big Content"},
		}}

	tempFile := writeDataToTestFile(blog)
	defer os.Remove(tempFile.Name())

	post1 := &domain.Post{
		Author:  "Anton",
		Title:   "Big title",
		Content: "Big Content",
	}

	repositoryMock.
		On("CreatePost", mock.Anything, post1).
		Once().
		Return(int64(0), errors.New("ERROR"))

	err := seeding.Seed(ctx, tempFile.Name(), repositoryMock)
	assert.Error(t, err)
	assert.Equal(t, "Could not seed data from a file. Error: ERROR", err.Error())
	repositoryMock.AssertExpectations(t)
}

func Test_Seed_FileDoesNotExist(t *testing.T) {
	ctx := context.Background()

	repositoryMock := new(mocks.IBlogRepository)

	err := seeding.Seed(ctx, "IdoNotExist.json", repositoryMock)
	assert.Error(t, err)
	assert.Equal(t, "open IdoNotExist.json: no such file or directory", err.Error())
	repositoryMock.AssertExpectations(t)
}

func writeDataToTestFile(dt any) *os.File {
	data, _ := json.Marshal(dt)
	tempFile, _ := os.CreateTemp("", "test.json")
	_, _ = tempFile.Write(data)
	return tempFile
}
