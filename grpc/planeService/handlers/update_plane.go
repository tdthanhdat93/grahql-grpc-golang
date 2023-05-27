package handlers

import (
	"FlightBookingAPI/grpc/planeService/models"
	"FlightBookingAPI/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlaneServer) UpdatePlane(ctx context.Context, req *pb.PlaneInfo) (*pb.PlaneInfo, error) {
	model := models.FromPBModel(req)
	plane, err := s.repository.UpdatePlane(model)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return plane.ToPBModel(), nil
}
