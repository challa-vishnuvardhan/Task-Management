package service

import (
	"Task-Management/constants"
	d "Task-Management/db"
	"Task-Management/models"
	"Task-Management/resources"
	"context"
	"fmt"
	"net/http"
)

type TaskService interface {
	CreateTask(ctx context.Context, input resources.TaskDetails) resources.ServiceResult
}

type taskService struct {
	taskDb d.TaskRepo
}

func NewTaskService(taskDb d.TaskRepo) *taskService {
	return &taskService{taskDb: taskDb}
}

func (s taskService) CreateTask(ctx context.Context, input resources.TaskDetails) resources.ServiceResult {
	functionDesc := "CreateActualLaunchDate Service"
	fmt.Println("Enter" + functionDesc)

	taskDetails := models.TaskDetails{
		TaskID:      input.TaskID,
		UserID:      input.UserID,
		TaskName:    input.TaskName,
		Description: input.Description,
		StartDate:   input.StartDate,
		DueDate:     input.DueDate,
		Status:      input.Status,
		Priority:    input.Priority,
	}

	taskDetails.AuditModel.SetCreateDefault(input.CreatedBy, 1)
	err := s.taskDb.CreateTask(ctx, &taskDetails)
	if err != nil {
		fmt.Println("Exit" + functionDesc)
		return resources.ServiceResult{
			Code:              http.StatusInternalServerError,
			ServiceResultData: resources.ServiceResultData{},
			IsError:           true,
			Error:             resources.ServiceError{ErrorMsg: constants.InternalServerError},
		}

	}

	return resources.ServiceResult{
		Code:              http.StatusOK,
		IsError:           false,
		ServiceResultData: resources.ServiceResultData{Data: taskDetails.TaskID},
	}

}
