package common

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

var RunType string
var EnvMap map[string] string


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