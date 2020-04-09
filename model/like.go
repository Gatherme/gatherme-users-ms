package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Like struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Category string        `bson:"category" json:"category"`
	Name     string        `bson:"name" json:"name"`
}

// UserID - for request
type LikeID struct {
	ID string `json:"id"`
}
