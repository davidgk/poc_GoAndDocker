package account_tests

import (
	"clientForm3Api/src"
	"clientForm3Api/test"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var a *int64

func TestDeleteAccount(t *testing.T) {
	fmt.Println("Starting Test of Account - Delete ")
	msg := sunnyDeleteByIdTest(t)
	msgs := []string{"when want to Delete an Account and it exists " + msg}

	msg = deleteByIdSendIdNullError(t)
	msgs = append(msgs, "and id as param is empty "+msg)

	test.PrintTestsMessages(msgs)
}

func sunnyDeleteByIdTest(t *testing.T) string {
	msg := "and id Exists"
	// Given
	accountDataSunny, _ := test.CreateDataToCreateAccount()
	data, _ := json.Marshal(accountDataSunny)
	// When
	aAccountCreated, _ := src.Create(string(data))
	message, _ := src.Delete(aAccountCreated.ID)
	require.Equal(t, message, "Account Deleted", "the account Should be deleted")
	_, err := src.Fetch(aAccountCreated.ID)
	require.Equal(t, err.Code, 404, "the account was not found ")
	require.Equal(t, err.ServerMessage, "record "+aAccountCreated.ID+" does not exist", "the account was not found ")
	msg += " should be found"
	return msg
}

func deleteByIdSendIdNullError(t *testing.T) string {
	id := ""
	_, err := src.Delete(id)
	require.Equal(t, err.Code, 500, "error should be 500 ")
	require.Equal(t, err.ServerMessage, "Id param should not be null", "should valid error message about attributes")
	msg := " should throw 500 error"
	return msg
}