package main

import (
	"fmt"
)

type Student struct {
	name  string
	score []int
}

func main() {
	students := make([]Student, 5)
	var totalScore, minScore, maxScore int
	var minStudent, maxStudent Student

	for i := 0; i < 5; i++ {
		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scan(&students[i].name)

		fmt.Printf("Input %d Student's Score: ", i+1)
		var score int
		fmt.Scan(&score)
		students[i].score = append(students[i].score, score)
		totalScore += score
	}
	rataRata := float64(totalScore) / float64(len(students))
	fmt.Printf("Average score : %.2f\n", rataRata)

	minScore = students[0].score[0]
	for _, student := range students {
		for _, score := range student.score {
			if score < minScore {
				minScore = score
				minStudent = student
			}
		}
	}
	fmt.Printf("Min Score of Students : %s (%d)\n", minStudent.name, minScore)

	maxScore = students[0].score[0]
	for _, student := range students {
		for _, score := range student.score {
			if score > maxScore {
				maxScore = score
				maxStudent = student
			}
		}
	}
	fmt.Printf("Max Score of Students : %s (%d)\n", maxStudent.name, maxScore)
}
