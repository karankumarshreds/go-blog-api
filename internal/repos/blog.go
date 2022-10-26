package repos

import (
	"context"
	"log"
	"time"

	"github.com/karankumarshreds/go-blog-api/constants"
	"github.com/karankumarshreds/go-blog-api/custom_errors"
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

func (b *BlogRepo) Create(payload core.CreateBlogDto) (*primitive.ObjectID, *custom_errors.CustomError) {
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
		msg := "Cannot insert document"
		log.Println(msg, err)
		return nil, custom_errors.InternalServerError(msg)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func (b *BlogRepo) Get(id string) (*core.Blog, *custom_errors.CustomError) {
	c := b.DB.Collection(constants.MongoCollections.BLOGS)
	blog := new(core.Blog)
	// convert the provided id from hex string -> mongo objectid
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		msg := "Object id is not valid"
		log.Println(msg)
		return nil, custom_errors.BadRequestError(msg)
	}
	filter := CreateBsonObject(map[string]interface{}{
		"_id": _id,
	})
	if err := c.FindOne(context.TODO(), filter).Decode(blog); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			msg := "Object with given id not found"
			log.Println(msg, err)
			return nil, custom_errors.NotFoundError(msg)
		} else {
			msg := "Cannot find the document"
			log.Println(msg, err)
			return nil, custom_errors.InternalServerError(msg)
		}
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
