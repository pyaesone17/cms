package datastore

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/pyaesone17/blog/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostDB interface {
	Find(id string) (*models.Post, error)
	Get() ([]*models.Post, error)
	CreatePost(p *models.Post) error
	UpdatePost(p *models.Post) error
	AddCategory(p *models.Post, category *models.Category)
}

type postdb struct {
	client *mongo.Client
}

func NewPostDataStore(client *mongo.Client) PostDB {
	return postdb{client}
}

func (db postdb) Get() ([]*models.Post, error) {
	type m struct {
		Post *models.Post       `json:"post" bson:"post"`
		Blah string             `json:"blah"`
		ID   primitive.ObjectID `json:"_id" bson:"_id"`
	}

	collection := db.client.Database(Database).Collection(PostCollection)
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, errors.Wrap(err, "decode failed")
	}

	var posts []*models.Post

	for cursor.Next(context.TODO()) {
		var result m
		err := cursor.Decode(&result)
		if err != nil {
			return nil, errors.Wrap(err, "decode failed")
		}
		result.Post.ID = result.ID.Hex()
		posts = append(posts, result.Post)
	}

	return posts, nil
}

func (db postdb) Find(id string) (*models.Post, error) {
	type m struct {
		Post *models.Post       `json:"post" bson:"post"`
		Blah string             `json:"blah"`
		ID   primitive.ObjectID `json:"_id" bson:"_id"`
	}
	var result m

	objectID, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectID}}
	if err != nil {
		return nil, errors.Wrap(err, "parsed from hex to id failed")
	}

	collection := db.client.Database(Database).Collection(PostCollection)
	singleresult := collection.FindOne(context.TODO(), filter)
	if singleresult.Err() != nil {
		return nil, errors.Wrap(singleresult.Err(), "find data failed")
	}

	err = singleresult.Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "decode failed")
	}

	result.Post.ID = result.ID.Hex()
	return result.Post, nil
}

func (db postdb) CreatePost(post *models.Post) error {
	collection := db.client.Database(Database).Collection(PostCollection)
	insertResult, err := collection.InsertOne(context.TODO(), bson.D{
		{"post", post.GetSaveModel()},
	})

	if err != nil {
		return errors.Wrap(err, "insert failed")
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	post.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (db postdb) UpdatePost(post *models.Post) error {
	objectID, err := primitive.ObjectIDFromHex(post.ID)
	filter := bson.D{{"_id", objectID}}

	update := bson.D{
		{"$set", bson.D{{"post", post.GetSaveModel()}}},
	}

	collection := db.client.Database(Database).Collection(PostCollection)
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func (db postdb) AddCategory(p *models.Post, category *models.Category) {
	collection := db.client.Database(Database).Collection(PostCollection)
	filter := bson.D{{"id", p.ID}}

	update := bson.D{{"category_id", category.ID}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
