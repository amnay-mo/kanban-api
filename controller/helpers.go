package controller

import (
	"encoding/json"
	"net/http"

	"github.com/amnay-mo/kanban-api/model"
)

func decodeUser(r *http.Request) (*model.User, error) {
	user := &model.User{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func decodeTask(r *http.Request) (*model.Task, error) {
	todo := &model.Task{}
	todo.Status = 1
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
