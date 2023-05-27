package handlers

import (
	"FlightBookingAPI/pb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FlightServer) BookFlight(ctx context.Context, req *pb.BookingRequest) (*pb.FlightInfo, error) {
	log.Printf("[Info]: Handle request BookFlight(): %v\n", req)
	flight, err := s.repository.GetFlight(req.FlightId)
	if err != nil {
		return nil, err
	}
	if flight.AvailableSeat == 0 {
		return nil, status.Error(codes.Internal, "Flight out of seat")
	}
	s.mutex.Lock() // Avoid data race
	defer s.mutex.Unlock()
	flight.AvailableSeat--
	updatedFlight, err := s.repository.UpdateFlightAvailableSeat(flight)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return updatedFlight.ToPBModel(), nil
}
