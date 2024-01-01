package dataprocessing

import (
	"sim/passenger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	passengers := []passenger.Passenger{
		{
			AisleTravelSpeed:  0,
			TimeBoardedPlane:  10,
			TimeSatDown:       16,
			Seat:              "",
			CurrentPOS:        0,
			TimeSinceLastMove: 0,
		},
		{
			AisleTravelSpeed:  0,
			TimeBoardedPlane:  10,
			TimeSatDown:       30,
			Seat:              "",
			CurrentPOS:        0,
			TimeSinceLastMove: 0,
		},
	}
	sut := DataProcess(passengers)
	assert.Equal(t, float32(13), sut.totalMetrics.avgTimeToBoard)
	// t.Fail()
}
func TestCalcPassengerData(t *testing.T) {
	p := passenger.Passenger{
		AisleTravelSpeed:  0,
		TimeBoardedPlane:  10,
		TimeSatDown:       16,
		Seat:              "",
		CurrentPOS:        0,
		TimeSinceLastMove: 0,
	}

	totData := TotalData{}
	totData.calcPassengerData(p)

	expectedData := passengerData{
		PassengerDetails: p,
		PassengerMetrics: passengerMetrics{
			timeToBoard: p.TimeSatDown - p.TimeBoardedPlane,
		},
	}

	assert.Equal(t, []passengerData{expectedData}, totData.PassengerList)
}
func TestCalcTotalMetrics(t *testing.T) {
	passengerList := []passengerData{
		{
			PassengerDetails: passenger.Passenger{
				AisleTravelSpeed:  0,
				TimeBoardedPlane:  10,
				TimeSatDown:       16,
				Seat:              "",
				CurrentPOS:        0,
				TimeSinceLastMove: 0,
			},
			PassengerMetrics: passengerMetrics{
				timeToBoard: 6,
			},
		},
		{
			PassengerDetails: passenger.Passenger{
				AisleTravelSpeed:  0,
				TimeBoardedPlane:  10,
				TimeSatDown:       30,
				Seat:              "",
				CurrentPOS:        0,
				TimeSinceLastMove: 0,
			},
			PassengerMetrics: passengerMetrics{
				timeToBoard: 20,
			},
		},
	}

	totalData := TotalData{
		PassengerList: passengerList,
		totalMetrics:  totalMetrics{},
	}

	totalData.calcTotalMetrics()

	expectedAvgTimeToBoard := float32(13)
	assert.Equal(t, expectedAvgTimeToBoard, totalData.totalMetrics.avgTimeToBoard)
}
