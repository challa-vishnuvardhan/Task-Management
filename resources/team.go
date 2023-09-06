package resources

import "Task-Management/models"

type TeamDetails struct {
	TeamID    uint64 `json:"team_id"`
	ProjectID uint64 `json:"project_id"`
	TeamtName string `json:"team_name"`
	models.AuditModel
}
