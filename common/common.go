package common

import (
	"log"
	"strings"
)

func calculate(data FixerResponse, converterRequest *ConverterRequest) float64 {
	upperTo := strings.ToUpper(converterRequest.OutputCurrency)
	return converterRequest.Amount*float64(data.Rates[upperTo])
}

func getCurrencies(converterRequest *ConverterRequest, fixerResponse *FixerResponse ) error {
	requestUrl := createUrl(converterRequest.InputCurrency, converterRequest.OutputCurrency,  EnvMap[APIKEY])
	return callClient(requestUrl, fixerResponse)
}

func ErrorHandler(err error){
	if err != nil {
		log.Fatalln(err)
	}
}