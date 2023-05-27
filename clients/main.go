package main

import (
	"FlightBookingAPI/clients/graphql/generated"
	"FlightBookingAPI/clients/graphql/resolver"
	"FlightBookingAPI/clients/services"
	"FlightBookingAPI/config"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

var (
	configPath = flag.String("config", "./config.json", "Location of config.json")
)

func main() {
	conf := config.LoadConfig(*configPath)
	planeService := services.NewPlaneAPIService(conf)
	flightService := services.NewFlightAPIService(conf)
	// resolver.NewResolver()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(planeService, flightService)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	addr := fmt.Sprintf("%v:%v", conf.APIConf.Host, conf.APIConf.Port)
	log.Printf("connect to http://%s/ for GraphQL playground", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
