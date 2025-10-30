package main

import (
	"fmt"
)

/*
 * Complete the 'gameOfThrones' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func gameOfThrones(s string) string {
	// Write your code here
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	var oddCount int = 0
	for _, count := range counts {
		if count%2 != 0 {
			oddCount++
		}
	}

	if oddCount > 1 {
		return "NO"
	}

	return "YES"
}


func main() {
	test1 := "aabbc"    
	test2 := "aabbccdd"
	test3 := "aabbcd"  
	test4 := "aaabbbb"  

	fmt.Printf("Teste 1 de '%s': %s\n", test1, gameOfThrones(test1))
	fmt.Printf("Teste 2 de '%s': %s\n", test2, gameOfThrones(test2))
	fmt.Printf("Teste 3 de '%s': %s\n", test3, gameOfThrones(test3))
	fmt.Printf("Teste 4 de '%s': %s\n", test4, gameOfThrones(test4))
}
