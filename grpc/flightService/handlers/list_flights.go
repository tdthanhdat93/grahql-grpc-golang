package handlers

import (
	"FlightBookingAPI/pb"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *FlightServer) ListFlights(_ *emptypb.Empty, stream pb.FlightService_ListFlightsServer) error {
	log.Printf("[Info]: Handle request ListFlights()\n")

	flights, err := s.repository.ListFlights()
	if err != nil {
		log.Printf("[Error]: ListFlights() query db failed: %v", err)
		return status.Error(codes.Internal, err.Error())
	}

	for _, flight := range flights {
		err := stream.Send(flight.ToPBModel())
		if err != nil {
			log.Printf("[Error]: ListFlights() send data (%v) failed: %v\n", flight, err)
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
