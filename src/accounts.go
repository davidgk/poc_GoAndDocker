package src

import (
	er "clientForm3Api/src/helpers"
	"clientForm3Api/src/models"
	"clientForm3Api/src/service"
)

var accService = service.AccountApiService{}

func Fetch(accountID string) (*models.AccountData, *er.MyError) {
	if validateNotNullParam(accountID) {
		accountData, myError := accService.GetAccountById(accountID)
		if myError != nil {
			return accountData, myError
		}
		return accountData, nil
	}
	myError := &er.MyError{ServerMessage: "Id param should not be null", Code: 500}
	return nil, myError
}

func validateNotNullParam(param string) bool {
	return len(param) > 0
}

func Create(accountDataJsonString string) (*models.AccountData, *er.MyError) {
	if validateNotNullParam(accountDataJsonString) {
		accountData, myError := accService.CreateAccount(accountDataJsonString)
		if myError != nil {
			return accountData, myError
		}
		return accountData, nil
	}
	myError := &er.MyError{ServerMessage: "Param should not be null", Code: 500}
	return nil, myError
}
