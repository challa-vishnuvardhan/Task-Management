package main

import (
	"Task-Management/config"
	"Task-Management/db"
	"Task-Management/routes"
)

func main() {
	//Intializing the Envinormental keys
	config.IntializeEnv()

    // connecting to db 
	db.DbConnect()

	//start
	routes.StartRoute()
}
