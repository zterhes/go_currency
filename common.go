package main

import (
	"log"
	"strings"
)

func Calculate(data FixerResponse, converterRequest *ConverterRequest) float64 {
	upperTo := strings.ToUpper(converterRequest.outputCurrency)
	return converterRequest.amount*float64(data.Rates[upperTo])
}

func GetCurrencies(converterRequest *ConverterRequest, fixerResponse *FixerResponse) error {
	requestUrl := CreateUrl(converterRequest.inputCurrency, converterRequest.outputCurrency, EnvMap[APIKEY])
	log.Println("Url:",requestUrl)
	return CallClient(requestUrl, fixerResponse)
}

func errorHandler(err error){
	if err != nil {
		log.Fatalln(err)
	}
}