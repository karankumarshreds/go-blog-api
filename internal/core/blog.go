package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Body        string             `json:"body"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateBlogDto struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	Body        string `validate:"required" json:"body"`
}
