package models

import (
	"FlightBookingAPI/pb"
	"time"

	graphql "FlightBookingAPI/clients/graphql/models"
)

type Flight struct {
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

func FromPBFlightModel(f *pb.FlightInfo) *Flight {
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

func (f *Flight) ToGraphQLModel() *graphql.Flight {
	return &graphql.Flight{
		ID:        f.Id,
		PlaneID:   f.PlaneId,
		Departure: f.Departure,
		Arrival:   f.Arrival,
		TimeDepart: &graphql.DateTime{
			Year:    f.TimeDepart.Year(),
			Month:   int(f.TimeDepart.Month()),
			Day:     f.TimeDepart.Day(),
			Hours:   f.TimeDepart.Hour(),
			Minutes: f.TimeDepart.Minute(),
		},
		TimeArrive: &graphql.DateTime{
			Year:    f.TimeArrive.Year(),
			Month:   int(f.TimeArrive.Month()),
			Day:     f.TimeArrive.Day(),
			Hours:   f.TimeArrive.Hour(),
			Minutes: f.TimeArrive.Minute(),
		},
		AvailableSeat: int(f.AvailableSeat),
	}
}

func FromGraphQLModel(f *graphql.NewFlight) *Flight {
	return &Flight{
		Id:            f.ID,
		PlaneId:       f.PlaneID,
		Departure:     f.Departure,
		Arrival:       f.Arrival,
		TimeDepart:    FromGraphQLDateTime(f.TimeDepart),
		TimeArrive:    FromGraphQLDateTime(f.TimeArrive),
		AvailableSeat: int32(f.AvailableSeat),
	}
}

func FromGraphQLDateTime(dt *graphql.InputDateTime) time.Time {
	if dt == nil {
		return time.Time{}
	}
	return time.Date(dt.Year, time.Month(dt.Month), dt.Day, dt.Hours, dt.Minutes, 0, 0, &time.Location{})
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
