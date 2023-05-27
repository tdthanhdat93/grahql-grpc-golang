package handlers

import (
	"FlightBookingAPI/grpc/flightService/repository"
	"FlightBookingAPI/grpc/flightService/requests"
	"FlightBookingAPI/pb"
	"sync"
)

type FlightServer struct {
	pb.UnimplementedFlightServiceServer
	repository repository.FlightRepository
	requester  requests.IPlaneRequests
	mutex      sync.Mutex
}

func NewFlightServer(flightRepository repository.FlightRepository, planeRequester requests.IPlaneRequests) (*FlightServer, error) {
	return &FlightServer{
		repository: flightRepository,
		requester:  planeRequester,
		mutex:      sync.Mutex{}}, nil
}
