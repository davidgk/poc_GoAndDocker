package account_tests

import (
	"clientForm3Api/src"
	"clientForm3Api/test"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFetchAccount(t *testing.T) {
	fmt.Println("Starting Test of Account - Fetch ")
	msg := sunnyGetByIdTest(t)
	msgs := []string{"when want to Fetch an Account by Id " + msg}

	msg = notFoundAccountCheck(t)
	msgs = append(msgs, "and it doesn't exist "+msg)

	msg = sendIdNullError(t)
	msgs = append(msgs, "and id as param is empty "+msg)

	test.PrintTestsMessages(msgs)
}

func sendIdNullError(t *testing.T) string {
	id := ""
	_, err := src.Fetch(id)
	require.Equal(t, err.Code, 500, "error shoyld be 500 ")
	require.Equal(t, err.ServerMessage, "Id param should not be null", "should valid error message about attributes")
	msg := " should throw 500 error"
	return msg
}

func notFoundAccountCheck(t *testing.T) string {
	id := uuid.New()
	_, err := src.Fetch(id.String())
	require.Equal(t, err.Code, 404, "the account was not found ")
	require.Equal(t, err.ServerMessage, "record "+id.String()+" does not exist", "the account was not found ")
	msg := " should throw 400 error"
	return msg
}

func sunnyGetByIdTest(t *testing.T) string {
	msg := "and id Exists"
	// Given
	accountDataSunny, _ := test.CreateDataToCreateAccount()
	data, _ := json.Marshal(accountDataSunny)
	// When
	aAccountCreated, _ := src.Create(string(data))
	aAccountFound, _ := src.Fetch(aAccountCreated.ID)
	require.Equal(t, aAccountFound.ID, aAccountCreated.ID, "the account created not contain the same ID ")
	msg += " should be found"
	src.Delete(aAccountCreated.ID)
	return msg
}
