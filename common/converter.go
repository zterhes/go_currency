package common

import "log"


func Convert(converterRequest ConverterRequest) float64{
	var fixerResponse FixerResponse
	err := getCurrencies(&converterRequest, &fixerResponse)
	ErrorHandler(err)
	if fixerResponse.Success == false{
		log.Panicln("Request is not succeded")
	}
	return calculate(fixerResponse, &converterRequest)
}