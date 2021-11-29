package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Models is the wrapper fot the database.
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *mongo.Database) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Post is the type for posts
type Post struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Text string             `json:"text,omitempty" bson:"text,omitempty"`
}
