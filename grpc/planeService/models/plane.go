package models

import "FlightBookingAPI/pb"

type PlaneStatus int32

const (
	STS_UNSPECIFIED = 0
	STS_CLEANING    = 1
	STS_MAINTAIN    = 2
	STS_READY       = 3
)

type Plane struct {
	// gorm.Model
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

func FromPBModel(p *pb.PlaneInfo) *Plane {
	return &Plane{
		Id:        p.Id,
		Type:      p.Type,
		TotalSeat: p.TotalSeat,
		Status:    PlaneStatus(p.Status),
	}
}
