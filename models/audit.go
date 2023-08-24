package models

type AuditModel struct {
	IsActive   int16  `gorm:"column:is_active;type:tinyint(1);" json:"is_active"`
	CreatedBy  string `gorm:"column:created_by;type:varchar(320);not null" json:"created_by"`
	CreatedAt  uint64 `gorm:"column:created_at;type:bigint(20);autoCreateTime" json:"created_at"`
	ModifiedBy string `gorm:"column:modified_by;type:varchar(320)" json:"modified_by"`
	ModifiedAt uint64 `gorm:"column:modified_at;type:bigint(20);autoUpdateTime" json:"modified_at"`
}

func (m *AuditModel) SetCreateDefault(user_email string, is_active int16) {
	m.IsActive = is_active
	m.CreatedBy = user_email
}

func (m *AuditModel) SetUpdateDefault(user_email string, is_active int16) {
	m.IsActive = is_active
	m.ModifiedBy = user_email
}

