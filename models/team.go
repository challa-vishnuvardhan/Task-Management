package models

type TeamDetails struct {
	TeamID    uint64 `gorm:"primaryKey;column:team_id;type:bigint(20) unsigned;not null" json:"team_id"`
	ProjectID uint64 `gorm:"index:project_id;column:project_id;type:bigint(20) unsigned;not null" json:"project_id"`
	TeamtName string `gorm:"column:team_name;type:text" json:"team_name"`
	AuditModel
}

func (m *TeamDetails) TableName() string {
	return "team_details"
}
