package main

import (
    "fmt"
)

/*
 * Complete the 'twoStrings' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func twoStrings(s1 string, s2 string) string {
    // Write your code here
    s1Map := make(map[rune]bool)

    for _, char := range s1 {
        s1Map[char] = true
    }
    
    for _, char := range s2 {
        if s1Map[char] {
            return "YES"
        }
    }
    
    return "NO"
}

func main() {
    s1_1 := "and"
    s1_2 := "art"
    fmt.Printf("Teste 1 ('and', 'art'): %s\n", twoStrings(s1_1, s1_2))

    s2_1 := "be"
    s2_2 := "cat"
    fmt.Printf("Teste 2 ('be', 'cat'): %s\n", twoStrings(s2_1, s2_2)) 
}
