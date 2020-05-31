package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var Collection []*mongo.Collection
var Collection2 = map[string]map[string]*mongo.Collection{}

func GetCollection(databaseexport []DbAddrCollectionList) {

	for _, db := range databaseexport {
		fmt.Println(db.DbName)
		Collection2[db.DbName] = map[string]*mongo.Collection{}

		for _, collectionName := range db.CollectionList {
			//Collection = append(Collection, c.Collection(collectionName))
			//colAdd := c.Collection(collectionName)
			//fmt.Println(colAdd)
			colmap := map[string]*mongo.Collection{
				collectionName: db.DbAddr.Collection(collectionName),
			}
			//fmt.Println(colmap)
			for key, val := range colmap {
				Collection2[db.DbName][key] = val
			}
		}
	}

	//fmt.Println(Collection[1])
	//Collection2["test"]["test"] = Collection[1]

	//fmt.Println(Collection2)
}

// func GetCollection(c *mongo.Database, col []string) {

// 	for _, collectionName := range col {
// 		Collection = append(Collection, c.Collection(collectionName))
// 	}
// }
