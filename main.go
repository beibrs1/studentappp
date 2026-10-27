package main

import (
	"fmt"
	"studentapp/logger"
	"studentapp/student"
	"studentapp/utils"
)

func main() {
	students := []student.Student{}
	initial := []student.Student{
		student.NewStudent("Алия", "Сулеймен", 2005, "IT", 2, "CS-32"),
		student.NewStudent("Ерлан", "Талғат", 2006, "Math", 1, "M-11"),
		student.NewStudent("Айгерім", "Серік", 2004, "IT", 3, "CS-31"),
	}
	for _, s := range initial {
		if err := student.AddStudent(&students, s); err != nil {
			logger.Error(err.Error())
		} else {
			logger.Info(fmt.Sprintf("Студент қосылды: %s %s (%s)", s.FirstName, s.LastName, s.Group))
		}
	}

	logger.Info("Барлық студенттер:")
	student.PrintAll(students)
	fmt.Println("Факультет бойынша сүзу: IT")
	student.PrintAll(student.FilterByFaculty(students, "IT"))
	fmt.Println("Курс бойынша сүзу: 2")
	student.PrintAll(student.FilterByCourse(students, 2))

	if err := student.AddStudent(&students, students[0]); err != nil {
		logger.Error(err.Error())
	}

	firstName := utils.ReadString("Аты: ")
	lastName := utils.ReadString("Тегі: ")
	birthYear := utils.ReadInt("Туған жылы: ")
	faculty := utils.ReadString("Факультет: ")
	course := utils.ReadInt("Курс: ")
	group := utils.ReadString("Топ: ")
	s := student.NewStudent(firstName, lastName, birthYear, faculty, course, group)
	if err := student.AddStudent(&students, s); err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Студент қосылды: " + firstName + " " + lastName)
	}
	student.PrintAll(students)
}
