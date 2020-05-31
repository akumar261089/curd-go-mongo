package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/akumar261089/curd-go-mongo/mongodb"
	"gopkg.in/mgo.v2/bson"
)

func WriteDB(w http.ResponseWriter, r *http.Request) {
	// podcastsCollection := Connect("quickstart", "QuickstartDatabase")

	// m := map[string]map[string][]string{
	// 	"stocks": {
	// 		"prediction_type":  {"trend_prediction_lstm", "sentiments_nlp"},
	// 		"available_values": {"Yahoo", "Google"},
	// 	},
	// 	"recipe": {
	// 		"prediction_type":  {"ingredents_list_texgenrnn", "process_texgenrnn"},
	// 		"available_values": {"idli", "pasta"},
	// 	},
	// }
	// jsonbyte, _ := json.Marshal(m)
	// jsonString := string(jsonbyte)
	// var bdoc interface{}
	// bson.UnmarshalJSON([]byte(jsonString), &bdoc)
	// fmt.Println(bdoc)
	// fmt.Println(podcastsCollection)
	// podcastResult, err := podcastsCollection.InsertOne(context.TODO(), (&bdoc))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(podcastResult)
	// // fmt.Fprint(w, podcastResult)

	// fmt.Println(podcastResult)

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetSavedModels(w http.ResponseWriter, r *http.Request) {

	cursor, err := mongodb.Collection2["quickstart"]["QuickstartDatabase"].Find(context.TODO(), bson.M{})
	//cursor, err := handlers.Collection[1].Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(context.TODO(), &episodes); err != nil {
		log.Fatal(err)
	}
	// jsonbyte, _ := json.Marshal(m)
	// jsonString := string(jsonbyte)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(episodes)
	// fmt.Fprint(w, episodes)

	// if err := json.NewEncoder(w).Encode(jsonString); err != nil {
	// 	panic(err)
	// }
}
func FittedModelDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func RawModelDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
