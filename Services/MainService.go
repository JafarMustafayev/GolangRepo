package Services

import (
	"awesomeProject/Utils"
	"fmt"
)

func MainService() {

	//Context.InitializeSampleData()
	for {
		Utils.CleanConsole()
		println("---------Main service---------")
		println("1. Class service")
		println("2. Student service")
		println("Any char - Exit")

		print("Enter your service : ")

		var input int
		_, err := fmt.Scan(&input)

		if err != nil {
			break
		}

		switch input {

		case 1:
			ClassService()

		case 2:
			StudentService()
		}
	}

	println("---------End---------")

}
