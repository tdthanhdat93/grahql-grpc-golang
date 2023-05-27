package handlers

import (
	"FlightBookingAPI/pb"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PlaneServer) ListPlanes(_ *emptypb.Empty, stream pb.PlaneService_ListPlanesServer) error {
	log.Printf("[Info]: Handle request ListPlanes()\n")

	planes, err := s.repository.ListPlanes()
	if err != nil {
		log.Printf("[Error]: ListPlanes() query db failed: %v", err)
		return status.Error(codes.Internal, err.Error())
	}

	for _, plane := range planes {
		err := stream.Send(plane.ToPBModel())
		if err != nil {
			log.Printf("[Error]: ListPlanes() send data (%v) failed: %v\n", plane, err)
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
