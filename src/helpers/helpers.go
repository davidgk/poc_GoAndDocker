package helpers

import (
	"clientForm3Api/src/models"
	"encoding/json"
	"net/http"
)

type MyError struct {
	Code          int
	Error         error
	ServerMessage string
}

func ManageApiError(err error, res *http.Response, body []byte) (*models.AccountData, *MyError, bool) {
	if err != nil || res.StatusCode > 299 {
		myError := buildMyError(err, res, body)
		return nil, myError, true
	}
	return nil, nil, false
}

func buildMyError(err error, res *http.Response, body []byte) *MyError {
	myError := &MyError{Code: res.StatusCode, Error: err}
	if myError.Error == nil {
		var result map[string]string
		json.Unmarshal(body, &result)
		myError.ServerMessage = result["error_message"]
	}
	return myError
}
