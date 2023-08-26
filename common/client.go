package common

import (
	"encoding/json"
	"net/http"
	"time"
)


var client *http.Client


func InitClient(){
	client = &http.Client{Timeout: 10 * time.Second}
}

func createUrl(base string, to string, apiKey string) string {
	requestUrl := "http://data.fixer.io/api/latest?access_key="+ apiKey +"&base="+ base +"&symbols="+to
	return requestUrl
}

func callClient(url string, target interface{}) error {
	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}