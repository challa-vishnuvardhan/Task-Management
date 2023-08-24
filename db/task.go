package db

import (
	"Task-Management/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type TaskDb interface {
	CreateTask(ctx context.Context, input *models.TaskDetails) error
}

type taskDb struct {
	db *gorm.DB
}

func NewTaskDb() *taskDb {
	return &taskDb{
		db: DbClient,
	}
}

func (d *taskDb) CreateTask(ctx context.Context, input *models.TaskDetails) error {
	functionDesc := "CreateActualLaunchDate DB"
	fmt.Println("Enter" + functionDesc)

	db := d.db

	memQuery := db.Debug().WithContext(ctx).Create(&input)
	if err := memQuery.Error; err != nil {
		fmt.Println("Exit" + functionDesc)
		return err
	}
	fmt.Println("Exit" + functionDesc)
	return nil

}
