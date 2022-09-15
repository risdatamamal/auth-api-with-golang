package main

import (
	"auth-api-with-golang/database"
	"auth-api-with-golang/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
