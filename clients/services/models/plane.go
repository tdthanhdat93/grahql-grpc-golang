package models

import (
	graphql "FlightBookingAPI/clients/graphql/models"
	"FlightBookingAPI/pb"
)

type PlaneStatus int32

const (
	STS_UNSPECIFIED = 0
	STS_CLEANING    = 1
	STS_MAINTAIN    = 2
	STS_READY       = 3
)

type Plane struct {
	Id        string
	Type      string
	TotalSeat int32
	Status    PlaneStatus
}

func (p *Plane) ToPBModel() *pb.PlaneInfo {
	return &pb.PlaneInfo{
		Id:        p.Id,
		Type:      p.Type,
		TotalSeat: p.TotalSeat,
		Status:    pb.PlaneStatus(p.Status),
	}
}

func FromPBPlaneModel(p *pb.PlaneInfo) *Plane {
	return &Plane{
		Id:        p.Id,
		Type:      p.Type,
		TotalSeat: p.TotalSeat,
		Status:    PlaneStatus(p.Status),
	}
}

func (p *Plane) ToGraphQLModel() *graphql.Plane {
	return &graphql.Plane{
		ID:        p.Id,
		Type:      p.Type,
		TotalSeat: int(p.TotalSeat),
		Status:    ToGraphQLEnum(p.Status),
	}
}

func ToGraphQLEnum(e PlaneStatus) graphql.PlaneStatus {
	switch e {
	case STS_UNSPECIFIED:
		return graphql.PlaneStatusStsUnspecified
	case STS_CLEANING:
		return graphql.PlaneStatusStsCleaning
	case STS_MAINTAIN:
		return graphql.PlaneStatusStsMaintain
	case STS_READY:
		return graphql.PlaneStatusStsReady
	default:
		return graphql.PlaneStatusStsUnspecified
	}
}

func FromGraphQLEnum(e *graphql.PlaneStatus) PlaneStatus {
	if e == nil {
		return STS_UNSPECIFIED
	}
	switch *e {
	case graphql.PlaneStatusStsCleaning:
		return STS_CLEANING
	case graphql.PlaneStatusStsMaintain:
		return STS_MAINTAIN
	case graphql.PlaneStatusStsReady:
		return STS_READY
	default:
		return STS_UNSPECIFIED
	}
}
