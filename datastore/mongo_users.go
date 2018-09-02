package datastore

import (
	"log"

	"github.com/amnay-mo/kanban-api/model"

	"github.com/globalsign/mgo/bson"
)

const dbNameUsers = "kanban"
const collNameUsers = "users"

// SaveUser adds a new user
func (mt MongoTasks) SaveUser(user *model.User) error {
	coll := mt.sess.DB(dbNameUsers).C(collNameUsers)
	err := coll.Insert(user)
	if err != nil {
		log.Println("Could not add user!")
		return err
	}
	return nil
}

// GetUserByEmail finds a user by email
func (mt MongoTasks) GetUserByEmail(email string) (*model.User, error) {
	coll := mt.sess.DB(dbNameUsers).C(collNameUsers)
	var user model.User
	err := coll.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
