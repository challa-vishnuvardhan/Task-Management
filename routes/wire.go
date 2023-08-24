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
	TaskDb d.TaskDb
)

func BuildDb() {
	TaskDb = d.NewTaskDb()
}
func BuildService() {
	TaskService = s.NewTaskService(TaskDb)
}
func BuildHandler() {
	BuildDb()
	BuildService()
	TaskHandler = h.NewTaskHandler(TaskService)
}
