package dataprocessing

import (
	"fmt"
	"sim/passenger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	fmt.Println("We in the test yo!")
	passengers := []passenger.Passenger{{
		AisleTravelSpeed:  0,
		TimeBoardedPlane:  10,
		TimeSatDown:       15,
		Seat:              "",
		CurrentPOS:        0,
		TimeSinceLastMove: 0,
	}}
	sut := DataProcess(passengers)
	assert.Equal(t, 5, sut.totalMetrics.avgTimeToBoard)
	// t.Fail()
}
