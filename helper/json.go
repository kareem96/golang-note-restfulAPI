package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
    writer.Header().Add("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
    err := encoder.Encode(response)
    if err != nil {
        log.Printf("Error encoding response: %v", err)
    }
    PanicIfError(err)
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
    decoder := json.NewDecoder(request.Body)
    err := decoder.Decode(result)
    if err != nil {
        log.Printf("Error decoding request body: %v", err)
    }
    PanicIfError(err)
}

// func WriteToResponseBody(writer http.ResponseWriter, response interface{})  {
// 	writer.Header().Add("Content-Type", "application/json")
// 	encoder := json.NewEncoder(writer)
// 	err := encoder.Encode(response)
// 	PanicIfError(err)
// }

// func ReadFromRequestBody(request *http.Request, result interface{})  {
// 	decoder := json.NewDecoder(request.Body)
// 	err := decoder.Decode(result)
// 	PanicIfError(err)
// }