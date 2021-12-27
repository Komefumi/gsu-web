package utils

import (
	"encoding/json"
	T "gsu-web/types"
	"net/http"
)

func MakeResponse(response T.StandardResponse, w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(response)
}
