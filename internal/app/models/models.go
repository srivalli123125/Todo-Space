package models

import "github.com/gocql/gocql"

type Todo struct {
	ID          gocql.UUID `json:"id"`
	UserID      gocql.UUID `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created     string     `json:"created"`
	Updated     string     `json:"updated"`
}