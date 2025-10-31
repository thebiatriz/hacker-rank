package main

import (
    "fmt"
)

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
  if v1 == v2 {
        if x1 == x2 {
            return "YES"
        } else {
            return "NO"
        }
    }
 
    if (x2 - x1) % (v1 - v2) == 0 {
        j := (x2 - x1) / (v1 - v2)
        
        if j >= 0 {
            return "YES"
        } else {
            return "NO"
        }
    }
    
    return "NO"
}


func main() {
	fmt.Println(kangaroo(0, 3, 4, 2)) 
	
	fmt.Println(kangaroo(0, 5, 1, 3))
	
	fmt.Println(kangaroo(0, 2, 1, 2))

	fmt.Println(kangaroo(0, 3, -1, 4)) 
}