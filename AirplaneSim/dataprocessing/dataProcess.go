package dataprocessing

import (
	"fmt"
	"sim/passenger"
)

type TotalData struct {
	PassengerList []passengerData
	totalMetrics
}

type passengerData struct {
	PassengerDetails passenger.Passenger
	passengerMetrics
}

type totalMetrics struct {
	avgTimeToBoard float32
}

type passengerMetrics struct {
	timeToBoard int
}

func DataProcess(passengers []passenger.Passenger) TotalData {
	fmt.Println("This is the passenger list we receive in collectData: ", passengers)
	data := TotalData{
		PassengerList: []passengerData{},
		totalMetrics:  totalMetrics{},
	}
	// get all the metrics for the passengers
	for _, p := range passengers {
		fmt.Println("This is my passenger: ", p)
		fmt.Printf("The passenger boarded at %v and sat down at %v\n", p.TimeBoardedPlane, p.TimeSatDown)
		d := passengerData{
			PassengerDetails: p,
			// eventually create a function to handle this as the number of metrics grows
			// this will support ease of adding metrics and testibility
			passengerMetrics: passengerMetrics{
				timeToBoard: p.TimeSatDown - p.TimeBoardedPlane,
			},
		}
		data.PassengerList = append(data.PassengerList, d)
	}

	fmt.Printf("heres what we have so far: %+v\n", data)
	return data
}
