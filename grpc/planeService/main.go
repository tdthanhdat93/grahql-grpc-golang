package main

import (
	"FlightBookingAPI/config"
	"FlightBookingAPI/grpc/planeService/handlers"
	"FlightBookingAPI/grpc/planeService/repository"
	"FlightBookingAPI/grpc/planeService/requests"
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
	port := conf.GRPCConf.PlaneGRPCConf.Port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("[Error]: Failed to listen (localhost:%v): %v", port, err)
	}

	log.Printf("[Info]: Start server (localhost:%v). Listening...", port)
	grpcServer := grpc.NewServer()
	planeRepository, err := repository.NewDBManager()
	if err != nil {
		log.Fatalf("[Error]: Create repository failed: %v\n", err)
	}
	flightRequester := requests.NewFlightRequester(conf) // Communicate with flightService to check flight info
	planeServer, err := handlers.NewPlaneServer(planeRepository, flightRequester)
	if err != nil {
		log.Fatalf("[Error]: Create handler failed: %v\n", err)
	}
	pb.RegisterPlaneServiceServer(grpcServer, planeServer)
	grpcServer.Serve(lis)
}

func main() {
	flag.Parse()
	startServer()
}
