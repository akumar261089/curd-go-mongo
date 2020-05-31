package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akumar261089/curd-go-mongo/mongodb"
	"github.com/akumar261089/curd-go-mongo/restapi"
	"github.com/gorilla/handlers"
)

func main() {

	dbDetails := map[string][]string{
		"quickstart": {"jhfhgfhgfhg", "QuickstartDatabase", "spjhbjhgvjh"},
		"seconddb":   {"firstcoll", "second"},
	}
	mongodb.Connect(dbDetails)
	fmt.Println("DB connected")

	router := restapi.NewRouter()

	// Route Handlers / Endpoints

	// log.Fatal(http.ListenAndServe(":8080", router))
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))

}
