package services

import (
	"github.com/karankumarshreds/go-blog-api/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogService struct {
	blogRepo BlogRepoPort
}

type BlogRepoPort interface {
	Create(payload core.CreateBlogDto) (*primitive.ObjectID, error)
}

func NewBlogService(blogRepo BlogRepoPort) *BlogService {
	return &BlogService{blogRepo}
}

func (b *BlogService) Create(payload core.CreateBlogDto) (*primitive.ObjectID, error) {
	return b.blogRepo.Create(payload)
}
