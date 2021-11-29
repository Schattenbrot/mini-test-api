package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBModel
type DBModel struct {
	DB *mongo.Database
}

// InsertPost is the method for inserting posts.
func (m *DBModel) InsertPost(post Post) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("posts")

	oid, err := collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	result := oid.InsertedID.(primitive.ObjectID)
	return &result, nil
}
