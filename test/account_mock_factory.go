package test

import (
	"clientForm3Api/src/models"
	"encoding/json"
	"github.com/google/uuid"
)

const sunnyAccountJsonData = `{
  "data": {
    "type": "accounts",
    "id": "a227e265-9605-4b4b-a0e5-3003ea9cc4dc",
    "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
    "attributes": {
      "country": "GB",
      "base_currency": "GBP",
      "bank_id": "400300",
      "bank_id_code": "GBDSC",
      "bic": "NWBKGB22",
      "name": [
        "David 02"
      ],
      "alternative_names": [
        "David 02"
      ],
      "account_classification": "Personal",
      "joint_account": false,
      "name_matching_status": "supported",
      "secondary_identification": "A1B2C3D4"
    }
  }
}`

func CreateMockAccountDataFromJsonString() models.AccountFetchData {
	var accData models.AccountFetchData
	json.Unmarshal([]byte(sunnyAccountJsonData), &accData)
	return accData
}

func CreateDataToCreateAccount() (models.AccountFetchData, string) {
	id := uuid.New()
	accountDataSunny := CreateMockAccountDataFromJsonString()
	accountDataSunny.Data.ID = id.String()
	return accountDataSunny, id.String()
}
