package resolver

import "FlightBookingAPI/clients/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	planeService  services.IPlaneAPIService
	flightService services.IFlightAPIService
}

func NewResolver(planeService services.IPlaneAPIService, flightService services.IFlightAPIService) *Resolver {
	return &Resolver{
		planeService:  planeService,
		flightService: flightService,
	}
}
