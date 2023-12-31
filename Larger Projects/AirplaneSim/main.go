package main

import (
	"fmt"
	"sim/dataprocessing"
	"sim/passenger"
	"sim/plane"
	"strconv"
)

var PASSENGER_COUNT = 2
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
			if a == 0 { //if we're at the first spot, check to see if we need to add a passenger
				if !plane.Aisle[a].Occupied { //if no one is on the spot
					if plane.LeftToBoardPassengers > 0 {
						plane.AddToAisle(0)
					} else {
						// fmt.Println("passengers all out of line")
					}
				}
			}
			//check if passenger is there
			if plane.Aisle[a].Occupied {
				//if yes move them
				// check to see if enough time has elapsed to move
				if plane.Aisle[a].Passenger.TimeSinceLastMove > plane.Aisle[a].Passenger.AisleTravelSpeed {
					// check to see if passenger in aisle that has their seat
					if plane.Aisle[a].Passenger.Seat == strconv.Itoa(a) {
						plane.Aisle[a].Passenger.TimeSatDown = plane.TimeCount
						seatedPassengers = append(seatedPassengers, plane.Aisle[a].Passenger)
						plane.Aisle[a].Occupied = !plane.Aisle[a].Occupied
						plane.Aisle[a].Passenger = passenger.Passenger{}

					} else if !plane.Aisle[a+1].Occupied {
						//check if next spot is clear
						//move the passenger forward
						plane.Aisle[a+1].Occupied = !plane.Aisle[a+1].Occupied
						plane.Aisle[a+1].Passenger = plane.Aisle[a].Passenger
						plane.Aisle[a].Occupied = !plane.Aisle[a].Occupied
						plane.Aisle[a].Passenger = passenger.Passenger{}
						// set time since last move to 0 since they moved
						plane.Aisle[a+1].Passenger.TimeSinceLastMove = 0
					} else { //spot not clear so can't move
						fmt.Println("the next spot is taken for: ", a)
						// didn't move so increment
					}
				}
				// increment time for this passenger
				plane.Aisle[a].Passenger.TimeSinceLastMove += 1
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
