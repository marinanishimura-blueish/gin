package main

import (
	"github.com/gin/backend/db"
)

func main() {
	db.Init()
	// server.Init()
	db.Close()
}
