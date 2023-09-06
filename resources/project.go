package resources

import "Task-Management/models"

type ProjectDetails struct {
	ProjectID   uint64 `json:"project_id"`
	ProjectName string `json:"project_name"`
	models.AuditModel
}
