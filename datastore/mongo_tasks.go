package datastore

import (
	"log"

	"github.com/amnay-mo/kanban-api/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoTasks is a mongo datastore
type MongoTasks struct {
	sess *mgo.Session
}

const dbNameTasks = "kanban"
const collNameTask = "tasks"

// NewMongoTasks returns a MongoTasks instance
func NewMongoTasks(connstring string) (MongoTasks, error) {
	sess, err := mgo.Dial(connstring)
	if err != nil {
		log.Fatal(err)
	}
	return MongoTasks{sess: sess}, nil
}

// GetTasks return all tasks
func (mt MongoTasks) GetTasks() ([]model.Task, error) {
	coll := mt.sess.DB(dbNameTasks).C(collNameTask)
	cursor := coll.Find(nil)
	tasks := []model.Task{}
	cursor.All(&tasks)
	return tasks, nil
}

// SaveTask saves a new task
func (mt MongoTasks) SaveTask(task *model.Task) error {
	coll := mt.sess.DB(dbNameTasks).C(collNameTask)
	err := coll.Insert(task)
	return err
}

// DeleteTask deletes a task
func (mt MongoTasks) DeleteTask(taskID string) error {
	coll := mt.sess.DB(dbNameTasks).C(collNameTask)
	err := coll.Remove(bson.M{"_id": taskID})
	return err
}

// UpdateTaskStatus updates a task's status
func (mt MongoTasks) UpdateTaskStatus(taskID string, status int) error {
	coll := mt.sess.DB(dbNameTasks).C(collNameTask)
	err := coll.Update(bson.M{"_id": taskID}, bson.M{"$set": bson.M{"status": status}})
	return err
}
