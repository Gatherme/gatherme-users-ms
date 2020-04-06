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
const DOCNAME_P = "pleasure"

var db *mgo.Database

// COLLECTION - name collection on Mongo
const (
	COLLECTION = "users"
)


// Insert - Insert a user
func InsertUser(user model.User) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()

	user.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(user)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Insert - Insert a pleasure
func InsertPleasure(pleasure model.Pleasure) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()

	pleasure.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME_P).Insert(pleasure)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Find user by ID - ...	
func FindUserByID(id string) (model.User, error) {
	var user model.User
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return user, err
	}

	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&user)
	return user, err
}


// Find user by ID - ...	
func FindPleasureByID(id string) (model.Pleasure, error) {
	var pleasure model.Pleasure
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return pleasure, err
	}

	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return pleasure, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME_P)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&pleasure)
	return pleasure, err
}

// Update User - ..
func UpdateUser(user model.User) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	err = c.UpdateId(user.ID, &user)
	return err
}


// Update Pleasure - ..
func UpdatePleasure(pleasure model.Pleasure) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME_P)
	err = c.UpdateId(pleasure.ID, &pleasure)
	return err
}

// Find User by username - ...
func FindByUsername(idUser string) ([]model.User, error) {
	var users []model.User
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return users, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.Find(bson.M{"username": idUser}).All(&users)
	return users, err
}

// Find User by username - ...
func FindPleasuresByCategory(category string) ([]model.Pleasure, error) {
	var pleasures []model.Pleasure
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return pleasures, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME_P)

	err = c.Find(bson.M{"category": category}).All(&pleasures)
	return pleasures, err
}


// Delete User by id- ...
func DeleteUser(id string) error {
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

// Delete Pleasure by id- ...
func DeletePleasure(id string) error {
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
	c := session.DB(DBNAME).C(DOCNAME_P)

	oid := bson.ObjectIdHex(id)
	err = c.RemoveId(oid)
	return err
}