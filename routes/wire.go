package routes

import (
	d "Task-Management/db"
	h "Task-Management/handler"
	s "Task-Management/service"
)

var (
	TaskHandler h.TaskHandler
)
var (
	TaskService s.TaskService
)

var (
	TaskRepo d.TaskRepo
)

func BuildDb() {
	TaskRepo = d.NewTaskRepo()
}
func BuildService() {
	TaskService = s.NewTaskService(TaskRepo)
}
func BuildHandler() {
	BuildDb()
	BuildService()
	TaskHandler = h.NewTaskHandler(TaskService)
}
