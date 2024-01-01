package plane

import (
	"fmt"
	"sim/passenger"
)

type aisleSpot struct {
	Occupied  bool
	Passenger passenger.Passenger
}

// Plane represents an airplane with an aisle and passenger-related information.
type Plane struct {
	Aisle                 []aisleSpot           // Aisle represents the spots in the aisle of the plane.
	TimeCount             int                   // TimeCount tracks the length of time boarding has been running.
	PassengerList         []passenger.Passenger // PassengerList represents the list of passengers in the boarding order.
	LeftToBoardPassengers int                   // LeftToBoardPassengers represents the number of passengers left to board.
	OnboardPassengers     int                   // OnboardPassengers represents the number of passengers already onboard.
}

// Adds a passenger to the aisle
// Idea is to add to any row, hence the rowNum variable
// There currently isn't any checks to ensure a person isn't added to a spot with someone else there
// These checks should be added once adding to any row is supported
func (a *Plane) AddToAisle(rowNum int) {
	a.Aisle[rowNum].Occupied = true
	a.Aisle[rowNum].Passenger = a.PassengerList[a.OnboardPassengers]
	a.Aisle[rowNum].Passenger.TimeBoardedPlane = a.TimeCount
	a.LeftToBoardPassengers--
	a.OnboardPassengers++
	fmt.Println("Added a new passenger to the Aisle")
}

func (a *Plane) IsAisleEmpty() bool {
	for _, i := range a.Aisle {
		if i.Occupied {
			return false
		}
	}
	return true
}

func NewPlane(aisleLength, passengerCount int) Plane {
	p := passenger.CreatePassengers(passengerCount, aisleLength)
	return Plane{
		Aisle:                 make([]aisleSpot, aisleLength),
		TimeCount:             0,
		PassengerList:         p,
		LeftToBoardPassengers: passengerCount,
		OnboardPassengers:     0,
	}
}
