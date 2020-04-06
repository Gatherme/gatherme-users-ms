package model

import (
	"gopkg.in/mgo.v2/bson"
)


type Pleasure struct{
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Category    string        `bson:"category" json:"category"`
	Name       	string        `bson:"name" json:"name"`	
}

// UserID - for request
type PleasureID struct {
	ID string `json:"id"`
}
