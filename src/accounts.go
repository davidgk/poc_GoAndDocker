package src

import (
	er "clientForm3Api/src/helpers"
	"clientForm3Api/src/models"
	"encoding/json"
	"io"
	"net/http"
)

const API_ADDRESS = "http://localhost:8080/v1/organisation/accounts"

func Fetch(accountID string) (*models.AccountData, *er.MyError) {
	res, err := http.Get(API_ADDRESS + "/" + accountID)
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	accountData, myError, done := er.ManageApiError(err, res, body)
	if done {
		return accountData, myError
	}
	data := parseToAccountData(body)
	return data, nil
}

func parseToAccountData(body []byte) *models.AccountData {
	data := new(models.AccountFetchData)
	json.Unmarshal(body, data)
	return &data.Data
}
