package mongoConnect

import (
	"context"
	"log"
	"time"

	"github.com/akumar261089/curd-go-mongo/mongoConnect/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoCollection *mongo.Collection
var Database *mongo.Database

func Connect() {

	// Database Config
	MongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	//Set up a context required by mongo.Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = MongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Database := MongoClient.Database("quickstart")
	handlers.GetCollection(Database)
	return
}
