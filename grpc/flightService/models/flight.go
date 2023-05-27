package models

import (
	"FlightBookingAPI/pb"
	"time"
)

type Flight struct {
	// gorm.Model
	Id            string
	PlaneId       string
	Departure     string
	Arrival       string
	TimeDepart    time.Time
	TimeArrive    time.Time
	AvailableSeat int32
}
type FindFlightParams struct {
	Departure      string
	Arrival        string
	StartTimeRange time.Time
	EndTimeRange   time.Time
}

func (f *Flight) ToPBModel() *pb.FlightInfo {
	return &pb.FlightInfo{
		Id:            f.Id,
		PlaneId:       f.PlaneId,
		Departure:     f.Departure,
		Arrival:       f.Arrival,
		TimeDepart:    ToPBDateTime(f.TimeDepart),
		TimeArrive:    ToPBDateTime(f.TimeArrive),
		AvailableSeat: f.AvailableSeat,
	}
}

func FromPBModel(f *pb.FlightInfo) *Flight {
	return &Flight{
		Id:            f.Id,
		PlaneId:       f.PlaneId,
		Departure:     f.Departure,
		Arrival:       f.Arrival,
		TimeDepart:    FromPBDateTime(f.TimeDepart),
		TimeArrive:    FromPBDateTime(f.TimeArrive),
		AvailableSeat: f.AvailableSeat,
	}
}

func ToPBDateTime(t time.Time) *pb.DateTime {
	return &pb.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
	}
}

func FromPBDateTime(dt *pb.DateTime) time.Time {
	if dt == nil {
		return time.Time{}
	}
	if dt.Year == 1 && dt.Month == 1 && dt.Day == 1 && dt.Hours == 0 && dt.Minutes == 0 { // Zero value of time: January 1, year 1, 00:00:00 UTC.
		return time.Time{}
	}
	return time.Date(
		int(dt.Year),
		time.Month(dt.Month),
		int(dt.Day),
		int(dt.Hours),
		int(dt.Minutes),
		0, 0,
		time.Local,
	)
}
