package main

import (
	"fmt"
)

type Student struct {
	Name  string
	ID    int
	Grade *int
}

func main() {
	students := make([]Student, 0)

	addStudent(&students, "John Doe", 1, 90)
	addStudent(&students, "Jane Smith", 2, 85)
	addStudent(&students, "Alice Johnson", 3, 95)

	updateGrade(students, 1, 95)
	updateGrade(students, 2, 90)

	displayStudents(students)

	averageGrade := calculateAverageGrade(students)
	fmt.Printf("Average Grade: %.2f\n", averageGrade)
}

func addStudent(students *[]Student, name string, id int, grade int) {
	newStudent := Student{
		Name:  name,
		ID:    id,
		Grade: &grade,
	}
	*students = append(*students, newStudent)
}

func updateGrade(students []Student, id int, grade int) {
	for i := range students {
		if students[i].ID == id {
			students[i].Grade = &grade
			break
		}
	}
}

func displayStudents(students []Student) {
	fmt.Println("Students:")
	for _, student := range students {
		fmt.Printf("Name: %s, ID: %d, Grade: %d\n", student.Name, student.ID, *student.Grade)
	}
}

func calculateAverageGrade(students []Student) float64 {
	totalGrades := 0
	for _, student := range students {
		totalGrades += *student.Grade
	}
	averageGrade := float64(totalGrades) / float64(len(students))
	return averageGrade
}
