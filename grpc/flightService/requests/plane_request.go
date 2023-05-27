package requests

import (
	"FlightBookingAPI/config"
	"FlightBookingAPI/pb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IPlaneRequests interface {
	FindPlane(id string) *pb.PlaneInfo
}

type PlaneRequester struct {
	client pb.PlaneServiceClient
}

func NewPlaneRequester(serviceConfig *config.Config) IPlaneRequests {
	return &PlaneRequester{
		client: connectService(serviceConfig),
	}
}
func connectService(serviceConfig *config.Config) pb.PlaneServiceClient {
	serverAddr := fmt.Sprintf("%v:%v", serviceConfig.GRPCConf.PlaneGRPCConf.Host, serviceConfig.GRPCConf.PlaneGRPCConf.Port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[Error]: PlaneRequests failed to dial (%v): %v\n", serverAddr, err)
	}
	return pb.NewPlaneServiceClient(conn)
}

func (p *PlaneRequester) FindPlane(id string) *pb.PlaneInfo {
	res, err := p.client.FindPlane(context.Background(), &pb.FindPlaneByIdRequest{Id: id})
	if err != nil {
		log.Printf("[Error]: FindPlane(%v) request failed: %v", id, err)
		return nil
	}

	return res
}
