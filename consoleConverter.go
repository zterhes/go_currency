package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var converterRequest ConverterRequest

func ConsoleConverter(){
	var err error
	converterRequest.amount, err = getInputFloat("amount")
	errorHandler(err)
	converterRequest.inputCurrency,err = getInputString("base")
	converterRequest.outputCurrency,err = getInputString("base")
	log.Println(converterRequest)
	var fixerResponse FixerResponse
	err = GetCurrencies(&converterRequest, &fixerResponse)
	errorHandler(err)
	if fixerResponse.Success == false{
		log.Panicln("Request is not succeded")
	}
	result := Calculate(fixerResponse, &converterRequest)
	fmt.Println(converterRequest.amount,converterRequest.inputCurrency,"is",result,converterRequest.outputCurrency)
}



func getInputString(title string) (string, error){
	fmt.Printf("Please add %s :", title)
	input := scanner()
	return input, nil
}



func getInputFloat(title string) (float64, error){
	fmt.Printf("Please add %s :", title)
	input := scanner()
	log.Println(input)
	return strconv.ParseFloat(input, 64)
}



func scanner() string{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
