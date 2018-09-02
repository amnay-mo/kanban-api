package model

// TaskDatastore is the datastore interface for tasks
type TaskDatastore interface {
	GetTasks() ([]Task, error)
	SaveTask(*Task) error
	DeleteTask(string) error
	UpdateTaskStatus(string, int) error
	SaveUser(*User) error
	GetUserByEmail(string) (*User, error)
}

var taskDatastore TaskDatastore

// SetDatastore sets the model's datastore
func SetDatastore(ds TaskDatastore) {
	taskDatastore = ds
}
