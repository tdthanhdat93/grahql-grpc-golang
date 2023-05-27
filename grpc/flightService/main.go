package main

import (
	"FlightBookingAPI/config"
	"FlightBookingAPI/grpc/flightService/handlers"
	"FlightBookingAPI/grpc/flightService/repository"
	"FlightBookingAPI/grpc/flightService/requests"
	"FlightBookingAPI/pb"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	configPath = flag.String("config", "./config.json", "Location of config.json")
)

func startServer() {
	conf := config.LoadConfig(*configPath)
	port := conf.GRPCConf.FlightGRPCConf.Port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("[Error]: Failed to listen (localhost:%v): %v", port, err)
	}

	log.Printf("[Info]: Start server (localhost:%v). Listening...", port)
	grpcServer := grpc.NewServer()
	flightRepository, err := repository.NewDBManager()
	if err != nil {
		log.Fatalf("[Error]: Create repository failed: %v\n", err)
	}
	planeRequester := requests.NewPlaneRequester(conf) // Communicate with planeSerive to check plane info
	flightServer, err := handlers.NewFlightServer(flightRepository, planeRequester)
	if err != nil {
		log.Fatalf("[Error]: Create handler failed: %v\n", err)
	}
	pb.RegisterFlightServiceServer(grpcServer, flightServer)
	grpcServer.Serve(lis)
}

func main() {
	flag.Parse()
	startServer()
}
