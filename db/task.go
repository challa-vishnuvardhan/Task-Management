package db

import (
	"Task-Management/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type TaskRepo interface {
	CreateTask(ctx context.Context, input *models.TaskDetails) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo() *taskRepo {
	return &taskRepo{
		db: DbClient,
	}
}

func (d *taskRepo) CreateTask(ctx context.Context, input *models.TaskDetails) error {
	functionDesc := "CreateTask DB"
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
