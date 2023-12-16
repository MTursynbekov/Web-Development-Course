package store

import (
	"twitter/internal/model"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Migrate()
	CreateUser(user *model.User) (int, error)
	GetUser(username string) (*model.User, error)
	CreatePost(post *model.Posts) error
	GetUsersPosts(userId int) ([]*model.Posts, error)
	GetUsersPostById(userId, postId int) (*model.Posts, error)
	UpdatePost(post *model.Posts) error
	DeletePost(postId int) error
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{
		db: db,
	}
}

func (s *store) Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR UNIQUE,
		password VARCHAR,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP
	);
	   
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		text TEXT NOT NULL,
		user_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP
	   );
	`

	s.db.Exec(query)
}
