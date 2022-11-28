package service

import (
	"bytes"
	"clientForm3Api/src/helpers"
	"clientForm3Api/src/models"
	"encoding/json"
	"os"
	"strconv"
)

type AccountApiService struct{}

// I investigate and I'd use this value from environment variables,
// but I think godotenv in not part of std library , right ?
var ApiAddress = os.Getenv("ACCOUNTAPI_ADDR") + "/v1/organisation/accounts"

const AccountDeleteConfirmation = "Account Deleted"

func (a AccountApiService) CreateAccount(accountDataJsonString string) (*models.AccountData, *helpers.MyError) {
	var jsonStr = []byte(accountDataJsonString)
	body, err := helpers.ExecuteAPI("POST", ApiAddress, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	data := parseToAccountData(body)
	return data, nil
}

func (a AccountApiService) GetAccountById(accountID string) (*models.AccountData, *helpers.MyError) {
	body, err := helpers.ExecuteAPI("GET", ApiAddress+"/"+accountID, nil)
	if err != nil {
		return nil, err
	}
	data := parseToAccountData(body)
	return data, nil
}

func (a AccountApiService) DeleteById(accountID string) (string, *helpers.MyError) {
	accountData, errOnFind := a.GetAccountById(accountID)
	if errOnFind == nil {
		deleteUrl := ApiAddress + "/" + accountData.ID + "?version=" + strconv.Itoa(accountData.Version)
		_, err := helpers.ExecuteAPI("DELETE", deleteUrl, nil)
		if err != nil {
			return "FAIL", err
		}
		return AccountDeleteConfirmation, nil
	}
	return "Fail", &helpers.MyError{ServerMessage: "Cannot delete an account with id: " + accountID + ", Error: " + errOnFind.ServerMessage, Code: 500}
}

func parseToAccountData(body []byte) *models.AccountData {
	data := new(models.AccountFetchData)
	json.Unmarshal(body, data)
	return &data.Data
}
