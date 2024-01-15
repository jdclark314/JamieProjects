package plane

import (
	"sim/passenger"
	"testing"
)

// add a test for the isAisleEmpty function
func TestAddToAisle(t *testing.T) {
	// create list of 2 passengers
	passengers := []passenger.Passenger{
		{
			AisleTravelSpeed:  0,
			TimeBoardedPlane:  0,
			TimeSatDown:       16,
			SeatRow:           "",
			CurrentPOS:        0,
			TimeSinceLastMove: 0,
			PassengerId:       0,
		},
		{
			AisleTravelSpeed:  0,
			TimeBoardedPlane:  0,
			TimeSatDown:       18,
			SeatRow:           "",
			CurrentPOS:        0,
			TimeSinceLastMove: 0,
			PassengerId:       1,
		},
	}

	// create a plane
	p := NewPlane(10, 10, passengers)
	// add a passenger to the aisle
	p.AddToAisle(0)
	// check that the passenger was added
	if !p.Aisle[0].Occupied {
		t.Error("Passenger was not added to the aisle")
	}
	// check that the passenger was added to the passenger list
	if p.Aisle[0].Passenger != p.PassengerList[0] {
		t.Error("Passenger was not added to the passenger list")
	}
	// check that the passenger count was updated
	if p.LeftToBoardPassengers != 9 {
		t.Error("Passenger count was not updated")
	}
	// check that the onboard passenger count was updated
	if p.OnboardPassengers != 1 {
		t.Error("Onboard passenger count was not updated")
	}
	// check that the passenger's time boarded was updated
	if p.Aisle[0].Passenger.TimeBoardedPlane != 0 {
		t.Error("Passenger's time boarded was not updated")
	}
}

func TestIsAisleEmpty(t *testing.T) {
	passengers := []passenger.Passenger{
		{
			AisleTravelSpeed:  0,
			TimeBoardedPlane:  10,
			TimeSatDown:       16,
			SeatRow:           "",
			CurrentPOS:        0,
			TimeSinceLastMove: 0,
		},
		{
			AisleTravelSpeed:  0,
			TimeBoardedPlane:  12,
			TimeSatDown:       18,
			SeatRow:           "",
			CurrentPOS:        0,
			TimeSinceLastMove: 0,
		},
	}
	// create a plane
	p := NewPlane(10, 10, passengers)
	// check that the aisle is empty
	if !p.IsAisleEmpty() {
		t.Error("Aisle is not empty")
	}
	// add a passenger to the aisle
	p.AddToAisle(0)
	// check that the aisle is not empty
	if p.IsAisleEmpty() {
		t.Error("Aisle is empty")
	}
}
