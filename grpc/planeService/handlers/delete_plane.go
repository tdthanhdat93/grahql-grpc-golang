package handlers

import (
	"FlightBookingAPI/pb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PlaneServer) DeletePlane(ctx context.Context, req *pb.DeletePlaneByIdRequest) (*emptypb.Empty, error) {
	log.Printf("[Info]: Handle request DeletePlane(): %v\n", req)

	flights, err := s.requester.FindFlightsByPlaneId(req.Id) // call flightService to check schedule of delete plane
	if err != nil {
		return nil, status.Error(codes.Internal, "Can not check schedule of plane")
	}
	if len(flights) > 0 {
		return nil, status.Error(codes.Internal, "Plane has scheduled for flight")
	}

	err = s.repository.DeletePlane(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
