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
	blog := CreateBsonDObject(map[string]interface{}{
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
	filter := CreateBsonDObject(map[string]interface{}{
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

func (b *BlogRepo) Update(id string, payload core.CreateBlogDto) (*core.Blog, *custom_errors.CustomError) {
	_id, err := ConvertToOjectId(id)
	if err != nil {
		return nil, err
	}

	c := b.DB.Collection(constants.MongoCollections.BLOGS)
	filter := CreateBsonDObject(map[string]interface{}{
		"_id": _id,
	})
	update := bson.M{
		"title":       payload.Title,
		"body":        payload.Body,
		"description": payload.Description,
		"updated_at":  time.Now(),
	}

	if _, err := c.UpdateOne(context.TODO(), filter, bson.M{"$set": update}); err != nil {
		if err == mongo.ErrNoDocuments {
			msg := logNoDocuments(err)
			return nil, custom_errors.NotFoundError(msg)
		} else {
			msg := "Cannot find document"
			log.Println(msg, err)
			return nil, custom_errors.InternalServerError(msg)
		}
	} else {
		blog := new(core.Blog)
		if err := c.FindOne(context.TODO(), filter).Decode(blog); err != nil {
			msg := "Cannot find document"
			log.Println(msg, err)
			return nil, custom_errors.InternalServerError(msg)
		}
		return blog, nil
	}
}

func (b *BlogRepo) Delete(id string) *custom_errors.CustomError {
	c := b.DB.Collection(constants.MongoCollections.BLOGS)
	_id, err := ConvertToOjectId(id)
	if err != nil {

		return err
	}
	filter := bson.M{"_id": _id}
	if res, err := c.DeleteOne(context.TODO(), filter); err != nil {
		if err == mongo.ErrNoDocuments {
			msg := logNoDocuments(err)
			return custom_errors.NotFoundError(msg)
		} else {
			msg := "Cannot delete document"
			log.Println(msg, err)
			return custom_errors.InternalServerError(msg)
		}
	} else {
		if res.DeletedCount == 0 {
			msg := "Cannot delete document"
			log.Println(msg, err)
			return custom_errors.InternalServerError(msg)
		}
	}
	return nil
}

func CreateBsonDObject(data map[string]interface{}) bson.D {
	var bsonObject bson.D
	for key, val := range data {
		bsonObject = append(
			bsonObject,
			primitive.E{Key: key, Value: val},
		)
	}
	return bsonObject
}

func CreateBsonMObject(data map[string]interface{}) bson.M {
	var bsonObject bson.M
	for key, val := range data {
		bsonObject[key] = val
	}
	return bsonObject
}

func ConvertToOjectId(id string) (*primitive.ObjectID, *custom_errors.CustomError) {
	if _id, err := primitive.ObjectIDFromHex(id); err != nil {
		return nil, custom_errors.BadRequestError("Invalid object id")
	} else {
		return &_id, nil
	}
}

func logNoDocuments(err error) string {
	msg := "Document with given id not found"
	log.Println(msg, err)
	return msg
}
