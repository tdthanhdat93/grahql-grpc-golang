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

type IFlightAPIService interface {
	CreateFlight(flight *models.Flight) (*models.Flight, error)
	UpdateFlight(flight *models.Flight) (*models.Flight, error)
	ListFlights() ([]*models.Flight, error)
	FindFlights(param *models.FindFlightParams) ([]*models.Flight, error)
	BookFlight(id string) (*models.Flight, error)
	DeleteFlight(id string) error
}

type FlightAPIService struct {
	client pb.FlightServiceClient
}

func NewFlightAPIService(serviceConfig *config.Config) IFlightAPIService {
	return &FlightAPIService{
		client: connectFlightService(serviceConfig),
	}
}
func connectFlightService(serviceConfig *config.Config) pb.FlightServiceClient {
	serverAddr := fmt.Sprintf("%v:%v", serviceConfig.GRPCConf.FlightGRPCConf.Host, serviceConfig.GRPCConf.FlightGRPCConf.Port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[Error]: FlightAPIService failed to dial (%v): %v\n", serverAddr, err)
	}
	return pb.NewFlightServiceClient(conn)
}

func (api *FlightAPIService) CreateFlight(flight *models.Flight) (*models.Flight, error) {
	res, err := api.client.CreateFlight(context.Background(), flight.ToPBModel())
	if err != nil {
		return flight, err
	}
	return models.FromPBFlightModel(res), nil
}

func (api *FlightAPIService) UpdateFlight(flight *models.Flight) (*models.Flight, error) {
	res, err := api.client.UpdateFlight(context.Background(), flight.ToPBModel())
	if err != nil {
		return flight, err
	}
	return models.FromPBFlightModel(res), nil
}

func (api *FlightAPIService) ListFlights() ([]*models.Flight, error) {
	log.Printf("[Info] Send ListFlights request\n")
	stream, err := api.client.ListFlights(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	log.Printf("[Info] ListFlights receiving response...\n")
	flights := []*models.Flight{}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return flights, err
		}
		log.Printf("%v\n", res)
		flights = append(flights, models.FromPBFlightModel(res))
	}
	return flights, nil
}

func (api *FlightAPIService) FindFlights(param *models.FindFlightParams) ([]*models.Flight, error) {
	req := &pb.FindFlightRequest{
		Departure:      param.Departure,
		Arrival:        param.Arrival,
		StartTimeRange: models.ToPBDateTime(param.StartTimeRange),
		EndTimeRange:   models.ToPBDateTime(param.EndTimeRange),
	}
	log.Printf("[Info] Send FindFlights request: %v\n", req)
	stream, err := api.client.FindFlights(context.Background(), req)
	if err != nil {
		return nil, err
	}

	log.Printf("[Info] FindFlights receiving response...\n")
	flights := []*models.Flight{}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return flights, err
		}
		log.Printf("%v\n", res)
		flights = append(flights, models.FromPBFlightModel(res))
	}
	return flights, nil
}

func (api *FlightAPIService) BookFlight(id string) (*models.Flight, error) {
	res, err := api.client.BookFlight(context.Background(), &pb.BookingRequest{FlightId: id})
	if err != nil {
		return nil, err
	}
	return models.FromPBFlightModel(res), nil
}

func (api *FlightAPIService) DeleteFlight(id string) error {
	_, err := api.client.DeleteFlight(context.Background(), &pb.DeleteFlightByIdRequest{Id: id})
	return err
}
