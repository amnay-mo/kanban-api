package controller

import (
	"net/http"

	"github.com/amnay-mo/kanban-api/model"
	"github.com/amnay-mo/kanban-api/utils"

	"github.com/julienschmidt/httprouter"
)

// GetTasks return all what's left to do
func GetTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tasks, _ := model.GetTasks()
	response := map[string][]model.Task{
		"tasks": tasks,
	}
	utils.Jsonify(w, r, response, http.StatusOK)
}

// AddTask adds a new ToDo
func AddTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	task, _ := decodeTask(r)
	newID, err := model.AddTask(task)
	if err != nil {
		response := map[string]string{
			"error": "could not save task to the datastore :(",
		}
		utils.Jsonify(w, r, response, http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"id": newID,
	}
	utils.Jsonify(w, r, response, http.StatusOK)
}

// DeleteTask adds a new ToDo
func DeleteTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	taskID := p.ByName("task_id")
	err := model.DeleteTask(taskID)
	if err != nil {
		response := map[string]string{
			"error": "No such task id",
		}
		utils.Jsonify(w, r, response, http.StatusNotFound)
		return
	}
	utils.Jsonify(w, r, "", http.StatusOK)
}

// UpdateTask updates a task
func UpdateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	taskID := p.ByName("task_id")
	task, err := decodeTask(r)
	if err != nil {
		response := map[string]string{
			"error": "Bad payload",
		}
		utils.Jsonify(w, r, response, http.StatusNotFound)
		return
	}
	newStatus := task.Status
	err = model.UpdateTask(taskID, newStatus)
	if err != nil {
		response := map[string]string{
			"error": "No such task id",
		}
		utils.Jsonify(w, r, response, http.StatusNotFound)
		return
	}
	utils.Jsonify(w, r, "", http.StatusOK)
}
