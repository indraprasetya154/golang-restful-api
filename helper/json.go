package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writter http.ResponseWriter, response interface{}) {
	writter.Header().Add("Content-type", "Application/json")
	encoder := json.NewEncoder(writter)
	err := encoder.Encode(response)
	PanicIfError(err)
}
