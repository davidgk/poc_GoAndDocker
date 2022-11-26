package helpers

import (
	"io"
	"io/ioutil"
	"net/http"
)

func ExecuteAPI(httpType string, url string, buffer io.Reader) ([]byte, *MyError) {
	req, err1 := http.NewRequest(httpType, url, buffer)
	if err1 != nil {
		myError := MyError{Code: 500, Error: err1, ServerMessage: "Something fails when create request:" + err1.Error()}
		return nil, &myError
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err2 := client.Do(req)
	if err2 != nil {
		myError := MyError{Code: 500, Error: err2, ServerMessage: "Something fails when resolve the request:" + err2.Error()}
		return nil, &myError
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	myError := ManageApiError(err, res, body)
	return body, myError
}
