package main

import (
	"fmt"
	"testing"
)

func TestCheckWins(t *testing.T) {
	testString := "one"
	r, i := whatDigitInString(testString)
	if r {
		fmt.Println("The test was true!")
		if i != "1" {
			fmt.Println("Wrong value given :", i)
			t.Fail()
		}
	}

	testString = "two"
	r, i = whatDigitInString(testString)
	if r {
		fmt.Println("The test was true!")
		if i != "2" {
			fmt.Println("Wrong value given")
			fmt.Println("Wrong value given :", i)

			t.Fail()
		}
	}

	testString = "free"
	r, i = whatDigitInString(testString)
	if !r {
		fmt.Println("The test was false!")
		if i != "" {
			fmt.Println("Wrong value given")
			t.Fail()
		}
	}

	testString = "checktwo"
	r, i = whatDigitInString(testString)
	if r {
		fmt.Println("The test was true!")
		if i != "2" {
			fmt.Println("Wrong value given")
			t.Fail()
		}
	}
}

func TestForPartTwo(t *testing.T) {
	entryString := "two1nine"
	result := forPartTwo(entryString)
	fmt.Println("this is the result: ", result)
}
