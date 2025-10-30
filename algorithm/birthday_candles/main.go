package main

import (
	"fmt"
)

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var maxValue int32
	var maxClandlesCount, equalClandlesCount int32 = 0, 1

	for _, value := range candles {
		if value > maxValue {
			maxClandlesCount = 1
			maxValue = value

		} else if value == maxValue && value != 0 {
			equalClandlesCount++
		}
	}

	if equalClandlesCount > maxClandlesCount {
		return equalClandlesCount
	}

	return maxClandlesCount
}

func main() {
	teste1 := []int32{3, 2, 1, 3}
	resultado1 := birthdayCakeCandles(teste1)
	fmt.Println(resultado1)

	teste2 := []int32{4, 4, 1, 3}
	resultado2 := birthdayCakeCandles(teste2)
	fmt.Println(resultado2)

	teste3 := []int32{1, 5, 5, 5, 2}
	resultado3 := birthdayCakeCandles(teste3)
	fmt.Println(resultado3)
}
