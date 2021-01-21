package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	BASE_URL         = "https://kontrolio.com/api/"
	RECORD_ENDPOINT  = "record"
	RECORDS_ENDPOINT = "records"
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

func CreateRecord(apiKey string) string {

	requestUrl := BASE_URL + RECORD_ENDPOINT

	now := time.Now().In(time.Local).Format(time.RFC3339)

	jsonBytes := []byte(`{ "time": "` + now + `", "apiKey": "` + apiKey + `" }`)

	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonBytes))

	// req.Header.Set("X-Custom-Header", "myvalue")
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
			log.Fatal("Something went wrong while parsing the response body.\n")
		}

		log.Fatal(responseBody.Error + "\n")
	}

	var responseBody createRecordResponseBody

	body, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Fatal("Something went wrong while parsing the response body.\n")
	}

	return responseBody.Data.RecordType
}
