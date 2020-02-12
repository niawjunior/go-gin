package main

import (
	"log"

	database "github.com/niawjunior/go-gin/database"
	"github.com/niawjunior/go-gin/routers"
)

func main() {
	log.Println("Starting server..")

	database.Init()
	defer database.DB.Close()

	r := routers.Router()

	err := r.Run(":3001")

	if err != nil {
		panic(err)
	}

}
