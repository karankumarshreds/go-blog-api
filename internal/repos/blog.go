package repos

import (
	"context"
	"log"

	"github.com/karankumarshreds/go-blog-api/constants"
	"github.com/karankumarshreds/go-blog-api/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepo struct {
	DB *mongo.Database
}

// Constructor function for the blog repo
func NewBlogRepo(db *mongo.Database) *BlogRepo {
	return &BlogRepo{db}
}

func (b *BlogRepo) Create(payload core.CreateBlogDto) (*primitive.ObjectID, error) {
	s := b.DB.Collection(constants.MongoCollections.BLOGS)
	blog := CreateBsonObject(map[string]interface{}{
		"title":       payload.Title,
		"description": payload.Description,
		"body":        payload.Body,
	})
	res, err := s.InsertOne(context.TODO(), blog)
	if err != nil {
		log.Println("error")
		return nil, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func CreateBsonObject(data map[string]interface{}) bson.D {
	var bsonObject bson.D
	for key, val := range data {
		bsonObject = append(
			bsonObject,
			primitive.E{Key: key, Value: val},
		)
	}
	return bsonObject
}
