package db

import (
	"Task-Management/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ProjectRepo interface {
	CreateProject(ctx context.Context, projectData *models.ProjectDetails) error
	UpdateProject(ctx context.Context, projectData *models.ProjectDetails) error
	DeleteProject(ctx context.Context, projectId uint64) error
	GetProjectsByUser(ctx context.Context, userId uint64)
}

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepo() *projectRepo {
	return &projectRepo{
		db: DbClient,
	}
}

func (d projectRepo) CreateProject(ctx context.Context, projectData *models.ProjectDetails) error {
	functionDesc := "CreateProject DB"
	fmt.Println("Enter" + functionDesc)

	db := d.db

	memQuery := db.Debug().Create(&projectData)
	if err := memQuery.Error; err != nil {
		fmt.Println("Exit" + functionDesc)
		return err
	}
	fmt.Println("Exit" + functionDesc)
	return nil
}
