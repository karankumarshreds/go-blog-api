package repos

import (
	"context"
	"errors"
	"log"
	"time"

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
	c := b.DB.Collection(constants.MongoCollections.BLOGS)
	blog := CreateBsonObject(map[string]interface{}{
		"title":       payload.Title,
		"description": payload.Description,
		"body":        payload.Body,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	})
	res, err := c.InsertOne(context.TODO(), blog)
	if err != nil {
		log.Println("error")
		return nil, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func (b *BlogRepo) Get(id string) (*core.Blog, error) {
	c := b.DB.Collection(constants.MongoCollections.BLOGS)
	blog := new(core.Blog)
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := CreateBsonObject(map[string]interface{}{
		"_id": _id,
	})
	if err := c.FindOne(context.TODO(), filter).Decode(blog); err != nil {
		msg := "Object with given id not found"
		log.Println(msg, err)
		return nil, errors.New(msg)
	}
	return blog, nil
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
