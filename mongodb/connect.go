package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Connect(dbDetails map[string][]string) {

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
	// var databasedetails map[string]map[string][]
	databaseexport := make([]DbAddrCollectionList, len(dbDetails))
	i := 0
	for dbname, cols := range dbDetails {
		// fmt.Println(dbname, cols)
		databaseexport[i].DbName = dbname
		// fmt.Println(databaseexport[i].DbName)
		databaseexport[i].DbAddr = MongoClient.Database(dbname)
		// fmt.Println(databaseexport[i].DbAddr)
		databaseexport[i].CollectionList = cols
		// fmt.Println(databaseexport[i].CollectionList)
		// fmt.Println(databaseexport[i])

		i++
	}

	GetCollection(databaseexport)
	// Database := MongoClient.Database(dbname)
	// var cols []string
	// cols = append(cols, "QuickstartDatabase")
	// handlers.GetCollection(Database, cols)
	return
}
