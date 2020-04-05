package model

import (
	"gopkg.in/mgo.v2/bson"
)


type Pleasures struct{
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Category    string        `bson:"category" json:"category"`
	Name       	string        `bson:"email" json:"email"`
	
}
