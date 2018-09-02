package model

import uuid "github.com/satori/go.uuid"

// Task is just a task struct with its json mapping
type Task struct {
	ID     string `json:"id" bson:"_id"`
	Text   string `json:"text" bson:"text"`
	Status int    `json:"status" bson:"status"`
}

// GetTasks return all tasks
func GetTasks() ([]Task, error) {
	result, err := taskDatastore.GetTasks()
	return result, err
}

// AddTask adds something new
func AddTask(task *Task) (string, error) {
	task.ID = uuid.NewV4().String()
	err := taskDatastore.SaveTask(task)
	return task.ID, err
}

// DeleteTask deletes a task
func DeleteTask(taskID string) error {
	err := taskDatastore.DeleteTask(taskID)
	return err
}

// UpdateTask updates a task
func UpdateTask(taskID string, status int) error {
	err := taskDatastore.UpdateTaskStatus(taskID, status)
	return err
}
