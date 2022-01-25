package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Fight(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	log.Printf("Request: %v , Params: %v , \n", request, params)
}
