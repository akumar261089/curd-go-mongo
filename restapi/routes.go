package restapi

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Welcome",
		"GET",
		"/welcome",
		WriteDB,
	},
	Route{
		"GetSavedModels",
		"GET",
		"/getsavedmodels",
		GetSavedModels,
	},
	Route{
		"FittedModelDetails",
		"GET",
		"/fittedmodeldetails",
		FittedModelDetails,
	},
	Route{
		"RawModelDetails",
		"GET",
		"/rawmodeldetails",
		RawModelDetails,
	},
}
