package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	T "github.com/komefumi/gsu-web/types"
)

func StringIsOfLength(candidate string, min uint, max uint) error {
	givenLength := uint(len(candidate))
	if min != 0 && givenLength < min {
		return errors.New(fmt.Sprintf("Error: Given string '%v' is only of length %v, must be at least %v", candidate, givenLength, min))
	}
	if max != 0 && givenLength > max {
		return errors.New(fmt.Sprintf("Error: Given string '%v' is of length %v, must be at max %v", candidate, givenLength, min))
	}

	return nil
}

func ValidateBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errors := []string{"Error: Make sure request is given with a valid request body"}
		json.NewEncoder(w).Encode(T.StandardResponse{Success: 0, Data: nil, Errors: errors})
		return nil, err
	}

	return reqBody, nil
}
