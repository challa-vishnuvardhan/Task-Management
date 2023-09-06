package resources

import (
	"Task-Management/models"
)

type TaskDetails struct {
	TaskID      uint64 `json:"task_id"`
	UserID      uint64 `json:"user_id"`
	TaskName    string `json:"task_name"`
	StartDate   uint64 `json:"start_date"`
	DueDate     uint64 `json:"due_date"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	models.AuditModel
}
