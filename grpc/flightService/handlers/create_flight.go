package handlers

import (
	"FlightBookingAPI/grpc/flightService/models"
	"FlightBookingAPI/pb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FlightServer) CreateFlight(ctx context.Context, req *pb.FlightInfo) (*pb.FlightInfo, error) {
	log.Printf("[Info]: Handle request CreateFlight(): %v\n", req)

	plane := s.requester.FindPlane(req.PlaneId) // Call plane service to check plane info
	if plane == nil {
		return nil, status.Error(codes.Internal, "Plane not found")
	}

	if plane.Status != pb.PlaneStatus_STS_READY {
		return nil, status.Error(codes.Internal, "Plane not ready")
	}
	if plane.TotalSeat < req.AvailableSeat {
		return nil, status.Error(codes.Internal, "Plane not enough seat")
	}

	model := models.FromPBModel(req)
	log.Printf("[Info]: CreateFlight write record to db: %v\n", model)
	err := s.repository.CreateFlight(model)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return req, nil
}
