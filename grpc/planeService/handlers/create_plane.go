package handlers

import (
	"FlightBookingAPI/grpc/planeService/models"
	"FlightBookingAPI/pb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlaneServer) CreatePlane(ctx context.Context, req *pb.PlaneInfo) (*pb.PlaneInfo, error) {
	log.Printf("[Info]: Handle request CreatePlane(): %v\n", req)
	model := models.FromPBModel(req)

	err := s.repository.CreatePlane(model)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return req, nil
}
