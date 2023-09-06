package resources

import "Task-Management/models"

type UserDetails struct {
	UserID    uint64 ` json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	EmailID   string `json:"email_id"`
	models.AuditModel
}
