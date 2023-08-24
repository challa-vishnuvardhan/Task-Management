package main

import (
	"Task-Management/config"
	"Task-Management/db"
	"Task-Management/routes"
)

func main() {
	config.IntializeEnv()
	db.DbConnect()
	routes.StartRoute()
}
