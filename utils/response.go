package utils

import (
	"encoding/json"
	"net/http"

	T "github.com/komefumi/gsu-web/types"
)

func MakeResponse(response T.StandardResponse, w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(response)
}
