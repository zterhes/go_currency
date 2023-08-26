package main

import (
	"go_currency/api"
	"go_currency/common"
	"go_currency/console"
)


func main() {
	if common.RunType == common.RUN_TYPE_CONSOLE{
		console.ConsoleConverter()
	}
	if common.RunType == common.RUN_TYPE_API{
		api.ApiConverter()
	}
}