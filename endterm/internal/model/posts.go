package model

import "time"

type Posts struct {
	Id        int
	Title     string
	Text      string
	CreatedAt time.Time `db:"created_at"`
	UserId    int       `db:"user_id"`
}
