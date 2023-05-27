package handlers

import (
	"FlightBookingAPI/grpc/planeService/repository"
	"FlightBookingAPI/grpc/planeService/requests"
	"FlightBookingAPI/pb"
)

type PlaneServer struct {
	pb.UnimplementedPlaneServiceServer
	repository repository.PlaneRepository
	requester  requests.IFlightRequests
}

func NewPlaneServer(planeRepository repository.PlaneRepository, flightRequester requests.IFlightRequests) (*PlaneServer, error) {
	return &PlaneServer{
		repository: planeRepository,
		requester:  flightRequester}, nil
}
