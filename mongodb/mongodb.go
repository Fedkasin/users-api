package mongodb

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Fedkasin/users-api/config"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client *mongo.Client

type InsertOneResult *mongo.InsertOneResult
type UpdateResult *mongo.UpdateResult
type SingleResult *mongo.SingleResult

var conf = config.Config

func Initialize() {
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	Client, err = mongo.Connect(ctx, "mongodb://" + conf.DBHostname + ":" + strconv.Itoa(conf.DBPort))
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongodb on " + conf.DBHostname + ":" + strconv.Itoa(conf.DBPort))
}

func Close() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := Client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}



