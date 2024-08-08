package main

import (
	"os"

	"apsdatabase/db"
	"apsdatabase/server"
	"github.com/joho/godotenv"
)

func init() {
	if len(os.Getenv("HOST")) == 0 {
		godotenv.Load()
	}
}

func init() {
	db.Init()
}

func main() {
	server.Router()
}
