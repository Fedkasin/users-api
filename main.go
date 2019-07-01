package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	//"time"

	"github.com/Fedkasin/users-api/config"
	"github.com/Fedkasin/users-api/mongodb"
	"github.com/Fedkasin/users-api/router"
)

var conf = config.Config

func init() {
	
	ENV := os.Getenv("ENVIRONMENT")
	if len(ENV) == 0 {
		ENV = "local"
	}

	file, err := os.Open("./config/config." + ENV +".json")
	if err != nil {
		log.Println("Cannot find %s config file!", ENV)
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config.Config)
	if err != nil {
		log.Println("Invalid %s config file!", ENV)
		return
	}
	if (ENV != "local") {
		config.Config.DBHostname = os.Getenv("MONGO_HOST")
		config.Config.DbName = os.Getenv("MONGO_DB")
		config.Config.DBPassword = os.Getenv("MONGO_PASSWORD")
		config.Config.DBPort, err = strconv.Atoi(os.Getenv(("MONGO_PORT")))
		config.Config.DBUser = os.Getenv("MONGO_USER")
	}
}

func main() {
	mongodb.Initialize()
	defer mongodb.Close()

	log.Println("App listening on port " + strconv.Itoa(conf.Port) + "...")

	if err := http.ListenAndServe(":" + strconv.Itoa(conf.Port), router.Router); err != nil {
		panic(err)
	}
}


