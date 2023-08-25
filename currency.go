package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var inputCurrency, outputCurrency string
var amount float64
var client *http.Client
var EnvMap map[string] string

const(
	APIKEY string = "API_KEY"
)


type FixerResponse struct {
	Success bool
	Timestamp int
	Base string
	Rates map[string]float32
}

func init(){
	flag.StringVar(&inputCurrency,"icurr","eur","input currency")
	flag.StringVar(&outputCurrency,"ocurr","huf","output currency")
	flag.Float64Var(&amount,"amount",0.,"amount")
	flag.Parse()
	var err error
	EnvMap, err = godotenv.Read(".env")
	if err != nil {
		log.Panicln(err)
	}
}


func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	requestUrl := CreateUrl(inputCurrency,outputCurrency,EnvMap[APIKEY])
	var fixerResponse FixerResponse
	err := GetCurrencies(requestUrl,&fixerResponse)
	if err != nil {
		log.Panicln(err)
	}
	if fixerResponse.Success == false{
		log.Panicln("Request is not succeded")
	}
	result := Calculate(fixerResponse, outputCurrency, amount)
	fmt.Println(amount,inputCurrency,"is",result,outputCurrency)
	
}