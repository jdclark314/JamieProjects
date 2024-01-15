package dataprocessing

import (
	"fmt"
	"os"
	"path/filepath"
	"sim/passenger"
)

type TotalData struct {
	PassengerList []passengerData
	totalMetrics
}

type passengerData struct {
	PassengerDetails passenger.Passenger
	PassengerMetrics passengerMetrics
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
		PassengerMetrics: passengerMetrics{
			timeToBoard: p.TimeSatDown - p.TimeBoardedPlane,
		},
	}
	t.PassengerList = append(t.PassengerList, d)
}

func (t *TotalData) calcTotalMetrics() {
	// loop through passenger list and add up metrics
	sum := 0
	for _, p := range t.PassengerList {
		sum += p.PassengerMetrics.timeToBoard
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
	data.outputIndividualData()

	return data
}

// output the data to a csv file
// writing data to  a file isn't part of the data process, this should be in its own package
func (t *TotalData) outputIndividualData() {
	// create the file
	fileName := filepath.Join("data", "data.csv")
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// write the header
	_, err = f.WriteString("PassengerId,AisleTravelSpeed,TimeBoardedPlane,TimeSatDown,Seat,CurrentPOS,TimeSinceLastMove,TimeToBoard\n")
	if err != nil {
		fmt.Println(err)
	}

	// write the data
	for _, p := range t.PassengerList {
		_, err = f.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v\n", p.PassengerDetails.PassengerId, p.PassengerDetails.AisleTravelSpeed, p.PassengerDetails.TimeBoardedPlane, p.PassengerDetails.TimeSatDown, p.PassengerDetails.SeatRow, p.PassengerDetails.CurrentPOS, p.PassengerDetails.TimeSinceLastMove, p.PassengerMetrics.timeToBoard))
		if err != nil {
			fmt.Println(err)
		}
	}
}
