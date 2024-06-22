package usecase

import (
	"context"

	"github.com/kondrushin/blog/internal/domain"
)

type IBlogRepository interface {
	GetPost(ctx context.Context, id int64) (*domain.Post, error)
	GetPosts(ctx context.Context) []*domain.Post
	CreatePost(ctx context.Context, post *domain.Post) (int64, error)
	UpdatePost(ctx context.Context, post *domain.Post, id int64) error
	DeletePost(ctx context.Context, id int64) error
}

type BlogUseCase struct {
	repository IBlogRepository
}

func NewBlogUseCase(repository IBlogRepository) *BlogUseCase {
	return &BlogUseCase{repository: repository}
}

func (b *BlogUseCase) GetPost(ctx context.Context, id int64) (*domain.Post, error) {
	return b.repository.GetPost(ctx, id)
}

func (b *BlogUseCase) GetPosts(ctx context.Context) []*domain.Post {
	return b.repository.GetPosts(ctx)
}

func (b *BlogUseCase) CreatePost(ctx context.Context, post *domain.Post) (int64, error) {
	return b.repository.CreatePost(ctx, post)
}

func (b *BlogUseCase) UpdatePost(ctx context.Context, post *domain.Post, id int64) error {
	return b.repository.UpdatePost(ctx, post, id)
}

func (b *BlogUseCase) DeletePost(ctx context.Context, id int64) error {
	return b.repository.DeletePost(ctx, id)
}
