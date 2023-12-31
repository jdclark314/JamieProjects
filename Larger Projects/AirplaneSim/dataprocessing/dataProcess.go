package dataprocessing

import (
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

func (t *TotalData) calcPassengerData(p passenger.Passenger) {
	d := passengerData{
		PassengerDetails: p,
		// eventually create a function to handle this as the number of metrics grows
		// this will support ease of adding metrics and testibility
		passengerMetrics: passengerMetrics{
			timeToBoard: p.TimeSatDown - p.TimeBoardedPlane,
		},
	}
	t.PassengerList = append(t.PassengerList, d)
}

func (t *TotalData) calcTotalMetrics() {
	// loop through passenger list and add up metrics
	sum := 0
	for _, p := range t.PassengerList {
		sum += p.timeToBoard
	}
	t.totalMetrics.avgTimeToBoard = float32(sum) / float32(len(t.PassengerList))
}

func DataProcess(passengers []passenger.Passenger) TotalData {
	data := TotalData{
		PassengerList: []passengerData{},
		totalMetrics:  totalMetrics{},
	}
	// get all the metrics for the passengers
	for _, p := range passengers {
		data.calcPassengerData(p)
	}
	data.calcTotalMetrics()

	return data
}
