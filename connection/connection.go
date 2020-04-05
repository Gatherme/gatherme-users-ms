package connection

import (
	"errors"
	"log"
	"time"

	"github.com/Gatherme/gatherme-users-ms/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// INFO - to connect mongo
var INFO = &mgo.DialInfo{
	Addrs:    []string{"127.0.0.1:27017"},
	Timeout:  60 * time.Second,
	Database: "admin",
	Username: "admin",
	Password: "admin",
}

// DBNAME the name of the DB instance
const DBNAME = "user_db"

// DOCNAME the name of the document
const DOCNAME = "user"

var db *mgo.Database

// COLLECTION - name collection on Mongo
const (
	COLLECTION = "users"
)

// Insert - Insert a Shopping
func Insert(shopping model.User) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()

	shopping.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(shopping)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// FindByID - ...
func FindByID(id string) (model.User, error) {
	var shopping model.User
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return shopping, err
	}

	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return shopping, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&shopping)
	return shopping, err
}

// Update - ..
func Update(shopping model.User) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	err = c.UpdateId(shopping.ID, &shopping)
	return err
}

// FindByUser - ...
func FindByUsername(idUser string) ([]model.User, error) {
	var shoppings []model.User
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return shoppings, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.Find(bson.M{"username": idUser}).All(&shoppings)
	return shoppings, err
}

// Delete - ...
func Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return err
	}
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.RemoveId(oid)
	return err
}