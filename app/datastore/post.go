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
	Get(id string) (models.PostFractal, error)
	CreatePost(p models.PostFractal)
	AddCategory(p models.PostFractal, category *models.Category)
}

type postdb struct {
	client *mongo.Client
}

func NewPostDataStore(client *mongo.Client) PostDB {
	return postdb{client}
}

func (db postdb) Get(id string) (models.PostFractal, error) {
	type m struct {
		Post models.Post        `json:"post" bson:"post"`
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

func (db postdb) CreatePost(p models.PostFractal) {
	collection := db.client.Database(Database).Collection(PostCollection)
	insertResult, err := collection.InsertOne(context.TODO(), bson.D{
		{"post", p},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func (db postdb) AddCategory(p models.PostFractal, category *models.Category) {
	collection := db.client.Database(Database).Collection(PostCollection)
	filter := bson.D{{"id", p.(*models.Post).ID}}

	update := bson.D{{"category_id", category.ID}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
