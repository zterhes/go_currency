package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

var EnvMap map[string] string
var RunType string

const(
	APIKEY string = "API_KEY"
	RUN_TYPE ="RUN_TYPE"
	RUN_TYPE_CONSOLE = "console"
	RUN_TYPE_API = "api"
)

func init(){
	flag.StringVar(&RunType,"run","console","Type of running: console/api")
	flag.Parse()
	log.Println(RunType)
	if RunType == "" {
		log.Fatalln("Run type is not available")
	} else{
		log.Println("Starting as",RunType,"...")
	}
	var err error
	EnvMap, err = godotenv.Read(".env")
	if err != nil {
		log.Panicln(err)
	}
	InitClient()
	log.Println("Started")
}


func main() {
	if RunType == RUN_TYPE_CONSOLE{
		ConsoleCalculator()
	}
}