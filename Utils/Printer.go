package Utils

import (
	"awesomeProject/Models"
	"fmt"
)

// Print student in detailed format
func PrintStudentDetailed(student Models.Student) {
	CleanConsole()
	fmt.Println("=== Student Details ===")
	fmt.Printf("ID: %d\n", student.ID)
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Surname: %s\n", student.Surname)
	fmt.Printf("Age: %d\n", student.Age)
	fmt.Printf("Class: %s (ID: %d)\n", student.ClassName, student.ClassId)
	fmt.Println("=====================")
}

// Print class in detailed format
func PrintClassDetailed(class Models.Class) {
	CleanConsole()
	fmt.Println("=== Class Details ===")
	fmt.Printf("ID: %d\n", class.ID)
	fmt.Printf("Name: %s\n", class.Name)
	fmt.Printf("Currently student count: %d\n", len(class.Students))
	fmt.Printf("Max student count: %d\n", class.MaxCount)
	fmt.Println("===================")
}

func PrintStudentsTable(students []Models.Student) {
	CleanConsole()
	if len(students) == 0 {
		fmt.Println("No students to display.")
		return
	}

	fmt.Println("┌──────┬──────────────────────────────┬─────┬──────────────────┐")
	fmt.Println("│  ID  │            Name              │ Age │      Class       │")
	fmt.Println("├──────┼──────────────────────────────┼─────┼──────────────────┤")

	for _, student := range students {
		fullName := fmt.Sprintf("%s %s", student.Name, student.Surname)
		fmt.Printf("│ %-4d │ %-28s │ %-3d │ %-16s │\n",
			student.ID, fullName, student.Age, student.ClassName)
	}

	fmt.Println("└──────┴──────────────────────────────┴─────┴──────────────────┘")
	fmt.Printf("Total: %d students\n", len(students))
}

// Print a list of classes in table format
func PrintClassesTable(classes []Models.Class) {

	CleanConsole()
	if len(classes) == 0 {
		fmt.Println("No classes to display.")
		return
	}

	fmt.Println("┌──────┬──────────────────────────────┬─────────────────┬───────────────────┐")
	fmt.Println("│  ID  │         Class Name           │ Student Count   │ Max Student Count │")
	fmt.Println("├──────┼──────────────────────────────┼─────────────────┼───────────────────┤")

	for _, class := range classes {
		// You'll need to import Context package to get student count
		// For now, we'll show placeholder
		fmt.Printf("│ %-4d │ %-28s │ %-15d │ %-17d │\n",
			class.ID, class.Name, len(class.Students), class.MaxCount)
	}

	fmt.Println("└──────┴──────────────────────────────┴─────────────────┴───────────────────┘")
	fmt.Printf("Total: %d classes\n", len(classes))
}

func CleanConsole() {
	fmt.Print("\033[H\033[2J")

}
