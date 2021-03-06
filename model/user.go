package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Users - Model
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"username" json:"username"`
	Name        string        `bson:"name" json:"name"`
	Email       string        `bson:"email" json:"email"`
	Picture     string        `bson:"picture" json:"picture"`
	Description string        `bson:"description" json:"description"`
	Gender      string        `bson:"gender" json:"gender"`
	Age         int           `bson:"age" json:"age"`
	City        string        `bson:"city" json:"city"`
	Likes       []string      `bson:"likes" json:"likes"`
	Communities []int         `bson:"communities" json:"communities"`
	Activities  []int         `bson:"activities" json:"activities"`
	Gathers     []string      `bson:"gathers" json:"gathers"`
}

// UserID - for request
type UserID struct {
	ID string `json:"id"`
}
