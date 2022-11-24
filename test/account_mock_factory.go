package test

import (
	m "clientForm3Api/src/models"
	"encoding/json"
)

const json_attributes = `{
	"account_classification": "",
	"account_matching_opt_out": "",
	"account_number": "",
	"alternative_names": "",
	"bank_id": "123456",
	"bank_id_code": "",
	"base_currency": "",
	"bic": "1",
	"country": "",
	"iban": "",
	"joint_account": "",
	"name": "",
	"secondary_identification": "",
	"status": "",
	"switched": "",
}`

const json_account = `{
	"attributes": ""
	"id": "1"
	"organisation_id": ""
	"type":""
	"version":""
}`

func CreateFakeAccountData() *m.AccountData {
	data := new(m.AccountData)
	json.Unmarshal([]byte(json_account), data)
	return data
}

func CreateFakeAccountAttributes() *m.AccountAttributes {
	data := new(m.AccountAttributes)
	json.Unmarshal([]byte(json_attributes), data)
	return data
}
