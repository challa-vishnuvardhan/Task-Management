package models

type TaskDetails struct {
	TaskDetailID uint64 `gorm:"primaryKey;column:task_detail_id;type:bigint(20) unsigned;not null" json:"task_detail_id"`
	UserID       uint64 `gorm:"index:user_id;column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`
	TaskName     string `gorm:"column:task_name;type:text" json:"task_name"`
	StartDate    uint64 `gorm:"column:start_date;type:bigint(20) unsigned" json:"start_date"`
	DueDate      uint64 `gorm:"column:due_date;type:bigint(20) unsigned" json:"due_date"`
	Description  string `gorm:"column:description;type:text" json:"description"`
	Priority     string `gorm:"column:priority;type:enum('low','medium', 'high'); not null" json:"priority"`
	Status       string `gorm:"column:status;type:enum('to do','in progress', 'done'); not null" json:"status"`
	AuditModel
}

func (m *TaskDetails) TableName() string {
	return "task_details"
}
