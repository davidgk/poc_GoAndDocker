package helpers

import (
	"clientForm3Api/src/models"
	"log"
	"net/http"
)

type MyError struct {
	Code  int
	Error error
}

func ManageApiError(err error, res *http.Response, body []byte) (*models.AccountData, *MyError, bool) {
	if err != nil || res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, &MyError{Code: res.StatusCode, Error: err}, true
	}
	return nil, nil, false
}
