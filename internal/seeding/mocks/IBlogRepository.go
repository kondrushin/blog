// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/kondrushin/blog/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// IBlogRepository is an autogenerated mock type for the IBlogRepository type
type IBlogRepository struct {
	mock.Mock
}

// CreatePost provides a mock function with given fields: ctx, post
func (_m *IBlogRepository) CreatePost(ctx context.Context, post *domain.Post) (int64, error) {
	ret := _m.Called(ctx, post)

	if len(ret) == 0 {
		panic("no return value specified for CreatePost")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Post) (int64, error)); ok {
		return rf(ctx, post)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Post) int64); ok {
		r0 = rf(ctx, post)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Post) error); ok {
		r1 = rf(ctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIBlogRepository creates a new instance of IBlogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIBlogRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IBlogRepository {
	mock := &IBlogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
