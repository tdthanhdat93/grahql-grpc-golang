package handlers

import (
	"FlightBookingAPI/pb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlaneServer) FindPlane(ctx context.Context, req *pb.FindPlaneByIdRequest) (*pb.PlaneInfo, error) {
	log.Printf("[Info]: Handle request FindPlane(): %v\n", req)

	plane, err := s.repository.FindPlane(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return plane.ToPBModel(), nil
}
