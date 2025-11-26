package day2

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name  string
	Age   int
	Marks []float64
}

func (s student) average() float64 {
	if len(s.Marks) == 0 {
		return 0.0
	}
	total := 0.0
	for _, mark := range s.Marks {
		total += mark
	}
	return total / float64(len(s.Marks))
}

func (s student) grade() string {
	avg := s.average()
	switch {
	case avg >= 90:
		return "A"
	case avg >= 80:
		return "B"
	case avg >= 70:
		return "C"
	case avg >= 60:
		return "D"
	default:
		return "F"
	}
}

func (s student) toJSON() (string, error) {
	data := map[string]interface{}{
		"name":    s.Name,
		"age":     s.Age,
		"marks":   s.Marks,
		"average": fmt.Sprintf("%.2f", s.average()),
		"grade":   s.grade(),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func createStudent(name string, age int, marks []float64) student {
	return student{
		Name:  name,
		Age:   age,
		Marks: marks,
	}
}

func updateStudentName(s *student, newName string) {
	s.Name = newName
}

func Day2() {
	student1 := createStudent("Alice", 20, []float64{85.5, 90.0, 78.5})
	student2 := createStudent("Bob", 22, []float64{92.0, 88.5, 95.0})

	updateStudentName(&student1, "Alicia")

	json1, err := student1.toJSON()
	if err != nil {
		fmt.Println("Error converting student1 to JSON:", err)
	} else {
		fmt.Println("Student 1 JSON:\n", json1)
	}

	json2, err := student2.toJSON()
	if err != nil {
		fmt.Println("Error converting student2 to JSON:", err)
	} else {
		fmt.Println("Student 2 JSON:\n", json2)
	}
}
