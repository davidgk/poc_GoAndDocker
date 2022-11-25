package service

import (
	"bytes"
	er "clientForm3Api/src/helpers"
	"clientForm3Api/src/models"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AccountApiService struct{}

const ApiAddress = "http://localhost:8080/v1/organisation/accounts"
const AccountDeleteConfirmation = "Account Deleted"

func (a AccountApiService) CreateAccount(accountDataJsonString string) (*models.AccountData, *er.MyError) {
	var jsonStr = []byte(accountDataJsonString)
	req, err := http.NewRequest("POST", ApiAddress, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
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

func (a AccountApiService) DeleteById(accountID string) (string, *er.MyError) {
	accountData, errOnFind := a.GetAccountById(accountID)
	if errOnFind == nil {
		deleteUrl := ApiAddress + "/" + accountData.ID + "?version=" + strconv.Itoa(accountData.Version)
		req, err := http.NewRequest("DELETE", deleteUrl, nil)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		_, myError, done := er.ManageApiError(err, resp, body)
		if done {
			return "Fail", myError
		}
		return AccountDeleteConfirmation, nil
	}
	return "Fail", &er.MyError{ServerMessage: "Cannot delete an account with id: " + accountID + ", Error: " + errOnFind.ServerMessage, Code: 500}
}

func parseToAccountData(body []byte) *models.AccountData {
	data := new(models.AccountFetchData)
	json.Unmarshal(body, data)
	return &data.Data
}
