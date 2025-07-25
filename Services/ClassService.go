package Services

import (
	"awesomeProject/Context"
	"awesomeProject/Models"
	"awesomeProject/Utils"
	"awesomeProject/Validations"
	"fmt"
	"strings"
)

func ClassService() {
	for {
		println("---------Class service---------")
		println("Get all classes : 1")
		println("Get class by id : 2")
		println("Get class by name : 3")
		println("Get students in class : 4")
		println("Create new class : 5")
		println("Update class : 6")
		println("Delete class : 7")
		println("Get class statistics : 8")
		println("Exit : 0")

		print("Enter your service : ")
		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			getAllClasses()
		case 2:
			getClassById()
		case 3:
			getClassByName()
		case 4:
			getStudentsInClass()
		case 5:
			createNewClass()
		case 6:
			updateClass()
		case 7:
			deleteClass()
		case 8:
			getClassStatistics()
		case 0:
			println("Exiting...")
			return
		default:
			println("Invalid option. Please try again.")
		}

		println("\nPress Enter to continue...")
		fmt.Scan()

	}
}

func getAllClasses() {
	Utils.CleanConsole()
	classes := Context.GetAllClasses()
	if len(classes) == 0 {
		println("No classes found.")
		return
	}

	println("All Classes:")

	Utils.PrintClassesTable(classes)

}

func getClassById() {
	Utils.CleanConsole()
	print("Enter class ID: ")
	var classId int
	_, err := fmt.Scan(&classId)
	if err != nil {
		println("Invalid class ID")
		return
	}

	class, err := Context.GetClassById(classId)
	if err != nil {
		println("Class not found")
		return
	}

	Utils.PrintClassDetailed(class)

	// Show students in this class
	students := Context.GetStudentsByClass(class.ID)
	fmt.Printf("Students count: %d\n", len(students))

	if len(students) > 0 {
		println("Students in this class:")
		for i, student := range students {
			fmt.Printf("  %d. %s %s (Age: %d)\n", i+1, student.Name, student.Surname, student.Age)
		}
	}
}

func getClassByName() {
	Utils.CleanConsole()
	print("Enter class name: ")
	var name string
	fmt.Scan(&name)

	if name == "" {
		println("Class name cannot be empty")
		return
	}

	classes := Context.GetAllClasses()
	var foundClasses []Models.Class

	for _, class := range classes {
		if strings.Contains(strings.ToLower(class.Name), strings.ToLower(name)) {
			foundClasses = append(foundClasses, class)
		}
	}

	if len(foundClasses) == 0 {
		println("No classes found with this name.")
		return
	}

	fmt.Printf("Classes matching '%s':\n", name)
	Utils.PrintClassesTable(classes)
}

func getStudentsInClass() {
	Utils.CleanConsole()
	print("Enter class ID: ")
	var classId int
	_, err := fmt.Scan(&classId)
	if err != nil {
		println("Invalid class ID")
		return
	}

	class, err := Context.GetClassById(classId)
	if err != nil {
		println("Class not found")
		return
	}

	students := Context.GetStudentsByClass(classId)
	if len(students) == 0 {
		fmt.Printf("No students found in class '%s'.\n", class.Name)
		return
	}

	fmt.Printf("Students in class '%s' (ID: %d):\n", class.Name, class.ID)

	Utils.PrintStudentsTable(students)

	fmt.Printf("\nTotal students: %d\n", len(students))
}

func createNewClass() {
	Utils.CleanConsole()
	print("Enter class name: ")
	var name string
	fmt.Scan(&name)
	name = strings.TrimSpace(name)

	if name == "" {
		println("Class name cannot be empty")
		return
	}

	// Check if class with same name already exists
	classes := Context.GetAllClasses()
	for _, class := range classes {
		if strings.EqualFold(class.Name, name) {
			println("A class with this name already exists")
			return
		}
	}

	// Validate input (assuming you have class validation)
	isValid := Validations.ClassNameValidation(name)
	if !isValid {
		println("Invalid class name - validation failed")
		return
	}
	var typeNum int

	fmt.Printf("Enter class type (0: BackEnd, 1: FrontEnd, 2: FullStack): ")

	fmt.Scan(&typeNum)

	var isValidType, maxCount = Validations.ClassTypeValidation(typeNum)
	// Create class

	if !isValidType {
		println("Invalid class type - validation failed")
		return
	}

	newClass := Models.Class{
		Name:     name,
		MaxCount: maxCount,
		Type:     typeNum,
	}

	err := Context.CreateClass(newClass)
	if err != nil {
		println("Error creating class:", err.Error())
		return
	}
	Utils.CleanConsole()
	println("Class created successfully!")
}

func updateClass() {
	Utils.CleanConsole()
	print("Enter class ID: ")
	var classId int
	_, err := fmt.Scan(&classId)
	if err != nil {
		println("Invalid class ID")
		return
	}

	class, err := Context.GetClassById(classId)
	if err != nil {
		println("Class not found")
		return
	}

	// Show current class info
	fmt.Printf("Current class info: %s (ID: %d)\n", class.Name, class.ID)

	// Show students that will be affected
	students := Context.GetStudentsByClass(classId)
	if len(students) > 0 {
		fmt.Printf("Note: This class has %d students. Updating will affect their records.\n", len(students))
	}

	print("Enter new class name (press Enter to keep current): ")
	var name string
	fmt.Scan(&name)
	name = strings.TrimSpace(name)

	if name == "" {
		name = class.Name
	}

	// Check if another class with same name already exists
	classes := Context.GetAllClasses()
	for _, existingClass := range classes {
		if existingClass.ID != class.ID && strings.EqualFold(existingClass.Name, name) {
			println("A class with this name already exists")
			return
		}
	}

	// Validate input
	isValid := Validations.ClassNameValidation(name)
	if !isValid {
		println("Invalid class name - validation failed")
		return
	}
	var typeNum int

	fmt.Printf("Enter class type (1: BackEnd, 2: FrontEnd, 3: FullStack): ")

	fmt.Scan(&typeNum)

	var isValidType, maxCount = Validations.ClassTypeValidation(typeNum - 1)
	// Create class

	if !isValidType {
		println("Invalid class type - validation failed")
		return
	}

	// Update class
	class.Name = name
	class.MaxCount = maxCount
	class.Type = typeNum
	err = Context.UpdateClass(class)
	if err != nil {
		println("Error updating class:", err.Error())
		return
	}
	Utils.CleanConsole()
	println("Class updated successfully!")

	if len(students) > 0 {
		println("All students in this class have been updated with the new class name.")
	}
}

func deleteClass() {
	Utils.CleanConsole()
	print("Enter class ID: ")
	var classId int
	_, err := fmt.Scan(&classId)
	if err != nil {
		println("Invalid class ID")
		return
	}

	class, err := Context.GetClassById(classId)
	if err != nil {
		println("Class not found")
		return
	}

	// Check if class has students
	students := Context.GetStudentsByClass(classId)
	if len(students) > 0 {
		fmt.Printf("Cannot delete class '%s' because it has %d students assigned.\n", class.Name, len(students))
		println("Please move or delete all students from this class first.")

		println("\nStudents in this class:")
		for i, student := range students {
			fmt.Printf("  %d. %s %s (ID: %d)\n", i+1, student.Name, student.Surname, student.ID)
		}
		return
	}

	// Show class details
	fmt.Printf("Class to delete: %s (ID: %d)\n", class.Name, class.ID)
	print("Are you sure you want to delete this class? (y/n): ")

	var answer string
	fmt.Scan(&answer)
	answer = strings.ToLower(strings.TrimSpace(answer))

	if answer != "y" && answer != "yes" {
		println("Delete operation cancelled")
		return
	}

	err = Context.DeleteClass(class.ID)
	if err != nil {
		println("Error deleting class:", err.Error())
		return
	}
	Utils.CleanConsole()
	println("Class deleted successfully!")
}

func getClassStatistics() {

	Utils.CleanConsole()
	classes := Context.GetAllClasses()
	if len(classes) == 0 {
		println("No classes found.")
		return
	}

	println("=== Class Statistics ===")
	fmt.Printf("Total classes: %d\n\n", len(classes))

	var totalStudents int
	var maxStudents, minStudents int
	var maxClassName, minClassName string

	// Initialize min/max with first class
	if len(classes) > 0 {
		firstClassStudents := len(Context.GetStudentsByClass(classes[0].ID))
		maxStudents = firstClassStudents
		minStudents = firstClassStudents
		maxClassName = classes[0].Name
		minClassName = classes[0].Name
	}

	println("Class breakdown:")
	for _, class := range classes {
		students := Context.GetStudentsByClass(class.ID)
		studentCount := len(students)
		totalStudents += studentCount

		fmt.Printf("- %s (ID: %d): %d students\n", class.Name, class.ID, studentCount)

		// Track max/min
		if studentCount > maxStudents {
			maxStudents = studentCount
			maxClassName = class.Name
		}
		if studentCount < minStudents {
			minStudents = studentCount
			minClassName = class.Name
		}
	}

	println("\n=== Summary ===")
	fmt.Printf("Total students across all classes: %d\n", totalStudents)

	if len(classes) > 0 {
		avgStudents := float64(totalStudents) / float64(len(classes))
		fmt.Printf("Average students per class: %.2f\n", avgStudents)
		fmt.Printf("Class with most students: %s (%d students)\n", maxClassName, maxStudents)
		fmt.Printf("Class with least students: %s (%d students)\n", minClassName, minStudents)
	}

	// Show empty classes
	var emptyClasses []string
	for _, class := range classes {
		if len(Context.GetStudentsByClass(class.ID)) == 0 {
			emptyClasses = append(emptyClasses, class.Name)
		}
	}

	if len(emptyClasses) > 0 {
		fmt.Printf("\nEmpty classes (%d): %s\n", len(emptyClasses), strings.Join(emptyClasses, ", "))
	}
}
