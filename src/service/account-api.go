package service

import (
	"bytes"
	er "clientForm3Api/src/helpers"
	"clientForm3Api/src/models"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type AccountApiService struct{}

const ApiAddress = "http://localhost:8080/v1/organisation/accounts"

func (a AccountApiService) CreateAccount(accountDataJsonString string) (*models.AccountData, *er.MyError) {
	var jsonStr = []byte(accountDataJsonString)
	req, err := http.NewRequest("POST", ApiAddress, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	accountData, myError, done := er.ManageApiError(err, resp, body)
	if done {
		return accountData, myError
	}
	data := parseToAccountData(body)
	return data, nil
}

func (a AccountApiService) GetAccountById(accountID string) (*models.AccountData, *er.MyError) {
	res, err := http.Get(ApiAddress + "/" + accountID)
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
