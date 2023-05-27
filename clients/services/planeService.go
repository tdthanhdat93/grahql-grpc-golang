package services

import (
	models "FlightBookingAPI/clients/services/models"
	"FlightBookingAPI/config"
	"FlightBookingAPI/pb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IPlaneAPIService interface {
	CreatePlane(plane *models.Plane) (*models.Plane, error)
	UpdatePlane(plane *models.Plane) (*models.Plane, error)
	ListPlanes() ([]*models.Plane, error)
	DeletePlane(id string) error
}

type PlaneAPIService struct {
	client pb.PlaneServiceClient
}

func NewPlaneAPIService(serviceConfig *config.Config) IPlaneAPIService {
	return &PlaneAPIService{
		client: connectService(serviceConfig),
	}
}
func connectService(serviceConfig *config.Config) pb.PlaneServiceClient {
	serverAddr := fmt.Sprintf("%v:%v", serviceConfig.GRPCConf.PlaneGRPCConf.Host, serviceConfig.GRPCConf.PlaneGRPCConf.Port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[Error]: PlaneAPIService failed to dial (%v): %v\n", serverAddr, err)
	}
	return pb.NewPlaneServiceClient(conn)
}

func (api *PlaneAPIService) CreatePlane(plane *models.Plane) (*models.Plane, error) {
	res, err := api.client.CreatePlane(context.Background(), plane.ToPBModel())
	if err != nil {
		return plane, err
	}
	return models.FromPBPlaneModel(res), nil
}

func (api *PlaneAPIService) UpdatePlane(plane *models.Plane) (*models.Plane, error) {
	res, err := api.client.UpdatePlane(context.Background(), plane.ToPBModel())
	if err != nil {
		return plane, err
	}
	return models.FromPBPlaneModel(res), nil
}

func (api *PlaneAPIService) ListPlanes() ([]*models.Plane, error) {
	stream, err := api.client.ListPlanes(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	planes := []*models.Plane{}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return planes, err
		}
		planes = append(planes, models.FromPBPlaneModel(res))
	}
	return planes, nil
}

func (api *PlaneAPIService) DeletePlane(id string) error {
	_, err := api.client.DeletePlane(context.Background(), &pb.DeletePlaneByIdRequest{Id: id})
	return err
}
