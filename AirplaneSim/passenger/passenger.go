package passenger

import (
	"math/rand"
	"strconv"
)

type Passenger struct {
	AisleTravelSpeed  int
	TimeBoardedPlane  int
	TimeSatDown       int
	Seat              string
	CurrentPOS        int
	TimeSinceLastMove int
}

// determines the AisleTravelSpeed of each passenger
func makeRandomWalkingSpeed() int {
	min := 1
	max := 10
	return rand.Intn(max-min+1) + min
}

// assigns a random seat to a passenger
// currently picks random aisle, but does not care if others in row or keep track of equal rows
func assignRandomSeat(ailseCount int) string {
	min := 0
	max := ailseCount - 1
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

// This passes in the passengerCount and number of ailse
// passenger count to find out how many passengers to create
// number of ailse to find limit of seat assignment numbers
func CreatePassengers(passengerCount int, ailseCount int) []Passenger {
	passengers := []Passenger{}
	for i := 0; i < passengerCount; i++ {
		// need a better way to produce random passenger information
		passengers = append(passengers, Passenger{makeRandomWalkingSpeed(), 0, 0, assignRandomSeat(ailseCount), 0, 0})
	}
	return passengers
}
