package handlers

import (
	"FlightBookingAPI/pb"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *FlightServer) DeleteFlight(ctx context.Context, req *pb.DeleteFlightByIdRequest) (*emptypb.Empty, error) {
	log.Printf("[Info]: Handle request DeleteFlight(): %v\n", req)

	err := s.repository.DeleteFlight(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
