package main

import (
	"log"

	db "github.com/niawjunior/go-gin/database"
	"github.com/niawjunior/go-gin/routers"
)

func main() {
	log.Println("Starting server..")

	db.Init()
	defer db.CloseDB()

	r := routers.Router()

	err := r.Run(":3001")

	if err != nil {
		panic(err)
	}

}
