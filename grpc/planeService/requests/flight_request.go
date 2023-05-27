package requests

import (
	"FlightBookingAPI/config"
	"FlightBookingAPI/pb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IFlightRequests interface {
	FindFlightsByPlaneId(id string) ([]*pb.FlightInfo, error)
}

type FlightRequester struct {
	client pb.FlightServiceClient
}

func NewFlightRequester(serviceConfig *config.Config) IFlightRequests {
	return &FlightRequester{
		client: connectService(serviceConfig),
	}
}

func connectService(serviceConfig *config.Config) pb.FlightServiceClient {
	serverAddr := fmt.Sprintf("%v:%v", serviceConfig.GRPCConf.FlightGRPCConf.Host, serviceConfig.GRPCConf.FlightGRPCConf.Port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[Error]: FlightRequests failed to dial (%v): %v\n", serverAddr, err)
	}
	return pb.NewFlightServiceClient(conn)
}

func (p *FlightRequester) FindFlightsByPlaneId(id string) ([]*pb.FlightInfo, error) {
	stream, err := p.client.FindFlightsByPlaneId(context.Background(), &pb.FindFlightByPlaneIdRequest{PlaneId: id})
	if err != nil {
		return nil, err
	}

	flights := []*pb.FlightInfo{}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return flights, err
		}
		flights = append(flights, res)
	}
	return flights, nil
}
