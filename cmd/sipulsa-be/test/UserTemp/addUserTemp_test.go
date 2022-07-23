package usertemp_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"sipulsa-be/models"
	"sipulsa-be/models/responses"
	"testing"
)

func TestAddUserTemp(t *testing.T) {
	var baseUrl = "http://localhost:8089/user"

	var body = models.UserTemp{
		Name:                     "Mardiana",
		Username:                 "medr0it",
		PhoneNumber:              "082125573465",
		Email:                    "madaresdhee@gmail.com",
		RegisteredByReferralCode: sql.NullString{Valid: false},
	}

	jsonBody, errJsonBody := json.Marshal(body)

	if errJsonBody != nil {
		panic(errJsonBody)
	}

	var payload = bytes.NewBuffer(jsonBody)

	res, err := http.Post(baseUrl, "application/json", payload)
	t.Fail()

	if err != nil {
		t.Fail()
		t.Log(err.Error())
	}
	var resp responses.Response

	decode := json.NewDecoder(res.Body)
	decode.Decode(&resp)

	t.Log(resp)
}
