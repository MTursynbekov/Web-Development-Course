package store

import (
	"database/sql"
	"twitter/internal/model"
)

func (s *store) CreatePost(post *model.Posts) error {
	query := `
	INSERT INTO posts (title, text, user_id)
	VALUES ($1, $2, $3)`

	_, err := s.db.Exec(query, post.Title, post.Text, post.UserId)

	return err
}

func (s *store) GetUsersPosts(userId int) ([]*model.Posts, error) {
	posts := make([]*model.Posts, 0)
	query := `
	SELECT id, title, text, user_id, created_at FROM posts
	WHERE user_id = $1`

	err := s.db.Select(&posts, query, userId)

	return posts, err
}

func (s *store) GetUsersPostById(userId, postId int) (*model.Posts, error) {
	var post model.Posts
	query := `
	SELECT id, title, text, user_id, created_at FROM posts
	WHERE user_id = $1 and id = $2`

	err := s.db.Get(&post, query, userId, postId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &post, err
}

func (s *store) UpdatePost(post *model.Posts) error {
	query := `
	UPDATE posts
	SET title = $1,
		text = $2,
		user_id = $3
		created_at = $4
	WHERE id = $5`

	_, err := s.db.Exec(query, post.Title, post.Text, post.UserId, post.CreatedAt)

	return err

}

func (s *store) DeletePost(postId int) error {
	query := `
	DELETE FROM posts
	WHERE id = $1`

	_, err := s.db.Exec(query, postId)

	return err
}
