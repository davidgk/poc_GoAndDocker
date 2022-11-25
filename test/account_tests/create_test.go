package account_tests

import (
	"clientForm3Api/src"
	"clientForm3Api/test"
	amf "clientForm3Api/test"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	fmt.Println("Starting Test of Account - Create ")
	msg := "and all values are Correct"
	// Given
	accountDataSunny, id := amf.CreateDataToCreateAccount()
	data, _ := json.Marshal(accountDataSunny)
	// When
	aAccountFound, _ := src.Create(string(data))
	// Then
	require.Equal(t, aAccountFound.ID, id, "the account created has fail")
	msg += " should be created"
	src.Delete(aAccountFound.ID)
	msgs := []string{"when want to Create an Account" + msg}
	test.PrintTestsMessages(msgs)
}

func TestCreateAccountFailsScenarios(t *testing.T) {
	fmt.Println("Starting Test of Account - Create - Fail Scenarios - when want to Create an Account ")
	msg := failScenarioSentSameRequestMoreThanOneTime(t)
	msgs := []string{"And same account posted twice" + msg}
	msg = failScenarioNotTypeAccountSent(t)
	msgs = append(msgs, "And type account is not send"+msg)
	msg = failScenarioNotIdSent(t)
	msgs = append(msgs, "And id is not send"+msg)
	msg = failScenarioNotOrganizationIdSent(t)
	msgs = append(msgs, "And Organization id is not send"+msg)
	msg = failScenarioNotAttributesSend(t)
	msgs = append(msgs, "And Attributes is not send"+msg)
	msg = failScenarioAttributesNotCountrySend(t)
	msgs = append(msgs, "And Country within is not send"+msg)
	msg = failScenarioAttributesNotCountrySend(t)
	msgs = append(msgs, "And Country within is not send"+msg)
	msg = failScenarioNotJsonStringSend(t)
	msgs = append(msgs, "And Country within is not send"+msg)
	test.PrintTestsMessages(msgs)
}

func failScenarioNotJsonStringSend(t *testing.T) string {
	// Given
	accountDataSunny := ""
	// When
	_, err := src.Create(accountDataSunny)
	// Then
	require.Equal(t, err.Code, 500, "should throw 500")
	errMessage := "Param should not be null"
	require.Equal(t, err.ServerMessage, errMessage, "should valid error message about null param")
	return " should throw 500 error and fail for null param"
}

func failScenarioSentSameRequestMoreThanOneTime(t *testing.T) string {
	// Given
	accountDataSunny, id := amf.CreateDataToCreateAccount()
	data, _ := json.Marshal(accountDataSunny)
	// When
	account, _ := src.Create(string(data))
	require.Equal(t, account.ID, id, "the account created has fail")
	_, err := src.Create(string(data))
	// Then
	require.Equal(t, err.Code, 409, "should throw 409")
	const errMessage = "Account cannot be created as it violates a duplicate constraint"
	require.Equal(t, err.ServerMessage, errMessage, "should valid error message about duplicate constraint")
	src.Delete(account.ID)
	return " should throw 409 error and fail for duplicated constraint"
}

func failScenarioAttributesNotCountrySend(t *testing.T) string {
	// Given
	accountDataSunny, _ := amf.CreateDataToCreateAccount()
	accountDataSunny.Data.Attributes.Country = nil
	data, _ := json.Marshal(accountDataSunny)
	// When
	_, err := src.Create(string(data))
	// Then
	require.Equal(t, err.Code, 400, "should throw 400")
	const countryErrMessage = "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body is required"
	require.Equal(t, err.ServerMessage, countryErrMessage, "should valid error message about attributes")
	return " should throw 400 error"
}

func failScenarioNotAttributesSend(t *testing.T) string {
	// Given
	accountDataSunny, _ := amf.CreateDataToCreateAccount()
	accountDataSunny.Data.Attributes = nil
	data, _ := json.Marshal(accountDataSunny)
	// When
	_, err := src.Create(string(data))
	// Then
	require.Equal(t, err.Code, 400, "should throw 400")
	require.Equal(t, err.ServerMessage, "validation failure list:\nvalidation failure list:\nattributes in body is required", "should valid error message about attributes")
	return " should throw 400 error"
}

func failScenarioNotTypeAccountSent(t *testing.T) string {
	// Given
	accountDataSunny, _ := amf.CreateDataToCreateAccount()
	accountDataSunny.Data.Type = ""
	data, _ := json.Marshal(accountDataSunny)
	// When
	_, err := src.Create(string(data))
	// Then
	require.Equal(t, err.Code, 400, "should throw 400")
	require.Equal(t, err.ServerMessage, "validation failure list:\nvalidation failure list:\ntype in body is required", "should valid error message about type")
	return " should throw 400 error"
}

func failScenarioNotIdSent(t *testing.T) string {
	// Given
	accountDataSunny, _ := amf.CreateDataToCreateAccount()
	accountDataSunny.Data.ID = ""
	data, _ := json.Marshal(accountDataSunny)
	// When
	_, err := src.Create(string(data))
	// Then
	require.Equal(t, err.Code, 400, "should throw 400")
	require.Equal(t, err.ServerMessage, "validation failure list:\nvalidation failure list:\nid in body is required", "should valid error message about id")
	return " should throw 400 error"
}

func failScenarioNotOrganizationIdSent(t *testing.T) string {
	// Given
	accountDataSunny, _ := amf.CreateDataToCreateAccount()
	accountDataSunny.Data.OrganisationID = ""
	data, _ := json.Marshal(accountDataSunny)
	// When
	_, err := src.Create(string(data))
	// Then
	require.Equal(t, err.Code, 400, "should throw 400")
	require.Equal(t, err.ServerMessage, "validation failure list:\nvalidation failure list:\norganisation_id in body is required", "should valid error message about organization")
	return " should throw 400 error"
}
