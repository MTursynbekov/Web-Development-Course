package service

import (
	"twitter/internal/model"
	"twitter/internal/store"
)

type PostsService interface {
	CreatePosts(post *model.Posts) error
	GetUsersPosts(userId int) ([]*model.Posts, error)
	GetPost(userId, postId int) (*model.Posts, error)
	UpdatePost(post *model.Posts) error
	DeletePost(postId int) error
}

type postService struct {
	store store.Store
}

func NewPostsService(store store.Store) PostsService {
	return &postService{
		store: store,
	}
}

func (s *postService) CreatePosts(post *model.Posts) error {
	err := s.store.CreatePost(post)
	return err
}

func (s *postService) GetUsersPosts(userId int) ([]*model.Posts, error) {
	Postss, err := s.store.GetUsersPosts(userId)
	return Postss, err
}

func (s *postService) GetPost(userId, postId int) (*model.Posts, error) {
	post, err := s.store.GetUsersPostById(userId, postId)
	return post, err
}

func (s *postService) UpdatePost(post *model.Posts) error {
	err := s.store.UpdatePost(post)
	return err
}

func (s *postService) DeletePost(postId int) error {
	err := s.store.DeletePost(postId)
	return err
}
