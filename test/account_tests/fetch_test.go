package account_tests

import (
	"clientForm3Api/src"
	"clientForm3Api/test"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFetchAccount(t *testing.T) {
	fmt.Println("Starting Test of Account - Fetch ")
	msg := sunnyGetByIdTest(t)
	msgs := []string{"when want to Fetch an Account by Id " + msg}

	msg = notFoundAccountCheck(t)
	msgs = append(msgs, "and it doesn't exist "+msg)

	test.PrintTestsMessages(msgs)
}

func notFoundAccountCheck(t *testing.T) string {
	id := "a227e265-9605-4b4b-a0e5-3003ea9cc11"
	_, err := src.Fetch(id)
	require.Equal(t, err.Code, 400, "the account was not found ")
	msg := " should throw 400 error"
	return msg
}

func sunnyGetByIdTest(t *testing.T) string {
	msg := "and id Exists"
	id := "a227e265-9605-4b4b-a0e5-3003ea9cc4dc"
	aAccountFound, _ := src.Fetch(id)
	require.Equal(t, aAccountFound.ID, id, "the account created not contain the same ID ")
	msg += " should be found"
	return msg
}
