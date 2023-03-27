package main

import (
	"log"
	"mux-mongo-api/configs"
	"mux-mongo-api/routes"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "mux-mongo-api/docs"

	"github.com/gorilla/mux"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server User server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000/
func main() {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:7000/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this

	log.Fatal(http.ListenAndServe(":7000", router))
}
