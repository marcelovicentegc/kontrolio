package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
)

type errorBody struct {
	Error string `json:"error"`
}

type ApiRecord struct {
	Time       string `json:"time"`
	RecordType string `json:"recordType"`
}

type createRecordResponseBody struct {
	Data ApiRecord `json:"data"`
}

type allRecordsResults struct {
	Results []ApiRecord `json:"results"`
}

type getAllRecordsResponseBody struct {
	Data allRecordsResults `json:"data"`
}

// CreateRecord creates an online record.
func CreateRecord(apiKey string) string {

	requestURL := config.GetBaseURL() + config.RecordEndpoint

	now := time.Now().In(time.Local).Format(time.RFC3339)

	jsonBytes := []byte(`{ "time": "` + now + `", "apiKey": "` + apiKey + `" }`)

	request, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBytes))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode > 400 {
		var responseBody errorBody

		err := json.NewDecoder(response.Body).Decode(&responseBody)

		if err != nil {
			log.Fatal(messages.FailedParsingRequest)
		}

		log.Fatal(responseBody.Error + "\n")
	}

	var responseBody createRecordResponseBody

	body, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Fatal(messages.FailedParsingRequest)
	}

	return responseBody.Data.RecordType
}

// GetAllRecords gets every record from the requesting user from the remote database.
func GetAllRecords(apiKey string) []ApiRecord {
	requestURL := config.GetBaseURL() + config.AllRecordsEndpoint

	jsonBytes := []byte(apiKey)

	request, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBytes))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode > 400 {
		var responseBody errorBody

		err := json.NewDecoder(response.Body).Decode(&responseBody)

		if err != nil {
			fmt.Println("Something went wrong while parsing the response body. [1]")
			log.Fatal(err.Error())
		}

		log.Fatal(responseBody.Error + "\n")
	}

	var responseBody getAllRecordsResponseBody

	body, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		fmt.Println("Something went wrong while parsing the response body. [2]")
		log.Fatal(err.Error())
	}

	return responseBody.Data.Results
}
