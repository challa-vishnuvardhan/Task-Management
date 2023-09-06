package models

type ProjectDetails struct {
	ProjectID   uint64 `gorm:"primaryKey;column:project_id;type:bigint(20) unsigned;not null" json:"project_id"`
	ProjectName string `gorm:"column:project_name;type:text" json:"project_name"`
	AuditModel
}

func (m *ProjectDetails) TableName() string {
	return "project_details"
}
