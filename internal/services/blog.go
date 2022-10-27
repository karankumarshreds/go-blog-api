package services

import (
	"github.com/karankumarshreds/go-blog-api/custom_errors"
	"github.com/karankumarshreds/go-blog-api/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogService struct {
	blogRepo BlogRepoPort
}

type BlogRepoPort interface {
	Create(payload core.CreateBlogDto) (*primitive.ObjectID, *custom_errors.CustomError)
	Get(id string) (*core.Blog, *custom_errors.CustomError)
	Update(id string, payload core.CreateBlogDto) (*core.Blog, *custom_errors.CustomError)
	Delete(id string) *custom_errors.CustomError
}

func NewBlogService(blogRepo BlogRepoPort) *BlogService {
	return &BlogService{blogRepo}
}

func (b *BlogService) Create(payload core.CreateBlogDto) (*primitive.ObjectID, *custom_errors.CustomError) {
	return b.blogRepo.Create(payload)
}

func (b *BlogService) Get(id string) (*core.Blog, *custom_errors.CustomError) {
	return b.blogRepo.Get(id)
}

func (b *BlogService) Update(id string, payload core.CreateBlogDto) (*core.Blog, *custom_errors.CustomError) {
	return b.blogRepo.Update(id, payload)
}

func (b *BlogService) Delete(id string) *custom_errors.CustomError {
	return b.blogRepo.Delete(id)
}
