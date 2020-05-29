package handlers

import "go.mongodb.org/mongo-driver/mongo"

var Collection *mongo.Collection

func GetCollection(c *mongo.Database) {
	Collection = c.Collection("QuickstartDatabase")
}
