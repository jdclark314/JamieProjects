package plane

import "testing"

// add a test for the isAisleEmpty function
func TestAddToAisle(t *testing.T) {
	// create a plane
	p := NewPlane(10, 10)
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
	// create a plane
	p := NewPlane(10, 10)
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
