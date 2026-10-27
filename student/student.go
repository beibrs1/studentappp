package student

import "fmt"

type Student struct {
	ID        int
	LastName  string
	FirstName string
	BirthYear int
	Faculty   string
	Course    int
	Group     string
}

var nextID = 1

// NewStudent assigns a unique ID to each student created in this process.
func NewStudent(firstName, lastName string, birthYear int, faculty string, course int, group string) Student {
	s := Student{nextID, lastName, firstName, birthYear, faculty, course, group}
	nextID++
	return s
}

func AddStudent(list *[]Student, s Student) error {
	for _, existing := range *list {
		if existing.ID == s.ID {
			return fmt.Errorf("Студент ID %d қайта қосылмақ болды!", s.ID)
		}
	}
	*list = append(*list, s)
	return nil
}

func PrintAll(list []Student) {
	if len(list) == 0 {
		fmt.Println("Студенттер табылмады.")
		return
	}
	for i, s := range list {
		fmt.Printf("%d. %s %s, ID: %d, Туған жылы: %d, Факультет: %s, Курс: %d, Топ: %s\n", i+1, s.FirstName, s.LastName, s.ID, s.BirthYear, s.Faculty, s.Course, s.Group)
	}
}

func FilterByFaculty(list []Student, faculty string) []Student {
	result := []Student{}
	for _, s := range list {
		if s.Faculty == faculty {
			result = append(result, s)
		}
	}
	return result
}

func FilterByCourse(list []Student, course int) []Student {
	result := []Student{}
	for _, s := range list {
		if s.Course == course {
			result = append(result, s)
		}
	}
	return result
}
