package handlers

import (
	"FlightBookingAPI/grpc/flightService/models"
	"FlightBookingAPI/pb"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FlightServer) FindFlightsByPlaneId(req *pb.FindFlightByPlaneIdRequest, stream pb.FlightService_FindFlightsByPlaneIdServer) error {
	log.Printf("[Info]: Handle request FindFlightsByPlaneId()\n")

	flights, err := s.repository.FindFlightsByPlaneId(req.PlaneId)
	if err != nil {
		log.Printf("[Error]: FindFlightsByPlaneId(%v) query db failed: %v", req.PlaneId, err)
		return status.Error(codes.Internal, err.Error())
	}

	for _, flight := range flights {
		err := stream.Send(flight.ToPBModel())
		if err != nil {
			log.Printf("[Error]: FindFlightsByPlaneId() send data (%v) failed: %v\n", flight, err)
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (s *FlightServer) FindFlights(req *pb.FindFlightRequest, stream pb.FlightService_FindFlightsServer) error {
	log.Printf("[Info]: Handle request FindFlights()\n")

	flights, err := s.repository.FindFlights(&models.FindFlightParams{
		Departure:      req.Departure,
		Arrival:        req.Arrival,
		StartTimeRange: models.FromPBDateTime(req.StartTimeRange),
		EndTimeRange:   models.FromPBDateTime(req.EndTimeRange),
	})
	if err != nil {
		log.Printf("[Error]: FindFlights(%v) query db failed: %v", req, err)
		return status.Error(codes.Internal, err.Error())
	}

	for _, flight := range flights {
		err := stream.Send(flight.ToPBModel())
		if err != nil {
			log.Printf("[Error]: FindFlights() send data (%v) failed: %v\n", flight, err)
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
