package models

type UserDetails struct {
	UserID    uint64 `gorm:"primaryKey;column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`
	FirstName string `gorm:"column:first_name;type:varchar(100)" json:"first_name"`
	LastName  string `gorm:"column:last_name;type:varchar(100)" json:"last_name"`
	EmailID   string `gorm:"column:email_id;type:varchar(200)" json:"email_id"`
	AuditModel
}

func (m *UserDetails) TableName() string {
	return "user_details"
}
