package models

import "go.mongodb.org/mongo-driver/mongo"

type DbAddrCollectionList struct {
	DbName         string
	DbAddr         *mongo.Database
	CollectionList []string
}
