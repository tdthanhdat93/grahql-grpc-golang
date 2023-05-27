package handlers

import (
	"FlightBookingAPI/grpc/flightService/models"
	"FlightBookingAPI/pb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FlightServer) UpdateFlight(ctx context.Context, req *pb.FlightInfo) (*pb.FlightInfo, error) {
	log.Printf("[Info]: Handle request UpdateFlight(): %v\n", req)
	model := models.FromPBModel(req)
	flight, err := s.repository.UpdateFlight(model)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return flight.ToPBModel(), nil
}
