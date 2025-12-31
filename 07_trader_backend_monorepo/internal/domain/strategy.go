package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Strategy struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id, omitempty"` // In MongoDB, each document must contain a unique _id field.
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
}
