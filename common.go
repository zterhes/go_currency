package main

import (
	"encoding/json"
	"strings"
)

func CreateUrl(base string, to string, apiKey string) string {
	requestUrl := "http://data.fixer.io/api/latest?access_key="+ apiKey +"&base="+ base +"&symbols="+to
	return requestUrl
}

func GetCurrencies(url string, target interface{}) error {
	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}

func Calculate(data FixerResponse, to string, amount float64) float64 {
	upperTo := strings.ToUpper(to)
	return amount*float64(data.Rates[upperTo])
}