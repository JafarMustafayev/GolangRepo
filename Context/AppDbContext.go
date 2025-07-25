package Context

import (
	"awesomeProject/Models"
	"errors"
	"sync"
)

var (
	students []Models.Student
	classes  []Models.Class
	mu       sync.RWMutex // Thread safety üçün
)

func GetAllStudents() []Models.Student {
	mu.RLock()
	defer mu.RUnlock()
	// Copy slice to prevent external modification
	result := make([]Models.Student, len(students))
	copy(result, students)
	return result
}

func GetStudentById(id int) (Models.Student, error) {
	if id <= 0 {
		return Models.Student{}, errors.New("invalid student ID")
	}

	mu.RLock()
	defer mu.RUnlock()

	for _, student := range students {
		if student.ID == id {
			return student, nil
		}
	}
	return Models.Student{}, errors.New("student not found")
}

func GetStudentsByClass(classId int) []Models.Student {
	if classId <= 0 {
		return []Models.Student{}
	}

	mu.RLock()
	defer mu.RUnlock()

	var result []Models.Student
	for _, student := range students {
		if student.ClassId == classId {
			result = append(result, student)
		}
	}
	return result
}

func GetStudentsByNameAndSurname(name string, surname string) []Models.Student {
	mu.RLock()
	defer mu.RUnlock()

	var result []Models.Student
	for _, student := range students {
		nameMatch := name == "" || student.Name == name
		surnameMatch := surname == "" || student.Surname == surname

		if nameMatch && surnameMatch {
			result = append(result, student)
		}
	}
	return result
}

func CreateStudent(student Models.Student) error {
	mu.Lock()
	defer mu.Unlock()

	// Generate unique ID
	student.ID = len(students) + 1

	// Check if ID already exists (safety check)
	for _, existingStudent := range students {
		if existingStudent.ID == student.ID {
			student.ID = getNextStudentID()
			break
		}
	}

	students = append(students, student)
	return nil
}

func UpdateStudent(student Models.Student) error {
	if student.ID <= 0 {
		return errors.New("invalid student ID")
	}

	mu.Lock()
	defer mu.Unlock()

	for i, s := range students {
		if s.ID == student.ID {
			students[i] = student
			return nil
		}
	}
	return errors.New("student not found")
}

func DeleteStudent(id int) error {
	if id <= 0 {
		return errors.New("invalid student ID")
	}

	mu.Lock()
	defer mu.Unlock()

	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}

func getNextStudentID() int {
	maxID := 0
	for _, student := range students {
		if student.ID > maxID {
			maxID = student.ID
		}
	}
	return maxID + 1
}

func GetClassById(id int) (Models.Class, error) {
	if id <= 0 {
		return Models.Class{}, errors.New("invalid class ID")
	}

	mu.RLock()
	defer mu.RUnlock()

	for _, class := range classes {
		if class.ID == id {
			return class, nil
		}
	}
	return Models.Class{}, errors.New("class not found")
}

func GetAllClasses() []Models.Class {
	result := make([]Models.Class, len(classes))
	return result
}

func CreateClass(class Models.Class) error {

	class.ID = getNextClassID()

	classes = append(classes, class)
	return nil
}

func UpdateClass(class Models.Class) error {
	if class.ID <= 0 {
		return errors.New("invalid class ID")
	}

	mu.Lock()
	defer mu.Unlock()

	for i, c := range classes {
		if c.ID == class.ID {
			classes[i] = class

			// Update class name in all students of this class
			for j, student := range students {
				if student.ClassId == class.ID {
					students[j].ClassName = class.Name
				}
			}
			return nil
		}
	}
	return errors.New("class not found")
}

func DeleteClass(id int) error {
	if id <= 0 {
		return errors.New("invalid class ID")
	}

	mu.Lock()
	defer mu.Unlock()

	// Check if any students are assigned to this class
	for _, student := range students {
		if student.ClassId == id {
			return errors.New("cannot delete class with assigned students")
		}
	}

	for i, c := range classes {
		if c.ID == id {
			classes = append(classes[:i], classes[i+1:]...)
			return nil
		}
	}
	return errors.New("class not found")
}

func getNextClassID() int {
	maxID := 0
	for _, class := range classes {
		if class.ID > maxID {
			maxID = class.ID
		}
	}
	return maxID + 1
}

func InitializeSampleData() {
	mu.Lock()
	defer mu.Unlock()

	if len(classes) == 0 {
		classes = append(classes, Models.Class{ID: 1, Name: "AB-102", Type: 0, MaxCount: 20})
		classes = append(classes, Models.Class{ID: 2, Name: "AB-101", Type: 1, MaxCount: 10})
		classes = append(classes, Models.Class{ID: 3, Name: "AB-202", Type: 0, MaxCount: 15})
		classes = append(classes, Models.Class{ID: 4, Name: "AF-102", Type: 1, MaxCount: 15})
	}

	if len(students) == 0 {
		students = append(students, Models.Student{
			ID: 1, Name: "Ali", Surname: "Məmmədov", Age: 19, ClassId: 2, ClassName: "AB-101",
		})
		students = append(students, Models.Student{
			ID: 2, Name: "Jafar", Surname: "Mustafayev", Age: 22, ClassId: 1, ClassName: "AB-102",
		})
		students = append(students, Models.Student{
			ID: 3, Name: "Ayşə", Surname: "Həsənova", Age: 22, ClassId: 2, ClassName: "AB-101",
		})
		students = append(students, Models.Student{
			ID: 4, Name: "Nigar", Surname: "Sultanova", Age: 20, ClassId: 4, ClassName: "AF-102",
		})
		students = append(students, Models.Student{
			ID: 5, Name: "Kamran", Surname: "Qasımov", Age: 21, ClassId: 3, ClassName: "AB-202",
		})
		students = append(students, Models.Student{
			ID: 6, Name: "Leyla", Surname: "Hüseynova", Age: 18, ClassId: 4, ClassName: "AF-102",
		})
		students = append(students, Models.Student{
			ID: 7, Name: "Elvin", Surname: "İsmayılov", Age: 19, ClassId: 1, ClassName: "AB-102",
		})
		students = append(students, Models.Student{
			ID: 8, Name: "Rövşən", Surname: "Səfərov", Age: 23, ClassId: 2, ClassName: "AB-101",
		})
		students = append(students, Models.Student{
			ID: 9, Name: "Zəhra", Surname: "Cəfərova", Age: 20, ClassId: 3, ClassName: "AB-202",
		})
		students = append(students, Models.Student{
			ID: 10, Name: "Murad", Surname: "Əliyev", Age: 21, ClassId: 1, ClassName: "AB-102",
		})
		students = append(students, Models.Student{
			ID: 11, Name: "Günay", Surname: "İbrahimli", Age: 22, ClassId: 4, ClassName: "AF-102",
		})
		students = append(students, Models.Student{
			ID: 12, Name: "Tural", Surname: "Süleymanov", Age: 19, ClassId: 3, ClassName: "AB-202",
		})
		students = append(students, Models.Student{
			ID: 13, Name: "Lalə", Surname: "Nağıyeva", Age: 20, ClassId: 2, ClassName: "AB-101",
		})
		students = append(students, Models.Student{
			ID: 14, Name: "Səid", Surname: "Əsədov", Age: 21, ClassId: 4, ClassName: "AF-102",
		})
	}

}
