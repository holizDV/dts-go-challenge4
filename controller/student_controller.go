package controller

import (
	"fmt"

	"github.com/holizDV/dts-go-challenge4/model"
)

func findByNumber(number int, student []model.Student) (*model.Student, error) {
	for i := 0; i < len(student); i++ {
		if student[i].Number == number {
			return &student[i], fmt.Errorf("")
		}

	}
	return nil, nil
}

func FetchDataStudent(number int, students []model.Student) {
	student, err := findByNumber(number, students)

	if err == nil {
		fmt.Printf(" Data number %d is not found, student data only up to %d\n", number, len(students))
		return
	}

	fmt.Printf("Number		: %d\n", student.Number)
	fmt.Printf("Name 		: %s\n", student.Name)
	fmt.Printf("Address 	: %s\n", student.Address)
	fmt.Printf("Profession 	: %s\n", student.Profession)
	fmt.Printf("Reason	 	: %s\n", student.Reason)
}
