package main

import (
	"context"
	"fmt"
	"log"

	"github.com/akumar261089/curd-go-mongo/handlers"
	"github.com/akumar261089/curd-go-mongo/mongoConnect"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	dbDetails := map[string][]string{
		"quickstart": {"jhfhgfhgfhg", "QuickstartDatabase", "spjhbjhgvjh"},
		"seconddb":   {"firstcoll", "second"},
	}

	mongoConnect.Connect(dbDetails)
	fmt.Println(handlers.Collection2)
	cursor, err := handlers.Collection2["quickstart"]["QuickstartDatabase"].Find(context.TODO(), bson.M{})
	//cursor, err := handlers.Collection[1].Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(context.TODO(), &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
}
