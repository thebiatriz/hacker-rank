package main

import (
	"fmt"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here

	time, err := time.Parse("03:04:05PM", s)

	if err != nil {
		return "Invalid time format"
	}

	militaryTime := time.Format("15:04:05")
	return militaryTime
}

func main() {
	testTime1 := "07:05:45PM"
	testTime2 := "12:01:00AM"

	result1 := timeConversion(testTime1)
	result2 := timeConversion(testTime2)


	fmt.Println(result1)
	fmt.Println(result2)
}
