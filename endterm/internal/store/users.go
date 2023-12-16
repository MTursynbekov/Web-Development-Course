package store

import (
	"database/sql"
	"twitter/internal/model"
)

func (s *store) CreateUser(user *model.User) (int, error) {
	query := `
	insert into users (username, password)
	values ($1, $2)
	returning id`

	var id int
	err := s.db.QueryRow(query, user.Username, user.Password).Scan(&id)
	return id, err
}

func (s *store) GetUser(username string) (*model.User, error) {
	var user model.User
	query := `
	select id, username, password from users
	where username = $1`
	err := s.db.Get(&user, query, username)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}
