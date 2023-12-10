package sqlcore

import (
	"bytes"
	"encoding/json"
	"github.com/mrveera/sxt-sdks/helpers"
	"io/ioutil"
	"net/http"
	"os"
)

// Run all DML queries
func DML(sqlText, originApp string, biscuitArray []string, resources []string, accesstoken string) (errMsg string, status bool) {
	apiEndPoint, _ := helpers.ReadEndPointGeneral()
	tokenEndPoint := apiEndPoint + "/sql/dml"

	at := accesstoken

	if accesstoken == "" {
		at = os.Getenv("accessToken")
	}

	client := http.Client{}
	postBody, _ := json.Marshal(map[string]interface{}{
		"biscuits":  biscuitArray,
		"resources": resources,
		"sqlText":   sqlText,
	})

	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", tokenEndPoint, responseBody)
	if err != nil {
		return err.Error(), false
	}

	req.Header.Add("Authorization", "Bearer "+at)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("originApp", originApp)

	res, err := client.Do(req)
	if err != nil {
		return err.Error(), false
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error(), false
	}

	if res.StatusCode != 200 {
		return string(body), false
	}

	return "", true
}
