package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"nerflog/constants"
	"nerflog/errors"
	"nerflog/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Parse(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var character string = params.ByName("character")
	log.Printf("Character: %v", character)

	//message
	response := models.ApiResponse{
		Message: constants.GetCharacterParseMessage + character,
		Status:  http.StatusOK,
	}

	//response
	jsonResp, err := json.Marshal(response)
	errors.CheckCommonError(err)
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%s", jsonResp)

	return
}
