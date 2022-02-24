package main

import (
	"github.com/alexguidi/api-go-gin/database"
	"github.com/alexguidi/api-go-gin/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
