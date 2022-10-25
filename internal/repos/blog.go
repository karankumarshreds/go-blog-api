package repos

import (
	"context"
	"log"

	"github.com/karankumarshreds/go-blog-api/constants"
	"github.com/karankumarshreds/go-blog-api/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepo struct {
	DB *mongo.Database
}

// Constructor function for the blog repo
func NewBlogRepo(db *mongo.Database) *BlogRepo {
	return &BlogRepo{db}
}

func (b *BlogRepo) Create(payload core.CreateBlogDto) (*int, error) {
	log.Println("A")
	s := b.DB.Collection(constants.MongoCollections.BLOGS)
	log.Println("A2")
	blog := bson.D{
		{"title", payload.Title},
		{"description", payload.Description},
		{"body", payload.Body},
	}
	res, err := s.InsertOne(context.TODO(), blog)
	if err != nil {
		log.Println("error")
		return nil, err
	} else {
		log.Println(res.InsertedID)
	}
	i := 69
	return &i, nil
}
