package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here

	var rest, finalGrade int32 = 0, 0
	var finalGrades = make([]int32, len(grades))

	for i, grade := range grades {
		if grade >= 38 {
			for m := 38; m <= 100; m++ {
				if m%5 == 0 && m >= int(grade) {
					rest = int32(m) - grade

					if rest >= 3 {
						finalGrades[i] = grade
					} else {
						finalGrade = grade + rest
						finalGrades[i] = finalGrade
					}
					break
				}
			}
		} else {
			finalGrades[i] = grade
		}
	}
	return finalGrades
}

func main() {
	inputGrades := []int32{73, 67, 38, 33}

	resultArray := gradingStudents(inputGrades)

	for _, grade := range resultArray {
		fmt.Println(grade)
	}
}
