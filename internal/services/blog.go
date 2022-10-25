package services

import (
	"github.com/karankumarshreds/go-blog-api/internal/core"
)

type BlogService struct {
	blogRepo BlogRepoPort
}

type BlogRepoPort interface {
	Create(payload core.CreateBlogDto) (*int, error)
}

func NewBlogService(blogRepo BlogRepoPort) *BlogService {
	return &BlogService{blogRepo}
}

func (b *BlogService) Create(payload core.CreateBlogDto) (*int, error) {
	return b.blogRepo.Create(payload)
}
