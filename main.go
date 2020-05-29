package main

import (
	"context"
	"fmt"
	"log"

	"github.com/akumar261089/curd-go-mongo/mongoConnect/handlers"
	"github.com/akumar261089/curd-go-mongo/mongoConnect/mongoConnect"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	mongoConnect.Connect()

	cursor, err := handlers.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(context.TODO(), &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
}
