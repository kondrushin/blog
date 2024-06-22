package seeding

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kondrushin/blog/internal/domain"
)

type Repository interface {
	CreatePost(ctx context.Context, post *domain.Post) (int64, error)
}

func Seed(ctx context.Context, filePath string, repository Repository) error {
	posts, err := getPostsFromFile(filePath)
	if err != nil {
		return err
	}

	err = addPostsToRepository(ctx, posts, repository)
	if err != nil {
		return err
	}

	return nil
}

func addPostsToRepository(ctx context.Context, posts []PostFileModel, repository Repository) error {
	for _, p := range posts {
		_, err := repository.CreatePost(ctx, &domain.Post{
			ID:      int64(p.ID),
			Author:  p.Author,
			Title:   p.Title,
			Content: p.Content,
		})

		if err != nil {
			return fmt.Errorf("Could not seed data from a file. Error: %w", err)
		}
	}

	return nil
}

func getPostsFromFile(filePath string) ([]PostFileModel, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	blogFromFile := BlogFileModel{}
	if err := json.Unmarshal(data, &blogFromFile); err != nil {
		return nil, err
	}

	return blogFromFile.Posts, nil
}

type PostFileModel struct {
	ID      int64  `json:"-"`
	Author  string `json:"author" `
	Title   string `json:"title"`
	Content string `json:"content"`
}

type BlogFileModel struct {
	Posts []PostFileModel `json:"Posts"`
}
