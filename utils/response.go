package utils

import (
	"encoding/json"
	"net/http"

	"github.com/dreddsa5dies/httprestapient/model"
)

// Return - encode the object to JSON and then return it to the client
func Return(w http.ResponseWriter, status bool, code int, err error, data interface{}) {

	response := model.Response{
		Status: status,
		Code:   code,
		Error:  "",
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
