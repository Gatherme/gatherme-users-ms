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
	Addrs:   []string{"gatherme-users-db:27019"},
	Timeout: 60 * time.Second,
}

// DBNAME the name of the DB instance
const DBNAME = "user_db"

// DOCNAME the name of the document
const DOCNAME = "user"
const DOCNAME_P = "like"

var db *mgo.Database

// COLLECTION - name collection on Mongo
const (
	COLLECTION = "users"
)

// Insert - Insert a user
func InsertUser(user model.User) error {
	session, err := mgo.Dial("gatherme-users-db:27019")
	if err != nil {
		log.Printf("Error de conexion con Mongo")
		log.Fatal(err)
		log.Fatalln("mongo err")
		return err
	}
	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)
	//user.ID = bson.NewObjectId()

	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return err
	}

	index2 := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index2)
	if err != nil {
		return err
	}

	err = c.Insert(user)

	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

// Insert - Insert a like
func InsertLike(like model.Like) error {
	session, err := mgo.Dial("gatherme-users-db:27019")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME_P)

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return err
	}
	//like.ID = bson.NewObjectId()
	err = c.Insert(like)

	if err != nil {
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

	session, err := mgo.Dial("gatherme-users-db:27019")
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
func FindLikeByID(id string) (model.Like, error) {
	var like model.Like
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return like, err
	}

	session, err := mgo.Dial("gatherme-users-db:27019")
	if err != nil {
		log.Fatal(err)
		return like, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME_P)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&like)
	return like, err
}

// Update User - ..
func UpdateUser(user model.User) error {
	session, err := mgo.Dial("gatherme-users-db:27019")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	err = c.UpdateId(user.ID, &user)
	return err
}

// Update Like - ..
func UpdateLike(like model.Like) error {
	session, err := mgo.Dial("gatherme-users-db:27019")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME_P)
	err = c.UpdateId(like.ID, &like)
	return err
}

// Find User by username - ...
func FindByUsername(idUser string) ([]model.User, error) {
	log.Printf("Find by username")
	var users []model.User
	session, err := mgo.Dial("127.0.0.1:27017")
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
func FindByEmail(email string) ([]model.User, error) {
	log.Printf(email)
	var users []model.User
	session, err := mgo.Dial("gatherme-users-db:27019")
	if err != nil {
		log.Fatal(err)
		return users, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.Find(bson.M{"email": email}).All(&users)

	return users, err
}

// Find User by username - ...
func FindLikesByCategory(category string) ([]model.Like, error) {
	var likes []model.Like
	session, err := mgo.Dial("gatherme-users-db:27019")
	if err != nil {
		log.Fatal(err)
		return likes, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME_P)

	err = c.Find(bson.M{"category": category}).All(&likes)
	return likes, err
}

// Delete User by id- ...
func DeleteUser(id string) error {
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return err
	}
	session, err := mgo.Dial("gatherme-users-db:27019")
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

// Delete Like by id- ...
func DeleteLike(id string) error {
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return err
	}
	session, err := mgo.Dial("gatherme-users-db:27019")
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
