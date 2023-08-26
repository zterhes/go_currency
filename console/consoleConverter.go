package console

import (
	"bufio"
	"fmt"
	"go_currency/common"
	"log"
	"os"
	"strconv"
)

func ConsoleConverter(){
	var err error
	var converterRequest common.ConverterRequest
	converterRequest.Amount, err = getInputFloat("amount")
	common.ErrorHandler(err)
	converterRequest.InputCurrency,err = getInputString("base")
	converterRequest.OutputCurrency,err = getInputString("base")
	result := common.Convert(converterRequest)
	fmt.Println(converterRequest.Amount,converterRequest.InputCurrency,"is",result,converterRequest.OutputCurrency)
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
