package main

import "main/bayarXYZ/database"

func main() {
	db, _ := database.ConnectToDB()

	callRouter(db)
}
