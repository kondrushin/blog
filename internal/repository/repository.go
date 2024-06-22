package repository

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/kondrushin/blog/internal/domain"
)

type Repository struct {
	mutex sync.RWMutex
	posts map[int64]*domain.Post

	sequenceId *int64
}

func NewRepository() *Repository {
	var startId int64 = 0 // it will start with 1
	return &Repository{
		sequenceId: &startId,
		posts:      map[int64]*domain.Post{},
	}
}

func (r *Repository) GetPost(ctx context.Context, id int64) (*domain.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	p, isIn := r.posts[id]
	if isIn {
		return p, nil
	}

	return nil, domain.ErrorPostNotFound
}

func (r *Repository) GetPosts(ctx context.Context) []*domain.Post {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	posts := make([]*domain.Post, 0, len(r.posts))
	for _, p := range r.posts {
		posts = append(posts, p)
	}

	return posts
}

func (r *Repository) CreatePost(ctx context.Context, post *domain.Post) (int64, error) {
	nextPostId := r.getNextSequenceId()
	post.ID = nextPostId

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.posts[nextPostId] = post
	return nextPostId, nil
}

func (r *Repository) UpdatePost(ctx context.Context, post *domain.Post, id int64) error {
	post.ID = id

	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, isIn := r.posts[id]
	if !isIn {
		return domain.ErrorPostNotFound
	}

	r.posts[id] = post
	return nil
}

func (r *Repository) DeletePost(ctx context.Context, id int64) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.posts, id)
	return nil
}

func (s *Repository) getNextSequenceId() int64 {
	return atomic.AddInt64(s.sequenceId, 1)
}
