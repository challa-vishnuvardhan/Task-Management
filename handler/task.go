package handler

import (
	"Task-Management/constants"
	"Task-Management/resources"
	s "Task-Management/service"
	"Task-Management/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type TaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
}
type taskHandler struct {
	taskService s.TaskService
}

func NewTaskHandler(taskService s.TaskService) *taskHandler {
	return &taskHandler{taskService: taskService}
}

func (h *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	functionDesc := "CreateActualLaunchDate Handler"
	fmt.Println("Enter" + functionDesc)
	ctx := r.Context()

	var input resources.TaskDetails

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		utils.JsonWriterString(w, http.StatusBadRequest, `{"status": "Not Ok","msg":"Error in decode the data"}`)
		fmt.Println("Exit" + functionDesc)
		return
	}
	result := h.taskService.CreateTask(ctx, input)
	if result.IsError {
		fmt.Println(result.Error)
		utils.JsonWriterString(w, http.StatusInternalServerError, constants.InternalServerError)
		fmt.Println("Exit" + functionDesc)
		return
	} else {
		utils.WriteResponseStruct(ctx, w, result.Code, result.Data)
	}
	fmt.Println("Exit" + functionDesc)
}
