package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Fedkasin/users-api/config"
	"github.com/Fedkasin/users-api/mongodb"
	"github.com/Fedkasin/users-api/router"
)

var conf = config.Config

func init() {
	env := os.Getenv("ENVIRONMENT")

	if len(env) == 0 {
		env = "local"
	}

	file, err := os.Open("./config/config." + env +".json")
	if err != nil {
		log.Println("Cannot find " + env + " config file!")
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config.Config)
	if err != nil {
		log.Println("Invalid %s config file!", env)
		return
	}
	log.Println(config.Config)
}

func main() {
	mongodb.Initialize()
	defer mongodb.Close()

	log.Println("App listening on port " + strconv.Itoa(conf.Port) + "...")

	if err := http.ListenAndServe(":" + strconv.Itoa(conf.Port), router.Router); err != nil {
		panic(err)
	}
}


