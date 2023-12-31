package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInputData() []string {
	// open file
	f, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	data := []string{}

	for scanner.Scan() {
		// do something with a line
		// fmt.Printf("line: %s\n", scanner.Text())
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func getNumberPartOne(s string) int {
	// fmt.Println("The string: ", s)
	split := strings.Split(s, "")
	l := 0
	r := len(s) - 1
	var leftNum string
	var rightNum string

	for rightNum == "" || leftNum == "" {
		// fmt.Println("This is l: ", split[l])
		// fmt.Println("this is r: ", split[r])
		if leftNum == "" {
			_, errL := strconv.Atoi(split[l])
			if errL == nil {
				leftNum = split[l]
				// fmt.Println("Found the left num: ", leftNum)
			}
		}
		if rightNum == "" {
			_, errR := strconv.Atoi(split[r])
			if errR == nil {
				rightNum = split[r]
				// fmt.Println("Found the right num: ", rightNum)

			}
		}

		if rightNum != "" && leftNum != "" {
			break
		}
		l++
		r--
	}
	answerString := leftNum + rightNum
	// fmt.Println("this is left: ", leftNum)
	// fmt.Println("this is right: ", rightNum)
	answer, err := strconv.Atoi(answerString)
	if err != nil {
		fmt.Println("Something went wrong int he answer string conversion: ", err)
		return -1
	}
	// fmt.Println("This is the answer: ", answer)
	return answer
}

func main() {
	data := getInputData()
	fmt.Println("Here is the length of data: ", len(data))
	getNumberPartOne(data[len(data)-1])
	sum := 0
	sumTwo := 0
	for _, d := range data {
		sum += getNumberPartOne(d)
		sumTwo += forPartTwo(d)
	}
	fmt.Println("This is the total for part 1: ", sum)
	fmt.Println("This is the total for part 2: ", sumTwo)

}

// for second part
/*
Check if single digit is number
	yes: answer
	no: add digit to read values

After digit added to read values
	Check if read value is a spelled out number
		Yes: Answer
		No: Next single digit
*/

// takes a string
// returns whether that string contains the text for the numbers 1-9
// if it does, returns that value as an int
func whatDigitInString(s string) (bool, string) {
	possibleWins := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	textToNumberTranslation := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for _, pw := range possibleWins {
		if strings.Contains(s, pw) {

			return true, textToNumberTranslation[pw]
		}
	}
	return false, ""
}

func forPartTwo(s string) int {
	// fmt.Println("The string: ", s)
	split := strings.Split(s, "")
	l := 0
	r := len(s) - 1
	var leftNum string
	var rightNum string

	lValues := ""
	rValues := ""

	for rightNum == "" || leftNum == "" {
		// fmt.Println("This is l: ", split[l])
		// fmt.Println("this is r: ", split[r])
		if leftNum == "" {
			_, errL := strconv.Atoi(split[l])
			if errL == nil {
				leftNum = split[l]
				// fmt.Println("Found the left num: ", leftNum)
			}
			lValues = lValues + split[l]
			b, s := whatDigitInString(lValues)
			if b {
				// we had a number, now gotta find the number
				leftNum = s
			}
		}
		if rightNum == "" {
			_, errR := strconv.Atoi(split[r])
			if errR == nil {
				rightNum = split[r]
				// fmt.Println("Found the right num: ", rightNum)
			}
			// fmt.Println("R Values before: ", rValues)
			rValues = split[r] + rValues
			// fmt.Println("R Values after: ", rValues)

			b, s := whatDigitInString(rValues)
			if b {
				// we had a number, now gotta find the number
				rightNum = s
			}
		}

		if rightNum != "" && leftNum != "" {
			break
		}
		l++
		r--
	}
	answerString := leftNum + rightNum
	// fmt.Println("this is left: ", leftNum)
	// fmt.Println("this is right: ", rightNum)
	answer, err := strconv.Atoi(answerString)
	if err != nil {
		fmt.Println("Something went wrong int he answer string conversion: ", err)
		return -1
	}
	// fmt.Println("This is the answer: ", answer)
	return answer
}
