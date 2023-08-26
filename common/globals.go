package common

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var RunType string
var EnvMap map[string] string


func init(){
	flag.StringVar(&RunType,"run","","Type of running: console/api")
	flag.Parse()
	fmt.Println(RunType)
	var err error
	EnvMap, err = godotenv.Read(".env")
	fmt.Println(EnvMap)
	ErrorHandler(err)
	if RunType == "" {
		RunType = EnvMap[RUN_TYPE]
		if RunType == "" {
			log.Fatalln("Run type is not available")
		}
	} else{
		log.Println("Starting as",RunType,"...")
	}
	InitClient()
	log.Println("Started")
}