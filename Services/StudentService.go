package Services

import (
	"awesomeProject/Context"
	"awesomeProject/Models"
	"awesomeProject/Utils"
	"awesomeProject/Validations"
	"fmt"
	"strings"
)

func StudentService() {
	for {
		println("---------Student service---------")
		println("Get all students : 1")
		println("Get all Students by class : 2")
		println("Get Student by id : 3")
		println("Get Student by name : 4")
		println("Get Student by surname : 5")
		println("Create New Student : 6")
		println("Add Student to Class : 7")
		println("Delete Student : 8")
		println("Update Student : 9")
		println("Exit : 0")

		print("Enter your service : ")
		var input string
		fmt.Scan(&input)

		switch input {
		case "1":
			getAllStudents()
		case "2":
			getStudentsByClass()
		case "3":
			getStudentById()
		case "4":
			getStudentByName()
		case "5":
			getStudentBySurname()
		case "6":
			createNewStudent()
		case "7":
			addStudentToClass()
		case "8":
			deleteStudent()
		case "9":
			updateStudent()
		case "0", "q", "Q":
			println("Exiting...")
			return
		default:
			println("Invalid option. Please try again.")
		}

		println("\nPress Enter to continue...")
		fmt.Scan()
	}
}

func getAllStudents() {
	Utils.CleanConsole()
	students := Context.GetAllStudents()
	if len(students) == 0 {
		println("No students found.")
		return
	}

	println("All Students:")

	Utils.PrintStudentsTable(students)

}

func getStudentsByClass() {
	Utils.CleanConsole()
	print("Enter class ID: ")
	var classId int
	_, err := fmt.Scan(&classId)
	if err != nil {
		println("Invalid class ID")
		return
	}

	// Check if class exists
	_, err = Context.GetClassById(classId)
	if err != nil {
		println("Class not found")
		return
	}

	students := Context.GetStudentsByClass(classId)
	if len(students) == 0 {
		println("No students found in this class.")
		return
	}

	fmt.Printf("Students in class ID %d:\n", classId)
	Utils.PrintStudentsTable(students)

}

func getStudentById() {
	Utils.CleanConsole()
	print("Enter student ID: ")
	var studentId int
	_, err := fmt.Scan(&studentId)
	if err != nil {
		println("Invalid student ID")
		return
	}

	student, err := Context.GetStudentById(studentId)
	if err != nil {
		println("Student not found")
		return
	}
	Utils.PrintStudentDetailed(student)
}

func getStudentByName() {
	Utils.CleanConsole()
	print("Enter student name: ")
	var name string
	fmt.Scan(&name)

	if name == "" {
		println("Name cannot be empty")
		return
	}

	students := Context.GetStudentsByNameAndSurname(name, "")
	if len(students) == 0 {
		println("No students found with this name.")
		return
	}
	Utils.CleanConsole()
	fmt.Printf("Students with name '%s':\n", name)
	Utils.PrintStudentsTable(students)

}

func getStudentBySurname() {
	Utils.CleanConsole()

	print("Enter student surname: ")
	var surname string
	fmt.Scan(&surname)
	surname = strings.TrimSpace(surname)

	if surname == "" {
		println("Surname cannot be empty")
		return
	}

	students := Context.GetStudentsByNameAndSurname("", surname)
	if len(students) == 0 {
		println("No students found with this surname.")
		return
	}

	fmt.Printf("Students with surname '%s':\n", surname)
	Utils.PrintStudentsTable(students)

}

func createNewStudent() {
	Utils.CleanConsole()

	print("Enter student name: ")

	var name string
	fmt.Scan(&name)

	print("Enter student surname: ")
	var surname string
	fmt.Scan(&surname)
	surname = strings.TrimSpace(surname)

	print("Enter student age: ")
	var age int
	_, err := fmt.Scan(&age)
	if err != nil {
		println("Invalid age")
		return
	}

	print("Enter student class id: ")
	var classId int
	_, err = fmt.Scan(&classId)
	if err != nil {
		println("Invalid class ID")
		return
	}

	// Validate input
	isValid := Validations.StudentValidation(name, surname, age)
	if !isValid {
		println("Invalid input - validation failed")
		return
	}

	// Check if class exists
	class, err := Context.GetClassById(classId)
	if err != nil {
		println("Class not found")
		return
	}

	// Create student
	newStudent := Models.Student{
		Name:      name,
		Surname:   surname,
		Age:       age,
		ClassId:   class.ID,
		ClassName: class.Name,
	}

	Context.CreateStudent(newStudent)
	Utils.CleanConsole()
	println("Student created successfully!\n\n")
}

func addStudentToClass() {
	Utils.CleanConsole()
	print("Enter student ID: ")
	var studentId int
	_, err := fmt.Scan(&studentId)
	if err != nil {
		println("Invalid student ID")
		return
	}

	print("Enter class ID: ")
	var classId int
	_, err = fmt.Scan(&classId)
	if err != nil {
		println("Invalid class ID")
		return
	}

	// Get student and class
	student, err := Context.GetStudentById(studentId)
	if err != nil {
		println("Student not found")
		return
	}

	class, err := Context.GetClassById(classId)
	if err != nil {
		println("Class not found")
		return
	}

	// Update student's class
	student.ClassId = class.ID
	student.ClassName = class.Name

	Context.UpdateStudent(student)
	Utils.CleanConsole()
	println("Student added to class successfully!")
}

func deleteStudent() {
	Utils.CleanConsole()
	print("Enter student ID: ")
	var studentId int
	_, err := fmt.Scan(&studentId)
	if err != nil {
		println("Invalid student ID")
		return
	}

	student, err := Context.GetStudentById(studentId)
	if err != nil {
		println("Student not found")
		return
	}

	// Show student details
	fmt.Printf("Student to delete: %s %s (ID: %d)\n", student.Name, student.Surname, student.ID)
	print("Are you sure you want to delete this student? (y/n): ")

	var answer string
	fmt.Scanln(&answer)
	answer = strings.ToLower(strings.TrimSpace(answer))

	if answer != "y" && answer != "yes" {
		println("Delete operation cancelled")
		return
	}

	Context.DeleteStudent(student.ID)
	Utils.CleanConsole()
	println("Student deleted successfully!")
}

func updateStudent() {
	Utils.CleanConsole()
	print("Enter student ID: ")
	var studentId int
	_, err := fmt.Scan(&studentId)
	if err != nil {
		println("Invalid student ID")
		return
	}

	student, err := Context.GetStudentById(studentId)
	if err != nil {
		println("Student not found")
		return
	}

	// Show current student info
	fmt.Printf("Current student info: %s %s, Age: %d\n", student.Name, student.Surname, student.Age)

	print("Enter new name (press Enter to keep current): ")
	var name string
	fmt.Scan(&name)
	if name == "" {
		name = student.Name
	}

	print("Enter new surname (press Enter to keep current): ")
	var surname string
	fmt.Scan(&surname)
	surname = strings.TrimSpace(surname)
	if surname == "" {
		surname = student.Surname
	}

	print("Enter new age (0 to keep current): ")
	var age int
	_, err = fmt.Scan(&age)
	if err != nil || age < 0 {
		println("Invalid age")
		return
	}
	if age == 0 {
		age = student.Age
	}

	// Validate input
	isValid := Validations.StudentValidation(name, surname, age)
	if !isValid {
		println("Invalid input - validation failed")
		return
	}

	// Update student
	student.Name = name
	student.Surname = surname
	student.Age = age

	Context.UpdateStudent(student)
	Utils.CleanConsole()
	println("Student updated successfully!")
}
