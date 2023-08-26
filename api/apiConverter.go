package api

import (
	"fmt"
	"go_currency/common"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var fooHandler http.Handler

const(
	INPUT string = "input"
	OUTPUT		 = "output"
	AMOUNT 		 = "amount"		
)


func ApiConverter(){
	fmt.Println(fooHandler)

	http.HandleFunc("/",requestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request){
	converterRequest := common.ConverterRequest{}
	varMap := transformRequestStringToVarMap(r.RequestURI)
	transformVarMapToConverterRequest(varMap,&converterRequest)
	result := common.Convert(converterRequest)
	fmt.Fprintf(w, "%v%v, is %v%v",converterRequest.Amount,converterRequest.InputCurrency,result,converterRequest.OutputCurrency )
}

func transformRequestStringToVarMap(requestURI string) map[string]string{
	var requestVarMap = map[string]string{}
	varList := strings.Split(requestURI, "?")[1:][0]
	varPairs := strings.Split(varList, "&")
	for _, pair := range varPairs{
		pairSlice := strings.Split(pair, "=")
		requestVarMap[pairSlice[0]]=pairSlice[1]
	}
	return requestVarMap
}

func transformVarMapToConverterRequest(varMap map[string]string, request *common.ConverterRequest){
	amount,err := strconv.ParseFloat(varMap[AMOUNT], 64)
	if err != nil {
		log.Fatalln(err)
	}
	request.Amount = amount
	request.InputCurrency = varMap[INPUT]
	request.OutputCurrency = varMap[OUTPUT]
}

