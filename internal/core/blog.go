package core

import "time"

type Blog struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateBlogDto struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	Body        string `validate:"required" json:"body"`
}
