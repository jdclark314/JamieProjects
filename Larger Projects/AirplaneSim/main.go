package main

import (
	"fmt"
	"sim/dataprocessing"
	"sim/passenger"
	"sim/plane"
	"strconv"
)

var PASSENGER_COUNT = 20
var AISLE_LENGTH = 5

/*


Next Steps: (no order)
- Need process to hold final values of passenger data
- clean up passenger moving through aisle logic
- write tests to see if we get the times we expect
- add ability to have a seat row and letter
- create an interface for passenger so that mock tests can recreate set values
- output the data for how long boarding was and how long each passenger took to board
- run multiple simulations of a plane at once
- generate different boarding orders
- improve logging so that I can only get logs I want to see during a simulation run

*/

// Runs the decision for getting all passengers on plane
// Returns the passengerList slice for processing of data
// Will most likely change that return value as improvements are made
func ProcessPlaneSim() []passenger.Passenger {
	plane := plane.NewPlane(AISLE_LENGTH, PASSENGER_COUNT)

	seatedPassengers := []passenger.Passenger{}

	for ; ; plane.TimeCount++ {
		if plane.IsAisleEmpty() {
			if plane.LeftToBoardPassengers > 0 {
				// 0 to add to first row
				plane.AddToAisle(0)
			} else {
				fmt.Println("passengers all out of line and aisle empty")
				break
			}
		}

		//loop through ailse and move passengers
		for a := len(plane.Aisle) - 1; a >= 0; a-- {
			if a == 0 && !plane.Aisle[a].Occupied {
				// Check if the first spot is empty and need to add a passenger
				if plane.LeftToBoardPassengers > 0 {
					plane.AddToAisle(0)
				}
			} else if plane.Aisle[a].Occupied {
				// Check if there is a passenger in the current spot
				if plane.Aisle[a].Passenger.TimeSinceLastMove > plane.Aisle[a].Passenger.AisleTravelSpeed {
					// Check if enough time has elapsed to move the passenger
					if plane.Aisle[a].Passenger.Seat == strconv.Itoa(a) {
						// Check if the passenger has reached their seat
						plane.Aisle[a].Passenger.TimeSatDown = plane.TimeCount
						seatedPassengers = append(seatedPassengers, plane.Aisle[a].Passenger)
						plane.Aisle[a].Occupied = false
						plane.Aisle[a].Passenger = passenger.Passenger{}
					} else if !plane.Aisle[a+1].Occupied {
						// Check if the next spot is clear to move the passenger forward
						plane.Aisle[a+1].Occupied = true
						plane.Aisle[a+1].Passenger = plane.Aisle[a].Passenger
						plane.Aisle[a].Occupied = false
						plane.Aisle[a].Passenger = passenger.Passenger{}
						plane.Aisle[a+1].Passenger.TimeSinceLastMove = 0
					} else {
						// The next spot is taken, so the passenger can't move
						fmt.Println("The next spot is taken for:", a)
					}
				}
				plane.Aisle[a].Passenger.TimeSinceLastMove++
			}
		}

		//this is the safety catch so I don't infinite loop
		if plane.TimeCount > 1000*PASSENGER_COUNT { //1000 is arbitrary number selected, just want to make sure it doesn't go too far
			fmt.Println("Got to 100 to break")
			break
		}
	}

	if len(seatedPassengers) != PASSENGER_COUNT {
		fmt.Println("ERROR: Not all passengers have been seated")
	}

	fmt.Println("All passengers boarded with time: ", plane.TimeCount)
	return seatedPassengers
}

func main() {
	passengerList := ProcessPlaneSim()
	fmt.Println("This is the passenger list we received in main: ", passengerList)
	dataprocessing.DataProcess(passengerList)
}
