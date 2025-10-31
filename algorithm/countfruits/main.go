package main

import (
	"fmt"
)

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var countApples, countOranges int32 = 0, 0

	for _, distance := range apples {
		finalApplePosition := a + distance

		if finalApplePosition >= s && finalApplePosition <= t {
			countApples++
		}
	}

	for _, distance := range oranges {
		finalOrangePosition := b + distance

		if finalOrangePosition >= s && finalOrangePosition <= t {
			countOranges++
		}
	}

	fmt.Printf("%d\n%d", countApples, countOranges)
}

func main() {
	var s int32 = 7
	var t int32 = 10
	var a int32 = 4
	var b int32 = 12
	
	apples := []int32{2, 3, -4}
	oranges := []int32{3, -2, -4}

	countApplesAndOranges(s, t, a, b, apples, oranges)
}