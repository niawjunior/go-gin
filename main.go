package main

import (
	"log"

	db "github.com/niawjunior/gin-app/database"
	"github.com/niawjunior/gin-app/routers"
)

func main() {
	log.Println("Starting server..")
	db.Init()
	r := routers.Router()
	r.Run(":3001")
}
