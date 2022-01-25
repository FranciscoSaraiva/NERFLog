package main

import (
	"log"
	"nerflog/config"
	"nerflog/errors"
	"nerflog/handlers"
	"net/http"
	"os"
)

func main() {
	log.Println("getting server")

	router := config.NewServerAPI()

	router.GET("/parse/:character", handlers.Parse)
	router.GET("/raid/:fightid", handlers.Fight)

	err := http.ListenAndServe("localhost:8080", router)
	if errors.CheckCommonError(err) {
		os.Exit(1)
	}

	log.Println("[ServerAPI] ServerAPI is running in http://localhost:8080/")
	os.Exit(0)
}
